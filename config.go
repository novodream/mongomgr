package mongomgr

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	*viper.Viper
	name string
}

func (c *Config) load() error {
	c.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	c.SetEnvKeyReplacer(replacer)

	c.SetConfigName(c.name)
	c.SetConfigType("yaml")
	c.AddConfigPath(".dream")
	if err := c.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error
		} else {
			// Config file was found but another error was produced
			return err
		}
	}
	return nil
}

func load(name string) (*Config, error) {
	c := &Config{viper.New(), name}
	return c, c.load()
}

func loadEnvConfig() (*Config, error) {
	return load("env.yml")
}
