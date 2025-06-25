package docker

import (
	"fmt"

	"github.com/docker/docker/client"
)

func CreateDockerClient() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("创建Docker客户端失败：%w", err)
	}
	fmt.Println("✅ 创建Docker客户端成功")
	return cli, nil
}
