package server

type Config struct {
	Host string `mapstructure:"host"`
	Post string `mapstructure:"port"`
}
