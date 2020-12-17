package app

import (
	"errors"
	"fmt"
	"github.com/sivaprasadreddy/spring-boot-cli/cmd/common"
	"os"
	"text/template"
)

type ProjectGenerator struct {
	templatesDir string
	targetDir    string
	config       GeneratorConfig
}

func (p ProjectGenerator) GenerateProject() (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovering from panic in GenerateProject, error is: %v \n", r)
			err = errors.New("failed to generate project")
		}
	}()
	p.cleanAndInitProjectDir()
	p.generateBuildConfig()
	p.generateSrcCode()
	p.generateConfig()
	return err
}

func (p ProjectGenerator) cleanAndInitProjectDir() {
	targetDir := p.targetDir
	err := os.RemoveAll(targetDir + "/")
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(targetDir+"/", 0700)
	if err != nil {
		panic(err)
	}
}

func (p ProjectGenerator) generateBuildConfig() {
	if p.config.Metadata.BuildTool == "gradle" {
		p.generateGradleConfig()
	} else {
		p.generateMavenConfig()
	}
}

func (p ProjectGenerator) createFromTemplate(templates *template.Template, templateName, filename string) {
	f := common.CreateFile(filename)
	err := templates.ExecuteTemplate(f, templateName, p.config)
	if err != nil {
		panic(err)
	}
}
