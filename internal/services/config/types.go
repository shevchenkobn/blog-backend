package config

import "github.com/spf13/viper"

type Config interface {
	Servers() []ServerConfig
	Db() DbConfig
	OpenApi() OpenApiConfig
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
type OpenApiConfig interface {
	ConfigPath() string
}

type config struct {
	ServersField []*serverConfig `mapstructure:"servers"`
	DbField      *dbConfig       `mapstructure:"db"`
	OpenApiField *openApiConfig `mapstructure:"openapi"`
}
func (c *config) Servers() []ServerConfig {
	returnServers := make([]ServerConfig, len(c.ServersField), len(c.ServersField))
	for i, server := range c.ServersField {
		returnServers[i] = server
	}
	return returnServers
}
func (c *config) Db() DbConfig {
	return c.DbField
}
func (c *config) OpenApi() OpenApiConfig {
	return c.OpenApiField
}

type serverConfig struct {
	 HostField string `mapstructure:"host"`
	 PortField int    `mapstructure:"port"`
}
func (c *serverConfig) Host() string {
	return c.HostField
}
func (c *serverConfig) Port() int {
	return c.PortField
}

type dbConfig struct {
	HostField     string `mapstructure:"host"`
	PortField     int    `mapstructure:"port"`
	DatabaseField string `mapstructure:"database"`
	UserField     string `mapstructure:"user"`
	PasswordField string `mapstructure:"password"`
}
func (c *dbConfig) Host() string {
	return c.HostField
}
func (c *dbConfig) Port() int {
	return c.PortField
}
func (c *dbConfig) Database() string {
	return c.DatabaseField
}
func (c *dbConfig) User() string {
	return c.UserField
}
func (c *dbConfig) Password() string {
	return c.PasswordField
}

type openApiConfig struct {
	ConfigPathField string `mapstructure:"configPath"`
}
func (c *openApiConfig) ConfigPath() string {
	return c.ConfigPathField
}

func newConfig(viper *viper.Viper) (Config, error) {
	config := new(config)
	err := viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
