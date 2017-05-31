package config

type Lastfm struct {
	ApiKey string `toml:"api_key"`
	ApiSecret string `toml:"api_secret"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}