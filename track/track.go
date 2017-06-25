package track

import (
	"fmt"
	"github.com/mattn/itunes-search-api"
	"github.com/polidog/go-itunes"
)

type Track struct {
	itunes.Track
	Image string
	Url   string
}

var cache = newCache()

func (t *Track) GetImageAndUrl(country string) {
	t.Image = "https://raw.githubusercontent.com/polidog/go-nowplaying/master/gopher.png"
	cacheKey := t.searchWord() + country

	if cache.isKey(cacheKey) {
		cache.bind(t)
		return
	}

	res, err := itunessearch.Search(t.searchWord(), country, "music")

	if err != nil {
		return
	}

	if len(res) == 0 {
		return
	}

	for _, result := range res {
		if result.CollectionName == t.Album {
			cache.set(cacheKey, result)
			cache.bind(t)
			return
		}
	}
	cache.set(cacheKey, res[0])
	cache.bind(t)
}

func (t Track) searchWord() string {
	return fmt.Sprintf("%s %s", t.Artist, t.Album)
}

func NewTrack(track itunes.Track) Track {
	return Track{
		Track: track,
	}
}
