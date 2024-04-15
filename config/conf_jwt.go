package config

type Jwt struct {
	Secret  string `json:"secret" yaml:"secret"`   //密钥
	Expires int    `json:"expires" yaml:"expires"` //过期时间
	Issuer  string `yaml:"issuer" json:"issuer"`   //颁发人
}
