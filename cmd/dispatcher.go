package cmd

import (
	"github.com/sivaprasadreddy/spring-boot-cli/cmd/app"
)

func invokeGenerator() {
	app.Run()
}

/*
func invokeGenerator(cmd *cobra.Command, args []string) error {
	generatorType := cmd.Flag("type").Value.String()
	if generatorType == "" {
		answers, err := getAppTypeAnswers()
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		generatorType = answers.AppType
	}

	if strings.EqualFold(generatorType, "microservice") {
		app.Run()
	} else if strings.EqualFold(generatorType, "spring-cloud-config-server") {
		fmt.Println("SpringCloud Config Server - Work In Progress")
	} else {
		fmt.Println("Unknown generator type")
	}
	return nil
}

type GeneratorType struct {
	AppType string
}

func getAppTypeAnswers() (GeneratorType, error) {
	var answers = []*survey.Question{
		{
			Name: "AppType",
			Prompt: &survey.Select{
				Message: "Choose application type:",
				Options: []string{"microservice", "spring-cloud-config-server"},
				Default: "microservice",
			},
		},
	}
	generatorType := GeneratorType{}
	err := survey.Ask(answers, &generatorType)
	if err == terminal.InterruptErr {
		os.Exit(0)
	} else if err != nil {
		fmt.Println(err.Error())
		return generatorType, err
	}
	return generatorType, nil
}
*/
