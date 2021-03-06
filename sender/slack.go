package sender

import (
	"bytes"
	"fmt"
	"github.com/polidog/go-nowplaying/config"
	"github.com/polidog/go-nowplaying/track"
	"net/http"
	"net/url"
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
	data.Add("username", fmt.Sprintf("%s - %s", track.Artist, track.Album)) // TODO config
	data.Add("icon_url", track.Image)

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
		return fmt.Sprintf("*%s* from <%s|%s>", track.Name, track.Url, track.Album)
	} else {
		return fmt.Sprintf("*%s* from %s", track.Name, track.Album)
	}
}
