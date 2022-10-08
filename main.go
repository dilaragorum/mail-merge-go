package main

import (
	"embed"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strings"
)

//go:embed template.xml
//go:embed variables.xml
var file embed.FS

func main() {
	variableFileByte, _ := file.ReadFile("variables.xml")

	// Parse xml
	studentInformation := Root{}
	err := xml.Unmarshal(variableFileByte, &studentInformation)
	if err != nil {
		log.Fatalln(err.Error())
	}

	templateFile, _ := file.ReadFile("template.xml")

	CreateMailTemplateXMLFileForEachStudent(string(templateFile), studentInformation.Students)
}

func CreateMailTemplateXMLFileForEachStudent(templateText string, students []Student) {
	for i := range students {
		student := students[i]
		studentTemplate := FullFillTemplateWithStudentInfos(templateText, student)

		fileName := fmt.Sprintf("%s.xml", student.ID)
		studentFile, _ := os.Create(fileName)
		studentFile.WriteString(studentTemplate)
	}
}

func FullFillTemplateWithStudentInfos(templateText string, student Student) string {
	studentTemplate := templateText

	studentTemplate = strings.ReplaceAll(studentTemplate, "<name/>", student.Name)
	studentTemplate = strings.ReplaceAll(studentTemplate, "<surname/>", student.Surname)
	studentTemplate = strings.ReplaceAll(studentTemplate, "<id/>", student.ID)
	studentTemplate = strings.ReplaceAll(studentTemplate, "<address/>", student.Address)
	studentTemplate = strings.ReplaceAll(studentTemplate, "<city/>", student.City)

	if student.Gpa != "" {
		studentTemplate = strings.ReplaceAll(studentTemplate, "<gpa/>", student.Gpa)
	}

	if student.Postcode != "" {
		studentTemplate = strings.ReplaceAll(studentTemplate, "<postcode/>", student.Postcode)
	}

	return studentTemplate
}
