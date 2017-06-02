package main

import (
	"github.com/polidog/go-itunes"
	"github.com/polidog/go-nowplaying/config"
	"github.com/polidog/go-nowplaying/sender"
	"os"
	"log"
	"github.com/polidog/go-nowplaying/track"
	"runtime"
)

func main() {
	var filename string

	if len(os.Args) != 2 {
		filename = homeDir() + "/.nowplaying.toml"
	} else {
		filename = os.Args[1]
	}

	config,err := config.NewConfig(filename)
	if err != nil {
		log.Fatal(err)
		return
	}

	watcher := itunes.NewWatcher(5000)
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

func homeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}