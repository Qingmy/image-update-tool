package docker

import (
	"fmt"
	"image-update-tool/internal/config"
	"image-update-tool/internal/flags"
	"os/exec"
	"path/filepath"
)

func Stop(config *config.Config, service flags.ServiceType) (bool, error) {
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

	ok, err := executeComposeDownCommand(path)
	if err != nil {
		return false, fmt.Errorf("停止%s旧服务失败: %s", service.String(), err)
	}
	return ok, nil
}

func executeComposeDownCommand(path string) (bool, error) {
	if path == "" {
		return false, fmt.Errorf("未指定 compose 路径")
	}
	files, err := filepath.Glob(filepath.Join(path, "*.yml"))
	if err != nil || len(files) == 0 {
		return false, fmt.Errorf("找不到任何 yml 文件")
	}

	args := []string{"down"}
	for _, f := range files {
		args = append([]string{"-f", f}, args...)
	}
	cmd := exec.Command("docker-compose", args...)
	cmd.Dir = path
	output, err := cmd.CombinedOutput()
	fmt.Println("命令输出:\n", string(output))
	if err != nil {
		return false, err
	}
	return true, nil
}
