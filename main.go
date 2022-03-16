package main

import (
	"Backend/ent"
	"Backend/ent/rom"
	"context"
	"crypto/md5"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"entgo.io/ent/dialect"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

var errNoGameIdentified = errors.New("no game identified")

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

func visitRoms(client *ent.Client, library *[]*ent.FileCreate, ctx context.Context) fs.WalkDirFunc {
	return func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		game, err := identityROM(path, client, ctx)
		if err != nil {
			if err == errNoGameIdentified {
				return nil
			}
			return err
		}
		*library = append(*library, client.File.Create().SetPath(path).SetRom(game))
		return nil
	}
}

func visitDAT(client *ent.Client, ctx context.Context) fs.WalkDirFunc {
	return func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		var datafile Datafile
		data, _ := ioutil.ReadFile(path)

		err = xml.Unmarshal(data, &datafile)
		if err != nil {
			log.Fatalf("%v", err)
		}

		err = fillTable(client, ctx, datafile)
		if err != nil {
			log.Fatalf("%v", err)
		}

		return nil
	}
}

func identityROM(path string, client *ent.Client, ctx context.Context) (game *ent.Rom, err error) {
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

	dbRom, err := client.Rom.Query().Where(rom.Md5EQ(sprintf)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errNoGameIdentified
		}
		if ent.IsNotSingular(err) {
			println("duplicated rom!")
			return nil, errNoGameIdentified
		}
		return nil, err
	} else {
		return dbRom, nil
	}

	return nil, errNoGameIdentified
}

func fillTable(client *ent.Client, ctx context.Context, datafile Datafile) error {
	dbDatafile, err := client.Datafile.Create().Save(ctx)
	if err != nil {
		return err
	}

	_, err = client.Header.Create().SetDatafile(dbDatafile).SetAuthor(datafile.Header.Author).SetName(datafile.Header.Name).SetDescription(datafile.Header.Description).SetVersion(datafile.Header.Version).SetDate(datafile.Header.Date).SetURL(datafile.Header.URL).Save(ctx)
	if err != nil {
		return err
	}

	bulkGames := make([]*ent.GameCreate, len(datafile.Games))
	for i, game := range datafile.Games {
		bulkGames[i] = client.Game.Create().SetCloneof(game.Cloneof).SetDatafile(dbDatafile).SetDescription(game.Description).SetName(game.Name)
	}
	games, err := client.Game.CreateBulk(bulkGames...).Save(ctx)
	if err != nil {
		return err
	}

	bulkReleases := make([]*ent.ReleaseCreate, 0)
	bulkRoms := make([]*ent.RomCreate, 0)

	for i, game := range datafile.Games {
		for _, release := range game.Release {
			bulkReleases = append(bulkReleases, client.Release.Create().SetGame(games[i]).SetName(release.Name).SetRegion(release.Region))
		}
		if game.Rom.Name != "" {
			bulkRoms = append(bulkRoms, client.Rom.Create().SetGame(games[i]).SetMd5(game.Rom.Md5).SetCrc(game.Rom.Crc).SetName(game.Rom.Name).SetSize(game.Rom.Size).SetSha1(game.Rom.Sha1).SetStatus(game.Rom.Status))
		}
	}

	var releaseChunks [][]*ent.ReleaseCreate
	for i := 0; i < len(bulkReleases); i += 999 {
		end := i + 999

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(bulkReleases) {
			end = len(bulkReleases)
		}
		releaseChunks = append(releaseChunks, bulkReleases[i:end])
	}

	for _, chunk := range releaseChunks {
		_, err = client.Release.CreateBulk(chunk...).Save(ctx)
		if err != nil {
			return err
		}
	}

	var romChunks [][]*ent.RomCreate
	for i := 0; i < len(bulkRoms); i += 999 {
		end := i + 999

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(bulkRoms) {
			end = len(bulkRoms)
		}
		romChunks = append(romChunks, bulkRoms[i:end])
	}

	for _, chunk := range romChunks {
		_, err = client.Rom.CreateBulk(chunk...).Save(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func initDatabase() {
	client, err := ent.Open(dialect.SQLite, "file:arcadia.db?mode=rwc&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	defer func(client *ent.Client) {
		_ = client.Close()
	}(client)
	ctx := context.Background()

	createDatabase(client, ctx, err)

}

func createDatabase(client *ent.Client, ctx context.Context, err error) {
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	start := time.Now()
	//err = filepath.WalkDir("./rom", visitRoms(&HashPointer, &GameLibrary))
	err = filepath.WalkDir("./dat", visitDAT(client, ctx))
	if err != nil {
		log.Fatalf("%v", err)
	}

	bulkFiles := make([]*ent.FileCreate, 0)

	err = filepath.WalkDir("/Volumes/Kuma/ROMs/Nintendo Gameboy Advance", visitRoms(client, &bulkFiles, ctx))
	if err != nil {
		println(err)
	}
	duration := time.Since(start)
	fmt.Println(duration)

	_, err = client.File.CreateBulk(bulkFiles...).Save(ctx)
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
