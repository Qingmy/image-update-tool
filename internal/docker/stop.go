package docker

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func Stop(cli *client.Client, service string, timeoutSeconds int) (bool, error) {
	ctx := context.Background()
	if timeoutSeconds <= 0 {
		timeoutSeconds = 10
	}
	timeout := &timeoutSeconds
	containers, err := cli.ContainerList(ctx, container.ListOptions{
		All: true,
	})
	if err != nil {
		return false, fmt.Errorf("获取容器列表失败: %w", err)
	}
	var containerID string
	found := false
	for _, c := range containers {
		for _, name := range c.Names {
			if strings.TrimPrefix(name, "/") == service {
				containerID = c.ID
				found = true
				break
			}
		}
	}
	if !found {
		return false, fmt.Errorf("未找到名为 '%s' 的容器", service)
	}
	err = cli.ContainerStop(ctx, containerID, container.StopOptions{
		Timeout: timeout,
	})
	if err != nil {
		return false, fmt.Errorf("停止容器失败: %w", err)
	}
	fmt.Println("✅ 停止旧镜像成功")
	return true, nil
}
