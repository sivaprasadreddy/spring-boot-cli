package main

import (
	"fmt"
	"log"
)

func main() {
	/*answers, err := getAnswers()
	if err != nil {
		panic(err)
	}*/
	answers := GeneratorConfig{
		Metadata: ProjectMetadata{
			ApplicationName:    "demo-app",
			GroupID:            "com.sivalabs",
			ArtifactID:         "demo-app",
			ApplicationVersion: "1.0.0-SNAPSHOT",
		},
	}

	fmt.Printf("%#v\n", answers)
	//fmt.Println(answers)
	//jsonF, _ := json.Marshal(answers)
	//fmt.Println(string(jsonF))

	if err := GenerateProject(answers); err != nil {
		log.Fatal(err)
	}

}



