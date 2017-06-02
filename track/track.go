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

	res,err := itunessearch.Search(t.Artist + " " + t.Album, country, "music")
	t.Image = "https://raw.githubusercontent.com/polidog/go-nowplaying/master/gopher.png"
	if err != nil {
		return
	}

	if len(res) == 0 {
		return
	}

	for _, result := range res {
		if result.CollectionName == t.Album {
			t.Image = result.ArtworkUrl60
			t.Url = result.CollectionViewUrl
			return
		}
	}

	t.Image = res[0].ArtworkUrl60
	t.Url = res[0].CollectionViewUrl
}

func NewTrack(track itunes.Track) Track {
	return Track{
		Track: track,
	}
}