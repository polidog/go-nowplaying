package sender

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"github.com/polidog/go-nowplaying/config"
	"github.com/polidog/go-nowplaying/track"
)

type Slack struct {
	Token   string
	Channel string
}

var apiUrl = "https://slack.com/api/chat.postMessage"

func (s Slack) Send(track track.Track) error {

	if len(s.Token) == 0 || len(s.Channel) == 0 {
		return nil
	}

	data := url.Values{}
	data.Set("token", s.Token)
	data.Add("channel", s.Channel)
	data.Add("username", "NowPlaying - " + track.Artist) // TODO config

	if len(track.Image) > 0 {
		data.Add("icon_url", track.Image)
	} else {
		data.Add("icon_url", "https://github.com/polidog/go-nowplaying/gopher.png")
	}

	data.Add("text", createText(track))

	client := &http.Client{}
	r, _ := http.NewRequest("POST", fmt.Sprintf("%s", apiUrl), bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	_, err := client.Do(r)
	if err != nil {
		return err
	}
	return nil
}

func NesSlackSender(slack config.Slack) Sender {
	return Slack{
		Token:   slack.Token,
		Channel: slack.Channel,
	}
}

func createText(track track.Track) string {
	if len(track.Url) > 0 {
		return fmt.Sprintf("%s from <%s|%s>",track.Name, track.Url, track.Album)
	} else {
		return fmt.Sprintf("%s from %s", track.Name, track.Album)
	}
}