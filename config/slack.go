package config

type Slack struct {
	Token string `toml:"token"`
	Channel string `toml:"channel"`
}
