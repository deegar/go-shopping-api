//Author: Ryan SU
//Email: yuansu.china.work@gmail.com

package config

type SysSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
	Host string `mapstructure:"host" json:"host"`
	Port int32  `mapstructure:"port" json:"port"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"signing_key" json:"signing-key"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int32  `mapstructure:"port" json:"port"`
}

type ServerConfig struct {
	Name         string   `mapstructure:"name" json:"name"`
	Host         string   `mapstructure:"host" json:"host"`
	Port         int32    `mapstructure:"port" json:"port"`
	Lang         string   `mapstructure:"lang" json:"lang"`
	Tags         []string `mapstructure:"tags" json:"tags"`
	SysSrvConfig `mapstructure:"sys_srv" json:"sys-srv"`
	JWTConfig    `mapstructure:"jwt" json:"jwt"`
	ConsulConfig `mapstructure:"consul" json:"consul"`
}

type NacosConfig struct {
	Host              string            `mapstructure:"host"`
	Port              uint64            `mapstructure:"port"`
	User              string            `mapstructure:"user"`
	Password          string            `mapstructure:"password"`
	Namespace         string            `mapstructure:"namespace"`
	DataId            string            `mapstructure:"data_id"`
	Group             string            `mapstructure:"group"`
	NacosClientConfig NacosClientConfig `mapstructure:"client"`
}

type NacosClientConfig struct {
	LogDir   string `mapstructure:"log_dir"`
	CacheDir string `mapstructure:"cache_dir"`
}
