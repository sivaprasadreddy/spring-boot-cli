package app

import (
	"github.com/sivaprasadreddy/spring-boot-cli/cmd/common"
)

func (p ProjectGenerator) generateMavenConfig() {
	err := common.CopyFile(p.templatesDir+"/maven/gitignore", p.targetDir+"/.gitignore")
	common.PanicIfErr(err)
	err = common.CopyDir(p.templatesDir+"/maven/wrapper", p.targetDir)
	common.PanicIfErr(err)

	paths := []string{
		p.templatesDir + "/maven/pom.xml.gotmpl",
	}
	templateMap := map[string]string{
		"pom.xml.gotmpl": "pom.xml",
	}
	p.generateConfigCode(paths, templateMap)
}
