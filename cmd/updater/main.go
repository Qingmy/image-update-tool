package main

import (
	"fmt"
	"image-update-tool/internal/config"
	"image-update-tool/internal/docker"
	"image-update-tool/internal/flags"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flags := flags.ParseFlags()
	configs, err := config.ReadYaml(flags.ConfigPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误：%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("开始更新服务 [%s]，镜像路径：%s\n", flags.Service, flags.ImagePath)
	cli, err := docker.CreateDockerClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误：%v\n", err)
		os.Exit(1)
	}
	_, err = docker.LoadImage(cli, flags.ImagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误：%v\n", err)
		os.Exit(1)
	}
	_, err = docker.Stop(configs, flags.Service)
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误：%v\n", err)
		os.Exit(1)
	}
	_, err = docker.CreateNewContainer(cli, configs, flags.Service, strings.TrimSuffix(filepath.Base(flags.ImagePath), ".tar"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误：%v\n", err)
		os.Exit(1)
	}
	fmt.Println("✅ 服务更新成功！")
}
