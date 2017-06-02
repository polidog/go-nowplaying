package sender

import (
	"github.com/polidog/go-nowplaying/track"
)

type Sender interface {
	Send(track track.Track) error
}
