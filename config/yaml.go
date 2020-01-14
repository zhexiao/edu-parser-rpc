package config

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
)

type CT_YamlSettings struct {
	Proj struct {
		Release bool `yaml:"release"`
	}

	Qiniu struct {
		AccessKey    string `yaml:"accessKey"`
		SecretKey    string `yaml:"secretKey"`
		Bucket       string `yaml:"bucket"`
		Zone         string `yaml:"zone"`
		Domain       string `yaml:"domain"`
		UploadPrefix string `yaml:"uploadPrefix"`
	}

	Wmf struct {
		Uri string `yaml:"uri"`
	}
}

func ReadYaml() *CT_YamlSettings {
	//读取配置文件
	yamlSettings := new(CT_YamlSettings)

	yamlFile, err := ioutil.ReadFile("./config/settings.yaml")
	if err != nil {
		log.Panicf("配置文件读取失败,err=%s", err)
	}

	err = yaml.Unmarshal(yamlFile, yamlSettings)
	if err != nil {
		log.Fatalf("Yaml解析失败: err=%s", err)
	}

	return yamlSettings
}
