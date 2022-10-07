package main

import "encoding/xml"

type Root struct {
	XMLName xml.Name `xml:"root"`
	Student []struct {
		ID       string `xml:"id,attr"`
		Name     string `xml:"name"`
		Surname  string `xml:"surname"`
		Gpa      string `xml:"gpa"`
		TermA    string `xml:"termA"`
		Address  string `xml:"address"`
		Postcode string `xml:"postcode"`
		City     string `xml:"city"`
	} `xml:"student"`
}
