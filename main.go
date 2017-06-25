package nowplaying

import (
	"github.com/polidog/go-itunes"
	"github.com/polidog/go-nowplaying/config"
	"github.com/polidog/go-nowplaying/sender"
	"github.com/polidog/go-nowplaying/track"
	"log"
	"os"

)


func main() {
	var filename string


	if len(os.Args) != 2 {
		filename = homeDir() + "/.nowplaying.toml"
	} else {
		filename = os.Args[1]
	}

	config, err := config.NewConfig(filename)
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