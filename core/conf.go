package core

import (
	"QAComm/config"
	"QAComm/global"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const ConfigFile = "settings.yaml"

// InitConf 读取yaml文件的配置
func InitConf() {

	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamFile load Init success.")
	global.Config = c //存放在global目录下的全局变量里
}

func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(ConfigFile, byteData, 0777)
	if err != nil {
		return err
	}
	global.Log.Info("配置文件修改成功")
	return nil
}
