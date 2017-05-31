package sender

import (
	"bytes"
	"fmt"
	"github.com/polidog/go-itunes"
	"net/http"
	"net/url"
	"github.com/polidog/go-nowplaying/config"
)

type Slack struct {
	Token   string
	Channel string
}

var apiUrl = "https://slack.com/api/chat.postMessage"

func (s Slack) Send(track itunes.Track) error {

	if len(s.Token) == 0 || len(s.Channel) == 0 {
		return nil
	}

	data := url.Values{}
	data.Set("token", s.Token)
	data.Add("channel", s.Channel)
	data.Add("username", "NowPlaying") // TODO config
	data.Add("text", "NowPlaying:"+track.Name+" by "+track.Artist+" from "+track.Artist)

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
