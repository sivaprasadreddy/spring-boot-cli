package app

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"os"
)

func getAnswers() (GeneratorConfig, error) {
	metadataAnswers, err := getMetadataAnswers()
	answers := GeneratorConfig{
		Metadata: metadataAnswers,
	}
	if err != nil {
		fmt.Println(err.Error())
		return answers, err
	}
	return answers, nil
}

func getMetadataAnswers() (ProjectMetadata, error) {
	var metadataQuestions = []*survey.Question{
		{
			Name: "ApplicationName",
			Prompt: &survey.Input{
				Message: "What is ApplicationName?",
				Help:    "Name of your application",
				Default: "myapp",
			},
			Validate: survey.Required,
		},
		{
			Name: "GroupID",
			Prompt: &survey.Input{
				Message: "What is GroupID?",
				Default: "com.mycompany",
			},
			Validate: func(val interface{}) error {
				if str, ok := val.(string); !ok || len(str) < 1 {
					return errors.New("invalid groupId")
				}
				return nil
			},
		},
		{
			Name: "ArtifactID",
			Prompt: &survey.Input{
				Message: "What is ArtifactID?",
				Default: "myapp",
			},
			Validate: survey.Required,
		},
		{
			Name: "ApplicationVersion",
			Prompt: &survey.Input{
				Message: "What is Application Version?",
				Default: "1.0.0-SNAPSHOT",
			},
			Validate: survey.Required,
		},
		{
			Name: "BasePackage",
			Prompt: &survey.Input{
				Message: "What is base package?",
				Help:    "Base package name",
				Default: "com.mycompany.myapp",
			},
			Validate: survey.Required,
		},
	}
	answers := ProjectMetadata{}
	err := survey.Ask(metadataQuestions, &answers)
	if err == terminal.InterruptErr {
		//fmt.Println("interrupted")
		os.Exit(0)
	} else if err != nil {
		fmt.Println(err.Error())
		return answers, err
	}
	return answers, nil
}
