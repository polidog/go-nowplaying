package nowplaying

import (
	"github.com/polidog/go-nowplaying/config"
	"github.com/polidog/go-itunes"
	"github.com/polidog/go-nowplaying/sender"
	"github.com/polidog/go-nowplaying/track"
)

func Run(config config.Config) {

	watcher := itunes.NewWatcher(config.Time)
	slack := sender.NesSlackSender(config.Slack)
	lastfm := sender.NewLastfmSender(config.Lastfm)

	for {
		if watcher.Watch() {
			t := track.NewTrack(watcher.Track)
			t.GetImageAndUrl(config.Country)
			go slack.Send(t)
			go lastfm.Send(t)
		}
	}

}