package app

import (
	"github.com/sivaprasadreddy/spring-boot-cli/cmd/common"
)

func (p ProjectGenerator) generateMavenConfig() {
	err := common.CopyFile(p.templatesDir+"/maven/gitignore", p.targetDir+"/.gitignore")
	if err != nil {
		panic(err)
	}
	err = common.CopyDir(p.templatesDir+"/maven/wrapper", p.targetDir)
	if err != nil {
		panic(err)
	}

	paths := []string{
		p.templatesDir + "/maven/pom.xml.gotmpl",
	}
	templateMap := map[string]string{
		"pom.xml.gotmpl": "pom.xml",
	}
	p.generateConfigCode(paths, templateMap)
}
