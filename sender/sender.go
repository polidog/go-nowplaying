package sender

import "github.com/polidog/go-itunes"

type Sender interface {
	Send(track itunes.Track) error
}


