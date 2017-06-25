package track

import (
	"github.com/mattn/itunes-search-api"
)

type apiCache struct {
	image string
	url   string
	key   string
}

func (a apiCache) isKey(key string) bool {
	return a.key == key
}

func (a *apiCache) set(key string, result itunessearch.Result) {
	a.image = result.ArtworkUrl60
	a.url = result.CollectionViewUrl
	a.key = key
}

func (a apiCache) bind(track *Track) {
	track.Image = a.image
	track.Url = a.url
}

func newCache() apiCache {
	return apiCache{}
}
