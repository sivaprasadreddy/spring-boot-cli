package app

import "github.com/sivaprasadreddy/spring-boot-cli/cmd/common"

func (p ProjectGenerator) generateGradleConfig() {
	err := common.CopyFile(p.templatesDir+"/gradle/gitignore", p.targetDir+"/.gitignore")
	if err != nil {
		panic(err)
	}
	err = common.CopyDir(p.templatesDir+"/gradle/wrapper", p.targetDir)
	if err != nil {
		panic(err)
	}

	paths := []string{
		p.templatesDir + "/gradle/build.gradle.gotmpl",
		p.templatesDir + "/gradle/settings.gradle.gotmpl",
	}

	templateMap := map[string]string{
		"build.gradle.gotmpl":    "build.gradle",
		"settings.gradle.gotmpl": "settings.gradle",
	}
	p.generateConfigCode(paths, templateMap)
}
