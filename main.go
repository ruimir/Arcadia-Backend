package main

import (
	"Backend/ent"
	"context"
	"crypto/md5"
	"encoding/xml"
	"entgo.io/ent/dialect"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var noGameIdentified = errors.New("no game identified")

type Game struct {
	Text        string `xml:",chardata"`
	Name        string `xml:"name,attr"`
	Cloneof     string `xml:"cloneof,attr"`
	Description string `xml:"description"`
	Release     []struct {
		Text   string `xml:",chardata"`
		Name   string `xml:"name,attr"`
		Region string `xml:"region,attr"`
	} `xml:"release"`
	Rom struct {
		Text   string `xml:",chardata"`
		Name   string `xml:"name,attr"`
		Size   string `xml:"size,attr"`
		Crc    string `xml:"crc,attr"`
		Md5    string `xml:"md5,attr"`
		Sha1   string `xml:"sha1,attr"`
		Status string `xml:"status,attr"`
	} `xml:"rom"`
}

type Datafile struct {
	XMLName xml.Name `xml:"datafile"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text        string `xml:",chardata"`
		Name        string `xml:"name"`
		Description string `xml:"description"`
		Version     string `xml:"version"`
		Date        string `xml:"date"`
		Author      string `xml:"author"`
		URL         string `xml:"url"`
	} `xml:"header"`
	Games []Game `xml:"game"`
}

func visit(datafile *map[string]*Game, library *[]string) fs.WalkDirFunc {
	return func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		game, err := identityGame(path, datafile)
		if err != nil {
			if err == noGameIdentified {
				return nil
			}
			return err
		}
		*library = append(*library, game.Name)
		return nil
	}
}

func identityGame(path string, datafile *map[string]*Game) (game *Game, err error) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	sprintf := fmt.Sprintf("%x", h.Sum(nil))

	ok := false
	game, ok = (*datafile)[sprintf]
	if ok {
		return game, nil
	}

	return &Game{}, noGameIdentified
}

func main() {
	var GBADatafile Datafile
	HashPointer := make(map[string]*Game, 0)
	GameLibrary := make([]string, 0)

	data, _ := ioutil.ReadFile("dat/Nintendo - Game Boy Advance (Parent-Clone) (20210506-095002).dat")

	_ = xml.Unmarshal(data, &GBADatafile)

	client, err := ent.Open(dialect.SQLite, "file:arcadia.db?mode=rwc&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	err, _ = fillTable(err, client, ctx, GBADatafile)

	for i, game := range GBADatafile.Games {
		HashPointer[game.Rom.Md5] = &GBADatafile.Games[i]
	}

	start := time.Now()
	err = filepath.WalkDir("D:\\ROMs\\Nintendo Gameboy Advance", visit(&HashPointer, &GameLibrary))
	if err != nil {
		println(err)
	}
	duration := time.Since(start)
	fmt.Println(duration)

	fmt.Printf("%v", GameLibrary)

}

func fillTable(err error, client *ent.Client, ctx context.Context, GBADatafile Datafile) (error, bool) {
	datafile, err := client.Datafile.Create().Save(ctx)
	if err != nil {
		return nil, true
	}

	_, err = client.Header.Create().SetDatafile(datafile).SetAuthor(GBADatafile.Header.Author).Save(ctx)
	if err != nil {
		return nil, true
	}
	return err, false
}
