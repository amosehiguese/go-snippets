package snippet

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	_ "github.com/lib/pq"
)

type NPost struct {
	XMLName 			xml.Name		`xml:"post"`
	Id     	      string      `xml:"id,attr"`
	Content       string      `xml:"content"`
	Author        Author     	`xml:"author"`
	Xml 					string 			`xml:",innerxml"`
	Comments      []NComment   `xml:"comments>comment"`
}

type NComment struct {
	Id      string     `xml:"id,attr"`
	Content string     `xml:"content"`
	Author  string     `xml:"author"`
}

type Author struct {
	Id     	string 			`xml:"id,attr"`
	Name    string     	`xml:",chardata"`
}

func XMlmain() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding XML into tokens:", err)
			return
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment Comment
				decoder.DecodeElement(&comment, &se)
				fmt.Println(se)
			}
		}
	}
}