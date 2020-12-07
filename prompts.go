package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
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
			Name:      "ApplicationName",
			Prompt:    &survey.Input{Message: "What is ApplicationName?"},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name:   "GroupID",
			Prompt: &survey.Input{Message: "What is GroupID?"},
		},
		{
			Name:   "ArtifactID",
			Prompt: &survey.Input{Message: "What is ArtifactID?"},
		},
		{
			Name:   "ApplicationVersion",
			Prompt: &survey.Input{Message: "What is Application Version?"},
		},
	}
	answers := ProjectMetadata{}
	err := survey.Ask(metadataQuestions, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return answers, err
	}
	return answers, nil
}