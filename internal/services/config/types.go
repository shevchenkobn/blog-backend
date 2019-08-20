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
	db dbConfig
}
func (c *config) Servers() []ServerConfig {
	returnServers := make([]ServerConfig, len(c.servers), len(c.servers))
	for i, server := range c.servers {
		returnServers[i] = &server
	}
	return returnServers
}
func (c *config) Db() DbConfig {
	return &c.db
}

type serverConfig struct {
	 host string
	 port int
}
func (c *serverConfig) Host() string {
	return c.host
}
func (c *serverConfig) Port() int {
	return c.port
}

type dbConfig struct {
	host string
	port int
	database string
	user string
	password string
}
func (c *dbConfig) Host() string {
	return c.host
}
func (c *dbConfig) Port() int {
	return c.port
}
func (c *dbConfig) Database() string {
	return c.database
}
func (c *dbConfig) User() string {
	return c.user
}
func (c *dbConfig) Password() string {
	return c.password
}

func newConfig(viper *viper.Viper) (Config, error) {
	config := new(config)
	err := viper.Unmarshal(config)
	return config, err
}
