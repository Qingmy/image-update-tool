package flags

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"os"
)

type ServiceType int

const (
	EmrWisdomServer ServiceType = iota
	EmrWisdomSync
	EmrWisdomWebui
	Redis
	Mysql
)

var ServiceString = map[ServiceType]string{
	EmrWisdomServer: "emr-wisdom-server",
	EmrWisdomSync:   "emr-wisdom-sync",
	EmrWisdomWebui:  "emr-wisdom-webui",
	Redis:           "redis",
	Mysql:           "mysql",
}

func (st ServiceType) String() string {
	return ServiceString[st]
}

type flags struct {
	Service    ServiceType
	ImagePath  string
	ConfigPath string
}

var validServices = map[string]bool{
	"emr-wisdom-server": true,
	"emr-wisdom-sync":   true,
	"emr-wisdom-webui":  true,
	"redis":             true,
	"mysql":             true,
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func exitWithError(msg string) {
	fmt.Fprintln(os.Stderr, "错误：", msg)
	os.Exit(1)
}

func ParseFlags() *flags {
	flags := &flags{}
	var serviceName string
	flag.StringVar(&serviceName, "service", "", "准备更新的服务")
	flag.StringVar(&serviceName, "S", "", "准备更新的服务(简写)")
	flag.StringVar(&flags.ImagePath, "path", "", "更新镜像路径")
	flag.StringVar(&flags.ImagePath, "P", "", "更新镜像路径(简写)")
	flag.StringVar(&flags.ConfigPath, "config", "./update.yml", "配置文件路径")
	flag.StringVar(&flags.ConfigPath, "C", "./update.yml", "配置文件路径")
	flag.Parse()

	//校验 service 是否是合法的服务
	if serviceName == "" {
		exitWithError("错误：--service 参数是必填的\n")
	}
	if !validServices[serviceName] {
		tpl := template.Must(template.New("errInfo").Parse("不存在此服务 {{.}} \n"))
		var buf bytes.Buffer
		err := tpl.Execute(&buf, flags.Service)
		if err != nil {
			exitWithError("模板执行失败")
		}
		exitWithError(buf.String())
	}
	for st, name := range ServiceString {
		if name == serviceName {
			flags.Service = st
			break
		}
	}
	if flags.ImagePath == "" {
		exitWithError("错误：--path 参数是必填的")
	}
	if !pathExists(flags.ImagePath) {
		tpl := template.Must(template.New("errInfo").Parse("路径不存在：{{.}}\n"))
		var buf bytes.Buffer
		err := tpl.Execute(&buf, flags.ImagePath)
		if err != nil {
			exitWithError("模板执行失败")
		}
		exitWithError(buf.String())
	}
	return flags
}
