package docker

import (
	"fmt"
	"image-update-tool/internal/config"
	"image-update-tool/internal/flags"

	"github.com/docker/docker/client"
)

func CreateNewContainer(cli *client.Client, config *config.Config, service flags.ServiceType, imageName string) (bool, error) {
	_, err := UpdateComposeFile(config, service, imageName)
	if err != nil {
		return false, fmt.Errorf("创建新容器失败：%v", err)
	}
	return true, nil
}
