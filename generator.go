package main

import (
	"log"
	"os"
	"text/template"
)

func GenerateProject(answers GeneratorConfig) error {
	targetDir := answers.Metadata.ApplicationName

	err := os.RemoveAll(targetDir + "/")
	os.MkdirAll(targetDir + "/", 0700)
	if err != nil {
		return err
	}

	paths := []string{
		"templates/maven/pom.xml.gotmpl",
	}
	templates, err := template.ParseFiles(paths...)

	err = CopyFile("templates/maven/gitignore", targetDir+"/.gitignore")
	if err != nil {
		return err
	}
	err = CopyDir("templates/maven/wrapper", targetDir)
	if err != nil {
		return err
	}

	f, err := os.Create(targetDir + "/pom.xml")
	if err != nil {
		log.Println("create file: ", err)
		return err
	}
	err = templates.ExecuteTemplate(f, "pom.xml.gotmpl", answers)
	if err != nil {
		return err
	}
	return nil
}
