package app

import (
	"strings"
	"text/template"
)

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
		p.templatesDir + "/src/main/java/package/config/WebMvcConfig.java.gotmpl",
		p.templatesDir + "/src/main/java/package/config/SwaggerConfig.java.gotmpl",
	}
	templateMap := map[string]string{
		"Application.java.gotmpl":   "Application.java",
		"WebMvcConfig.java.gotmpl":  "config/WebMvcConfig.java",
		"SwaggerConfig.java.gotmpl": "config/SwaggerConfig.java",
	}
	p.generateSrcJavaCode(mainJavaFolder, paths, templateMap)
}

func (p ProjectGenerator) generateSrcTestJavaCode() {
	testJavaFolder := "src/test/java"
	paths := []string{
		p.templatesDir + "/src/test/java/package/ApplicationTests.java.gotmpl",
		p.templatesDir + "/src/test/java/package/common/AbstractIntegrationTest.java.gotmpl",
		p.templatesDir + "/src/test/java/package/common/TestContainersInitializer.java.gotmpl",
	}
	templateMap := map[string]string{
		"ApplicationTests.java.gotmpl":          "ApplicationTests.java",
		"AbstractIntegrationTest.java.gotmpl":   "common/AbstractIntegrationTest.java",
		"TestContainersInitializer.java.gotmpl": "common/TestContainersInitializer.java",
	}
	p.generateSrcJavaCode(testJavaFolder, paths, templateMap)
}

func (p ProjectGenerator) generateSrcMainResourcesCode() {
	mainResFolder := "src/main/resources"
	paths := []string{
		p.templatesDir + "/src/main/resources/application.properties.gotmpl",
		p.templatesDir + "/src/main/resources/application-prod.properties.gotmpl",
		p.templatesDir + "/src/main/resources/application-heroku.properties.gotmpl",
	}
	templateMap := map[string]string{
		"application.properties.gotmpl":        "application.properties",
		"application-prod.properties.gotmpl":   "application-prod.properties",
		"application-heroku.properties.gotmpl": "application-heroku.properties",
	}
	if p.config.Options.DBMigrationType == "flyway" {
		paths = append(paths, p.templatesDir+"/src/main/resources/db/migration/flyway/V1__01_init.sql.gotmpl")
		templateMap["V1__01_init.sql.gotmpl"] = "/src/main/resources/db/migration/V1__01_init.sql"
	}
	if p.config.Options.DBMigrationType == "liquibase" {
		paths = append(paths, p.templatesDir+"/src/main/resources/db/migration/liquibase/liquibase-changelog.xml.gotmpl")
		paths = append(paths, p.templatesDir+"/src/main/resources/db/migration/liquibase/changelog/01-init.xml.gotmpl")

		templateMap["liquibase-changelog.xml.gotmpl"] = "/src/main/resources/db/migration/liquibase-changelog.xml"
		templateMap["01-init.xml.gotmpl"] = "/src/main/resources/db/migration/changelog/01-init.xml"
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
