package docker

import (
	"fmt"
	"image-update-tool/internal/config"
	"image-update-tool/internal/flags"
	"os/exec"
	"path/filepath"

	"github.com/docker/docker/client"
)

func CreateNewContainer(cli *client.Client, config *config.Config, service flags.ServiceType, imageName string) (bool, error) {
	_, err := UpdateComposeFile(config, service, imageName)
	if err != nil {
		return false, fmt.Errorf("创建新容器失败：%v", err)
	}
	var path string
	switch service {
	case flags.EmrWisdomServer:
		path = config.EmrWisdom
	case flags.EmrWisdomSync:
		path = config.EmrWisdom
	case flags.EmrWisdomWebui:
		path = config.EmrWisdomWebUi
	case flags.Mysql:
		path = config.Mysql
	case flags.Redis:
		path = config.Redis
	default:
		return false, fmt.Errorf("未知服务类型：%v", service)
	}
	_, err = executeComposeUpCommand(path)
	if err != nil {
		return false, fmt.Errorf("创建新容器失败：%v", err)
	}
	fmt.Println("✅ 启动容器成功")
	return true, nil
}

func executeComposeUpCommand(path string) (bool, error) {
	if path == "" {
		return false, fmt.Errorf("未指定 compose 路径")
	}
	files, err := filepath.Glob(filepath.Join(path, "*.yml"))
	if err != nil || len(files) == 0 {
		return false, fmt.Errorf("找不到任何 yml 文件")
	}

	args := []string{"up"}
	for _, f := range files {
		args = append([]string{"-f", f}, args...)
	}
	cmd := exec.Command("docker-compose", args...)
	cmd.Dir = path
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	if err != nil {
		return false, err
	}
	return true, nil
}
