package config

var Conf *Config

type Config struct {
	Server Server `mapstructure:"server"`
	KV     KeeperVault
}

type Server struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type KeeperVault struct {
	KsmConfig string
}
