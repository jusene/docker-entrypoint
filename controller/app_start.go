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
}

func NewAppStart() *AppStart {
	yaml := tools.LoadYaml()
	return &AppStart{
		yaml: yaml,
	}
}

func (a *AppStart) Start() {
	switch a.yaml.GetString("runningtime") {
	case "java8", "java11":
		a.startJava()
	case "node":
		a.startNode()
	default:
		fmt.Println("unknow runningtime! please check app.yaml")
		os.Exit(2)
	}
}

func (a *AppStart) startJava() {
	app := a.yaml.GetString("app")
	c := utils.NewConsul()
	k, _ := c.GetKV(app)
	if len(string(k)) == 0 {
		args := strings.Split(conf.JAVA_OPTS, " ")
		args = append(args, app+".jar")
		cmd := exec.Command("java", args...)
		tools.CmdStreamOut(cmd)
	} else {
		v := viper.New()
		v.SetConfigType("yaml")
		v.ReadConfig(bytes.NewReader(k))
		args := strings.Split(v.GetString("java_opts"), " ")
		args = append(args,  app+".jar")
		cmd := exec.Command("java", args...)
		tools.CmdStreamOut(cmd)
	}
}

func (a *AppStart) startNode() {
	args := strings.Split(a.yaml.GetString("entrypoint"), " ")
	cmd := exec.Command("exec", args...)
	tools.CmdStreamOut(cmd)
}
