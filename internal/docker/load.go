package docker

import (
	"archive/tar"
	"context"
	"fmt"
	"image-update-tool/internal/utils"
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
	stop := utils.StartSpinner("⏳ 正在加载更新镜像")
	// 加载镜像
	response, err := cli.ImageLoad(ctx, tarFile)
	if err != nil {
		return false, fmt.Errorf("加载镜像失败: %w", err)
	}
	defer response.Body.Close()
	stop()
	fmt.Println("✅ 加载更新镜像成功")
	return true, nil
}
