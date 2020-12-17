package app

import (
	"github.com/sivaprasadreddy/spring-boot-cli/cmd/common"
	"text/template"
)

func (p ProjectGenerator) generateConfig() {
	paths := []string{
		p.templatesDir + "/Jenkinsfile.gotmpl",
		p.templatesDir + "/Dockerfile.gotmpl",
		p.templatesDir + "/ci/travis/.travis.yml.gotmpl",
		p.templatesDir + "/ci/github/workflows/dev.yml.gotmpl",
		p.templatesDir + "/ci/github/workflows/master.yml.gotmpl",
		p.templatesDir + "/docker-compose.yml.gotmpl",
		p.templatesDir + "/sonar-project.properties.gotmpl",
		p.templatesDir + "/lombok.config.gotmpl",
		p.templatesDir + "/.editorconfig.gotmpl",
	}
	templateMap := map[string]string{
		"Jenkinsfile.gotmpl":              "Jenkinsfile",
		"Dockerfile.gotmpl":               "Dockerfile",
		".travis.yml.gotmpl":              ".travis.yml",
		"dev.yml.gotmpl":                  ".github/workflows/dev.yml",
		"master.yml.gotmpl":               ".github/workflows/master.yml",
		"docker-compose.yml.gotmpl":       "docker/docker-compose.yml",
		"sonar-project.properties.gotmpl": "sonar-project.properties",
		"lombok.config.gotmpl":            "lombok.config",
		".editorconfig.gotmpl":            ".editorconfig",
	}
	p.generateConfigCode(paths, templateMap)

	err := common.CopyDir(p.templatesDir+"/build-config", p.targetDir+"/build-config")
	if err != nil {
		panic(err)
	}
}

func (p ProjectGenerator) generateConfigCode(paths []string, templateMap map[string]string) {
	rootDir := p.targetDir + "/"
	templates, err := template.ParseFiles(paths...)
	if err != nil {
		panic(err)
	}
	for tmpl, filePath := range templateMap {
		p.createFromTemplate(templates, tmpl, rootDir+filePath)
	}
}
