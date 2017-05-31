package main

import (
	"github.com/polidog/go-itunes"
	"github.com/polidog/go-nowplaying/config"
	"github.com/polidog/go-nowplaying/sender"
	"os"
)

func main() {
	filename := os.Args[1]

	config := config.NewConfig(filename)
	watcher := itunes.NewWatcher(300)
	slack := sender.NesSlackSender(config.Slack.Token, config.Slack.Channel)
	lastfm := sender.NewLastfmSender(config.Lastfm)
	for {
		if watcher.Watch() {
			go slack.Send(watcher.Track)
			go lastfm.Send(watcher.Track)
		}
	}
}
