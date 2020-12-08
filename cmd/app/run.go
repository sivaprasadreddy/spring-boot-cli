package app

import (
	"fmt"
	"log"
)

func Run() {
	answers, err := getAnswers()
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%#v\n", answers)
	p := ProjectGenerator{
		templatesDir: "templates",
		targetDir:    answers.Metadata.ApplicationName,
		config:       answers,
	}
	if err := p.GenerateProject(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Project generated successfully")
	}
}
