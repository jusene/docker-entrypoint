package controller

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"gitlab.hho-inc.com/devops/docker-entrypoint/conf"
	"gitlab.hho-inc.com/devops/docker-entrypoint/utils"
	tools "gitlab.hho-inc.com/devops/flowctl-go/utils"
	"os"
	"os/exec"
	"strings"
)

type AppStart struct {
	yaml *viper.Viper
	debug bool
}

func NewAppStart(debug bool) *AppStart {
	yaml := tools.LoadYaml()
	return &AppStart{
		yaml: yaml,
		debug: debug,
	}
}

func (a *AppStart) Start() {
	switch a.yaml.GetString("runningtime") {
	case "java8", "java11":
		a.startJava()
	case "node":
		a.startNode()
	case "golang":
		a.startGolang()
	default:
		fmt.Println("unknow runningtime! please check app.yaml")
		os.Exit(2)
	}
}

func (a *AppStart) startJava() {
	app := a.yaml.GetString("app")
	c := utils.NewConsul("/devops/cicd/build/controller")
	k, _ := c.GetKV(app)
	if len(string(k)) == 0 {
		var args []string
		if a.yaml.GetString("runningtime") == "java8" {
			args = strings.Split(conf.JAVA8_OPTS, " ")
		} else if a.yaml.GetString("runningtime") == "java11" {
			args = strings.Split(conf.JAVA11_OPTS, " ")
		}
		args = append(args, app+".jar")
		a.debugPrint(args)
		cmd := exec.Command("java", args...)
		tools.CmdStreamOut(cmd)
	} else {
		v := viper.New()
		v.SetConfigType("yaml")
		v.ReadConfig(bytes.NewReader(k))
		args := strings.Split(v.GetString("java_opts"), " ")
		args = append(args,  app+".jar")
		a.debugPrint(args)
		cmd := exec.Command("java", args...)
		tools.CmdStreamOut(cmd)
	}
}

func (a *AppStart) startNode() {
	args := strings.Split(a.yaml.GetString("entrypoint"), " ")
	cmd := exec.Command(args[0], args[1:]...)
	tools.CmdStreamOut(cmd)
}

func (a *AppStart) startGolang() {
	cmd := exec.Command("./" + a.yaml.GetString("app"))
	tools.CmdStreamOut(cmd)
}

func (a *AppStart) debugPrint(args []string) {
	if a.debug {
		fmt.Println("java", strings.Join(args, " "))
		os.Exit(1)
	}
}