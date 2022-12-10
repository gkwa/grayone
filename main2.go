package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

type description struct {
	Name    string `xml:"name"`
	Version string `xml:"version"`
}

func main() {
	file, err := os.Open("application.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var buf bytes.Buffer
	decoder := xml.NewDecoder(file)
	encoder := xml.NewEncoder(&buf)

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error getting token: %v\n", err)
			break
		}

		switch v := token.(type) {
		case xml.StartElement:
			if v.Name.Local == "description" {
				var desc description
				if err = decoder.DecodeElement(&desc, &v); err != nil {
					log.Fatal(err)
				}
				// modify the version value and encode the element back
				desc.Version = "2.0"
				if err = encoder.EncodeElement(desc, v); err != nil {
					log.Fatal(err)
				}
				continue
			}
		}

		if err := encoder.EncodeToken(xml.CopyToken(token)); err != nil {
			log.Fatal(err)
		}
	}

	// must call flush, otherwise some elements will be missing
	if err := encoder.Flush(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf.String())
}
