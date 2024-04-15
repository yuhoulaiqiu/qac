// 配置文件的总入口

package config

type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
	Jwt    Jwt    `yaml:"jwt"`
	Redis  Redis  `yaml:"redis"`
}
