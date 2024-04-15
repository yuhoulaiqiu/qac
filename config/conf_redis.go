package config

// Redis 读取yaml文件中的Redis配置
type Redis struct {
	Addr     string `yaml:"addr"`
	Db       int    `yaml:"db"`
	Password string `yaml:"password"`
}
