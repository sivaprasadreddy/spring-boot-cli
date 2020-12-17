package app

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"os"
)

func getAnswers() (GeneratorConfig, error) {
	answers := GeneratorConfig{}
	metadataAnswers, err := getMetadataAnswers()
	if err != nil {
		fmt.Println(err.Error())
		return answers, err
	}
	answers.Metadata = metadataAnswers

	optionsAnswers, err := getOptionsAnswers()
	if err != nil {
		fmt.Println(err.Error())
		return answers, err
	}
	answers.Options = optionsAnswers
	return answers, nil
}

func getMetadataAnswers() (ProjectMetadata, error) {
	var questions = []*survey.Question{
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
		{
			Name: "BuildTool",
			Prompt: &survey.Select{
				Message: "Choose Build Tool:",
				Options: []string{"gradle", "maven"},
				Default: "gradle",
			},
		},
	}
	answers := ProjectMetadata{}
	err := survey.Ask(questions, &answers)
	if err == terminal.InterruptErr {
		os.Exit(0)
	} else if err != nil {
		fmt.Println(err.Error())
		return answers, err
	}
	return answers, nil
}

func getOptionsAnswers() (Options, error) {
	answers, err := getDBOptionsAnswers()
	/*
		var questions = []*survey.Question{
			{
				Name: "Features",
				Prompt: &survey.MultiSelect{
					Message: "What additional features you want to add?",
					Options: []string{"TestContainers"},
					Default: []string{"TestContainers"},
				},
			},
		}
		err := survey.Ask(questions, &answers)
	*/
	if err == terminal.InterruptErr {
		os.Exit(0)
	} else if err != nil {
		fmt.Println(err.Error())
		return answers, err
	}
	return answers, nil
}

func getDBOptionsAnswers() (Options, error) {
	useRdbms := true
	prompt := &survey.Confirm{
		Message: "Do you want to use Relational Database?",
		Default: true,
	}
	_ = survey.AskOne(prompt, &useRdbms)
	if !useRdbms {
		return Options{
			UseRDBMS:        false,
			ProdDbType:      "none",
			TestDbType:      "none",
			DBMigrationType: "none",
		}, nil
	}

	answers := Options{UseRDBMS: true}
	var questions = []*survey.Question{
		{
			Name: "ProdDbType",
			Prompt: &survey.Select{
				Message: "Choose Production Database Type:",
				Options: []string{"postgresql", "mysql", "mariadb"},
				Default: "postgresql",
			},
		},
		{
			Name: "TestDbType",
			Prompt: &survey.Select{
				Message: "Choose Test Database Type:",
				Options: []string{"h2", "hsqldb"},
				Default: "h2",
			},
		},
		{
			Name: "DBMigrationType",
			Prompt: &survey.Select{
				Message: "Choose DB Migration Tool:",
				Options: []string{"flyway", "liquibase", "none"},
				Default: "flyway",
			},
		},
	}
	err := survey.Ask(questions, &answers)
	if err == terminal.InterruptErr {
		os.Exit(0)
	} else if err != nil {
		fmt.Println(err.Error())
		return answers, err
	}
	return answers, nil
}
