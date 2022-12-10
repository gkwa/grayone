package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

type rawXML struct {
	Inner []byte `xml:",innerxml"`
}

type description struct {
	Name    string `xml:"name"`
	Version string `xml:"version"`
}

type app struct {
	Description description `xml:"description"`
	Developer   rawXML      `xml:"developer"`
	Market      rawXML      `xml:"market"`
}

func main() {
	data, err := ioutil.ReadFile("application.xml")
	if err != nil {
		log.Fatal(err)
	}

	var myapp app
	if err := xml.Unmarshal(data, &myapp); err != nil {
		log.Fatal(err)
	}

	// modify the version value
	myapp.Description.Version = "2.0"

	modified, err := xml.MarshalIndent(&myapp, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", modified)
}
