// 日志相关配置结构

package config

// Logger 读取yaml文件中的logger信息
type Logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"showLine"`       //是否显示行号
	LogInConsole bool   `yaml:"log_in_console"` //是否显示打印路径
}
