package track

import (
	"github.com/polidog/go-itunes"
	"github.com/mattn/itunes-search-api"
)

type Track struct {
	itunes.Track
	Image string
	Url string
}

func (t *Track) GetImageAndUrl(country string) {
	res,err := itunessearch.Search(t.Name + "　" + t.Artist + "　" + t.Album, country, "music")
	if err == nil {
		for _, result := range res {
			t.Image = result.ArtworkUrl60
			t.Url = result.CollectionViewUrl
		}
	}
}

func NewTrack(track itunes.Track) Track {
	return Track{
		Track: track,
	}
}