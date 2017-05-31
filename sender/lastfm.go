package sender

import (
	"github.com/polidog/go-itunes"
	"github.com/shkh/lastfm-go/lastfm"
	"github.com/polidog/go-nowplaying/config"
	"time"
)

type Lastfm struct {
	api lastfm.Api
}

func (l Lastfm) Send(track itunes.Track) error {
	p := lastfm.P{"artist": track.Artist, "track": track.Name, "album": track.Album}
	_, err := l.api.Track.UpdateNowPlaying(p)
	if err != nil {
		return err
	}

	start := time.Now().Unix()
	time.Sleep(15 * time.Second)
	p["timestamp"] = start
	_, err = l.api.Track.Scrobble(p)
	if err != nil {
		return err
	}

	return nil
}

func NewLastfmSender(config config.Lastfm) Sender {
	api := lastfm.New(config.ApiKey, config.ApiSecret)
	err := api.Login(config.Username, config.Password)
	if err != nil {
		panic(err)
	}

	return Lastfm{
		api: *api,
	}
}