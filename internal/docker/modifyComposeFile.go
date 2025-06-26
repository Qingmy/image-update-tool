package docker

import (
	"fmt"
	"image-update-tool/internal/config"
	"image-update-tool/internal/flags"
	"os"

	"gopkg.in/yaml.v3"
)

func UpdateComposeFile(config *config.Config, service flags.ServiceType, imageName string) (bool, error) {
	switch service {
	case flags.EmrWisdomServer:
		_, err := modifyComposeFile(config.EmrWisdom+"/emr-wisdom.yml", imageName, flags.EmrWisdomServer)
		if err != nil {
			return false, fmt.Errorf("更新compose文件失败->%v", err)
		}
	case flags.EmrWisdomSync:
		_, err := modifyComposeFile(config.EmrWisdom+"/emr-wisdom.yml", imageName, flags.EmrWisdomSync)
		if err != nil {
			return false, fmt.Errorf("更新compose文件失败->%v", err)
		}
	case flags.EmrWisdomWebui:
		_, err := modifyComposeFile(config.EmrWisdomWebUi+"/emr-wisdom-webui.yml", imageName, flags.EmrWisdomWebui)
		if err != nil {
			return false, fmt.Errorf("更新compose文件失败->%v", err)
		}
	case flags.Mysql:
		_, err := modifyComposeFile(config.Mysql+"/mysql.yml", imageName, flags.Mysql)
		if err != nil {
			return false, fmt.Errorf("更新compose文件失败->%v", err)
		}
	case flags.Redis:
		_, err := modifyComposeFile(config.Redis+"/redis.yml", imageName, flags.Redis)
		if err != nil {
			return false, fmt.Errorf("更新compose文件失败->%v", err)
		}
	default:
		return false, fmt.Errorf("未知服务类型->%v", service)
	}
	return true, nil
}

func modifyComposeFile(filePath string, imageName string, service flags.ServiceType) (bool, error) {
	serviceList := map[flags.ServiceType]string{
		flags.EmrWisdomServer: "api",
		flags.EmrWisdomSync:   "sync",
		flags.EmrWisdomWebui:  "frontend",
		flags.Mysql:           "db",
		flags.Redis:           "redis",
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return false, fmt.Errorf("读取文件失败")
	}
	var compose map[string]interface{}
	if err := yaml.Unmarshal(data, &compose); err != nil {
		return false, fmt.Errorf("用map解码yml文件失败")
	}
	services := compose["services"].(map[string]interface{})
	item := services[serviceList[service]].(map[string]interface{})
	item["image"] = imageName
	newData, err := yaml.Marshal(compose)
	if err != nil {
		return false, fmt.Errorf("编码回yml文件失败")
	}
	if err := os.WriteFile(filePath, newData, 0644); err != nil {
		return false, fmt.Errorf("写回原路径失败")
	}
	fmt.Println("✅ compose.yml 文件已成功修改")
	return true, nil
}
