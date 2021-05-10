package main

import (
	"encoding/xml"
	"io/ioutil"
)


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
	Game []struct {
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
	} `xml:"game"`
}

func main() {
	var GBADatafile Datafile

	data, _ := ioutil.ReadFile("dat/Nintendo - Game Boy Advance (Parent-Clone) (20210506-095002).dat")

	_ = xml.Unmarshal([]byte(data), &GBADatafile)

	print("done!")

}
