package sender

import (
	"github.com/polidog/go-nowplaying/config"
	"github.com/shkh/lastfm-go/lastfm"
	"time"
	"github.com/polidog/go-nowplaying/track"
)

type Lastfm struct {
	api lastfm.Api
	isLogin bool
}

func (l Lastfm) Send(track track.Track) error {

	if l.isLogin == false {
		return nil
	}

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
	sender := Lastfm{
		isLogin: false,
	}

	api := lastfm.New(config.ApiKey, config.ApiSecret)
	err := api.Login(config.Username, config.Password)
	if err == nil {
		sender.isLogin = true
	}
	sender.api = *api

	return sender

}
