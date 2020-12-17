package app

import "github.com/sivaprasadreddy/spring-boot-cli/cmd/common"

func (p ProjectGenerator) generateGradleConfig() {
	err := common.CopyFile(p.templatesDir+"/gradle/gitignore", p.targetDir+"/.gitignore")
	common.PanicIfErr(err)
	err = common.CopyDir(p.templatesDir+"/gradle/wrapper", p.targetDir)
	common.PanicIfErr(err)

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
