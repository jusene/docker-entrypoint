package controller

import (
	"bytes"
	"docker-entrypoint/conf"
	"docker-entrypoint/utils"
	"fmt"
	"github.com/spf13/viper"
	tools "gitlab.hho-inc.com/devops/flowctl-go/utils"
	"os"
	"os/exec"
)

type AppStart struct {
   yaml *viper.Viper
}

func NewAppStart() *AppStart {
	 yaml := tools.LoadYaml()
	 return  &AppStart{
		 yaml: yaml,
	 }
}

func (a *AppStart) Start() {
	switch  a.yaml.GetString("runningtime") {
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
		cmd := exec.Command("java", conf.JAVA_OPTS, app+".jar")
		tools.CmdStreamOut(cmd)
	} else {
		v := viper.New()
		v.SetConfigType("yaml")
		v.ReadConfig(bytes.NewReader(k))
		cmd := exec.Command("java", v.GetString("java_opts"), " -jar", app+".jar")
		tools.CmdStreamOut(cmd)
	}
}

func (a *AppStart) startNode() {
	cmd := exec.Command(a.yaml.GetString("entrypoint"))
	tools.CmdStreamOut(cmd)
}


