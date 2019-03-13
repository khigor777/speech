package vendorapi

import (
	"github.com/khigor777/speech/pkg/cache"
)

type cacher interface {
	Set(key string, val []byte)
	Get(key string) ([]byte, error)
}

type api interface {
	TextToVoice(string) ([]byte, error)
}

type YandexApi struct {
	cache cacher
}

func NewYandexApi(cache cacher) *YandexApi {
	return &YandexApi{
		cache: cache,
	}
}

func (y *YandexApi) auth() (string, error) {
	b, err := y.cache.Get("token")
	if err == cache.ErrNotFound {

	}
	return string(b), nil
}

func (y *YandexApi) TextToVoice(text string) ([]byte, error) {
	panic("create")
}
