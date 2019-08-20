package config

import "github.com/spf13/viper"

type Config interface {
	Servers() []ServerConfig
	Db() DbConfig
}
type ServerConfig interface {
	Host() string
	Port() int
}
type DbConfig interface {
	Host() string
	Port() int
	Database() string
	User() string
	Password() string
}

type config struct {
	servers []serverConfig
}
func (c *config) Servers() []ServerConfig {
	return c.servers[:]
}

type serverConfig struct {
	 host string
	 port int
}

type dbConfig struct {
	host string
	port int
	database string
	user string
	password string
}

func newConfig(viper *viper.Viper) Config {
	settings := viper.Get
}

func getDefaultConfig
