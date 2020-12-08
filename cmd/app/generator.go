package app

import (
	"errors"
	"fmt"
	"github.com/sivaprasadreddy/spring-boot-cli/cmd/common"
	"os"
	"path/filepath"
	"strings"
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
	p.generateMavenConfig()
	p.generateSrcCode()
	return err
}

func (p ProjectGenerator) generateSrcCode() {
	p.generateSrcMainJavaCode()
	p.generateSrcMainResourcesCode()
	p.generateSrcTestJavaCode()
	p.generateSrcTestResourcesCode()
}

func (p ProjectGenerator) generateSrcMainJavaCode() {
	mainJavaFolder := "src/main/java"
	paths := []string{
		p.templatesDir + "/src/main/java/package/Application.java.gotmpl",
	}
	templateMap := map[string]string{
		"Application.java.gotmpl": "Application.java",
	}
	p.generateSrcJavaCode(mainJavaFolder, paths, templateMap)
}

func (p ProjectGenerator) generateSrcTestJavaCode() {
	testJavaFolder := "src/test/java"
	paths := []string{
		p.templatesDir + "/src/test/java/package/ApplicationTests.java.gotmpl",
	}
	templateMap := map[string]string{
		"ApplicationTests.java.gotmpl": "ApplicationTests.java",
	}
	p.generateSrcJavaCode(testJavaFolder, paths, templateMap)
}

func (p ProjectGenerator) generateSrcMainResourcesCode() {
	mainResFolder := "src/main/resources"
	paths := []string{
		p.templatesDir + "/src/main/resources/application.properties.gotmpl",
	}
	templateMap := map[string]string{
		"application.properties.gotmpl": "application.properties",
	}
	p.generateSrcResourcesCode(mainResFolder, paths, templateMap)
}

func (p ProjectGenerator) generateSrcTestResourcesCode() {
	testResFolder := "src/test/resources"
	paths := []string{
		p.templatesDir + "/src/test/resources/application-test.properties.gotmpl",
	}
	templateMap := map[string]string{
		"application-test.properties.gotmpl": "application-test.properties",
	}
	p.generateSrcResourcesCode(testResFolder, paths, templateMap)
}

func (p ProjectGenerator) generateSrcJavaCode(javaFolder string, paths []string, templateMap map[string]string) {
	basePackagePath := strings.ReplaceAll(p.config.Metadata.BasePackage, ".", "/")
	testJavaRootDir := p.targetDir + "/" + javaFolder + "/" + basePackagePath + "/"
	templates, err := template.ParseFiles(paths...)
	if err != nil {
		panic(err)
	}
	for tmpl, filePath := range templateMap {
		p.createFromTemplate(templates, tmpl, testJavaRootDir+filePath)
	}
}

func (p ProjectGenerator) generateSrcResourcesCode(resFolder string, paths []string, templateMap map[string]string) {
	resRootDir := p.targetDir + "/" + resFolder + "/"
	templates, err := template.ParseFiles(paths...)
	if err != nil {
		panic(err)
	}
	for tmpl, filePath := range templateMap {
		p.createFromTemplate(templates, tmpl, resRootDir+filePath)
	}
}

func (p ProjectGenerator) generateMavenConfig() {
	paths := []string{
		p.templatesDir + "/maven/pom.xml.gotmpl",
	}
	templates, err := template.ParseFiles(paths...)

	err = common.CopyFile(p.templatesDir+"/maven/gitignore", p.targetDir+"/.gitignore")
	if err != nil {
		panic(err)
	}
	err = common.CopyDir(p.templatesDir+"/maven/wrapper", p.targetDir)
	if err != nil {
		panic(err)
	}

	p.createFromTemplate(templates, "pom.xml.gotmpl", p.targetDir+"/pom.xml")
}

func (p ProjectGenerator) createFromTemplate(templates *template.Template, templateName, filename string) {
	f := createFile(filename)
	err := templates.ExecuteTemplate(f, templateName, p.config)
	if err != nil {
		panic(err)
	}
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

func createFile(filePath string) *os.File {
	parent := filepath.Dir(filePath)
	_ = os.MkdirAll(parent, 0700)
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	return f
}
