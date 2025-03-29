package config

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}

type AliOssConfig struct {
	BucketName string `mapstructure:"bucket_name" json:"bucket_name,omitempty"`
	Region     string `mapstructure:"region" json:"region,omitempty"`
	ObjectName string `mapstructure:"object_name" json:"object_name,omitempty"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type ServerConfig struct {
	Name       string       `mapstructure:"name" json:"name,omitempty"`
	ConsulInfo ConsulConfig `mapstructure:"consul" json:"consul,omitempty"`
	OssInfo    AliOssConfig `mapstructure:"oss" json:"oss,omitempty"`
}
