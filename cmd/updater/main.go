package main

import (
	"fmt"
	"image-update-tool/internal/docker"
	"image-update-tool/internal/flags"
	"os"
)

func main() {
	flags := flags.ParseFlags()
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
	_, err = docker.Stop(cli, flags.Service, 5)
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误：%v\n", err)
		os.Exit(1)
	}
}
