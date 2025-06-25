package docker

import (
	"archive/tar"
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/client"
)

func isTarFileByContent(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()

	tr := tar.NewReader(f)
	_, err = tr.Next()
	return err == nil
}

func LoadImage(cli *client.Client, imagePath string) (bool, error) {
	ctx := context.Background()

	// 判断是否为 tar 文件
	if !isTarFileByContent(imagePath) {
		return false, fmt.Errorf("'%s' 不是有效的 tar 格式", imagePath)
	}

	// 打开文件
	tarFile, err := os.Open(imagePath)
	if err != nil {
		return false, fmt.Errorf("打开 tar 文件失败: %w", err)
	}
	defer tarFile.Close()

	fmt.Println("✅ 正在加载更新镜像")
	// 加载镜像
	response, err := cli.ImageLoad(ctx, tarFile)
	if err != nil {
		return false, fmt.Errorf("加载镜像失败: %w", err)
	}
	defer response.Body.Close()

	fmt.Println("✅ 更新镜像加载完成")
	return true, nil
}
