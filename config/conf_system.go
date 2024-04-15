// 系统相关配置结构

package config

import "fmt"

// System 读取yaml文件中的system信息
type System struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Env     string `yaml:"env"`
	TimeOut int    `yaml:"timeout"`
}

// Addr 拼接url
func (s System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
