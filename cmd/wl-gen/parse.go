package main

import "encoding/xml"

type Protocol struct {
	Name       string      `xml:"name,attr"`
	Copyright  string      `xml:"copyright"`
	Interfaces []Interface `xml:"interface"`
}

type Interface struct {
	Name  string `xml:"name,attr"`
	Enums []Enum `xml:"enum"`
}

type Enum struct {
	Name        string  `xml:"name,attr"`
	Description string  `xml:"description"`
	Entries     []Entry `xml:"entry"`
}

type Entry struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

func parseProtocolXML(input []byte) (Protocol, error) {
	var protocol Protocol

	err := xml.Unmarshal(input, &protocol)
	if err != nil {
		return protocol, err
	}

	return protocol, nil
}
