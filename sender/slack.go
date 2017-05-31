package sender

import (
	"github.com/polidog/go-itunes"
	"net/url"
	"net/http"
	"fmt"
	"bytes"
)

type Slack struct {
	Token string
	Channel string
}

var apiUrl = "https://slack.com/api/chat.postMessage"

func (s Slack) Send(track itunes.Track) error {

	data := url.Values{}
	data.Set("token", s.Token)
	data.Add("channel",s.Channel)
	data.Add("username", "NowPlaying") // TODO config
	data.Add("text", "NowPlaying:" + track.Name + " by " + track.Artist + " from " + track.Artist)

	client := &http.Client{}
	r, _ := http.NewRequest("POST",  fmt.Sprintf("%s",apiUrl), bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	_, err := client.Do(r)
	if err != nil {
		return err
	}
	return nil
}

func NesSlackSender(token string, channel string) Sender {
	return Slack{
		Token: token,
		Channel: channel,
	}
}
