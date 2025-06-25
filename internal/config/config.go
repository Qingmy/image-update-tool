package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	EmrWisdom      string `yaml:"EmrWisdom"`
	EmrWisdomWebUi string `yaml:"EmrWisdomWebUi"`
	Mysql          string `yaml:"Mysql"`
	Redis          string `yaml:"Redis"`
}

func ReadYaml(configPath string) (*Config, error) {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("配置文件不存在: %s", configPath)
	}
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("解析配置失败: %v", err)
	}
	return &cfg, nil
}
