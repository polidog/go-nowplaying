# go-nowplaying

go-nowplyaing is scrobble your local Sonos queue to a Slack channel And Lastfm.
Supported by mac only.

![](/screenshot.png)

## using

```
$ go install github.com/polidog/go-nowplaying
```

Create config.toml.

```
// ~/.nowplaying.toml

[slack]
token = "your token"
channel = "#general"

[lastfm]
api_key = "your api key"
api_secret = "youar api secret"
username = "username"
password = "password"
```


```
$ go-nowplaying
```

