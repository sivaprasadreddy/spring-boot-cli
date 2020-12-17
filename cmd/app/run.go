package app

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/sivaprasadreddy/spring-boot-cli/cmd/common"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

const VERSION = "v0.0.4"
const SbCliBaseUri = "https://github.com/sivaprasadreddy/spring-boot-cli/releases/download/"

func Run() {
	downloadTemplates()
	templateDir := getTemplatesDir()
	answers, err := getAnswers()
	common.PanicIfErr(err)
	//fmt.Printf("%#v\n", answers)

	p := ProjectGenerator{
		templatesDir: templateDir,
		targetDir:    answers.Metadata.ApplicationName,
		config:       answers,
	}
	if err := p.GenerateProject(); err != nil {
		log.Fatal(err)
	} else {
		file, _ := json.MarshalIndent(answers, "", " ")
		_ = ioutil.WriteFile(answers.Metadata.ApplicationName+"/.spring-boot-cli.json", file, 0644)
		fmt.Println("Project generated successfully")
	}
}

func getTemplatesDir() string {
	home, err := homedir.Dir()
	common.PanicIfErr(err)
	templateDir := home + "/.sbcli/" + VERSION + "/templates"
	return templateDir
}

func downloadTemplates() {
	home, err := homedir.Dir()
	common.PanicIfErr(err)
	if templatesExists() {
		return
	}
	sbcliBaseDir := home + "/.sbcli"
	sbCliCurrentVersionDir := sbcliBaseDir + "/" + VERSION
	archivesDir := sbcliBaseDir + "/archives"
	osType := runtime.GOOS
	filename := "spring-boot-cli-" + VERSION + "_" + osType + "_amd64.zip"
	err = common.DownloadFile(archivesDir+"/"+filename, SbCliBaseUri+VERSION+"/spring-boot-cli_"+osType+"_amd64.zip")
	common.PanicIfErr(err)
	_, err = common.Unzip(archivesDir+"/"+filename, sbCliCurrentVersionDir)
	common.PanicIfErr(err)
	_, err = common.Unzip(sbCliCurrentVersionDir+"/templates.zip", sbCliCurrentVersionDir)
	common.PanicIfErr(err)
	_ = os.Remove(sbCliCurrentVersionDir + "/templates.zip")
	//_ = os.Remove(sbCliCurrentVersionDir+"/spring-boot-cli")
}

func templatesExists() bool {
	home, err := homedir.Dir()
	common.PanicIfErr(err)
	ok, err := common.IsDirEmpty(home + "/.sbcli/" + VERSION + "/templates")
	return err == nil && !ok
}
