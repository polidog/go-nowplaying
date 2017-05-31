package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Slack  Slack
	Lastfm Lastfm
	isLoad bool
}

func (c *Config) Load(filename string) error {
	if c.isLoad == false {
		_, err := toml.DecodeFile(filename, &c)
		if err != nil {
			return err
		}
		c.isLoad = true
	}
	return nil
}

func NewConfig(filename string) Config {
	config := Config{
		isLoad: false,
	}
	config.Load(filename)
	return config
}
