package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	xmlFile, err := os.Open("variables.xml")
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("Successflly open variables.xml")

	byteValue, _ := io.ReadAll(xmlFile)

	studentInformation := Root{}
	xml.Unmarshal(byteValue, &studentInformation)

	fmt.Println(studentInformation.Student)

	if err != nil {
		log.Fatalln(err.Error())
	}

	output, err := xml.MarshalIndent(studentInformation, " ", "   ")
	if err != nil {
		fmt.Println("Error when marshalIndent")
		fmt.Println(err.Error())
	}

	os.WriteFile("template.xml", output, 0)
}
