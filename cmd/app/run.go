package app

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"log"
)

const VERSION = "0.0.3"

func Run() {
	templateDir := getTemplatesDir()
	answers, err := getAnswers()
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%#v\n", answers)

	p := ProjectGenerator{
		templatesDir: templateDir,
		targetDir:    answers.Metadata.ApplicationName,
		config:       answers,
	}
	if err := p.GenerateProject(); err != nil {
		log.Fatal(err)
	} else {
		file, _ := json.MarshalIndent(answers, "", " ")
		_ = ioutil.WriteFile(answers.Metadata.ApplicationName+"/.spring-boot-cli.json", file, 0644)
		fmt.Println("Project generated successfully")
	}
}

func getTemplatesDir() string {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	templateDir := home + "/.sbcli/" + VERSION + "/templates"
	return templateDir
}
