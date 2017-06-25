package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Slack   Slack
	Lastfm  Lastfm
	Country string `toml:"country"`
	Time	int `toml:"time"`
	isLoad  bool
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

func NewConfig(filename string) (Config, error) {
	config := Config{
		isLoad:  false,
		Country: "US",
		Time: 5000,
	}
	err := config.Load(filename)
	return config, err
}
