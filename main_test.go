package main

import (
	_ "embed"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

//go:embed template.xml
var templateText string

func Test_Create_Mail_TemplateXMLFile_For_Each_Student(t *testing.T) {
	student1 := Student{
		ID:       "1",
		Name:     "James",
		Surname:  "West",
		Gpa:      "3.00",
		TermA:    "a",
		Address:  "2656 South Loop West Suite 380 Houston TX 77054 USA",
		Postcode: "77054",
		City:     "Houston",
	}

	var students []Student
	students = append(students, student1)

	CreateMailTemplateXMLFileForEachStudent(templateText, students)

	fileName := fmt.Sprintf("%s.xml", student1.ID)
	studentFile, _ := os.Open(fileName)
	studentFileByte, _ := io.ReadAll(studentFile)

	actualText := string(studentFileByte)
	expectedText := "<mail>\n    Dear Student James West,\n    " +
		"This letter is an example to inform you about the content of template file. " +
		"Your student should be 1 and your GPA is to be 3.00. " +
		"Thanks for reading.\n    " +
		"Your address is: 2656 South Loop West Suite 380 Houston TX 77054 USA 77054 Houston\n</mail>"

	assert.Equal(t, expectedText, actualText)
}

func Test_FulFill_Template_With_Student_Infos(t *testing.T) {
	t.Run("Success Case - When all information is given", func(t *testing.T) {
		student1 := Student{
			ID:       "1",
			Name:     "James",
			Surname:  "West",
			Gpa:      "3.00",
			TermA:    "a",
			Address:  "2656 South Loop West Suite 380 Houston TX 77054 USA",
			Postcode: "77054",
			City:     "Houston",
		}

		actualResult := FullFillTemplateWithStudentInfos(templateText, student1)
		expectedResult := "<mail>\n    Dear Student James West,\n    " +
			"This letter is an example to inform you about the content of template file. " +
			"Your student should be 1 and your GPA is to be 3.00. Thanks for reading.\n    " +
			"Your address is: 2656 South Loop West Suite 380 Houston TX 77054 USA 77054 Houston\n</mail>"

		assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("Success Case - When gpa and postcode information is missed", func(t *testing.T) {
		student1 := Student{
			ID:       "1",
			Name:     "James",
			Surname:  "West",
			Gpa:      "",
			TermA:    "a",
			Address:  "2656 South Loop West Suite 380 Houston TX 77054 USA",
			Postcode: "",
			City:     "Houston",
		}

		actualResult := FullFillTemplateWithStudentInfos(templateText, student1)
		expectedResult := "<mail>\n    Dear Student James West,\n    " +
			"This letter is an example to inform you about the content of template file. " +
			"Your student should be 1 and your GPA is to be <gpa/>. Thanks for reading.\n    " +
			"Your address is: 2656 South Loop West Suite 380 Houston TX 77054 USA <postcode/> Houston\n</mail>"

		assert.Equal(t, expectedResult, actualResult)
	})
}
