package ytrwrap

import (
	"net/http"
	"time"
)

//
// DefaultClientTimeout is waiting for result interval
// used by default in default client
//
const DefaultClientTimeout = time.Second * 50

//
// YandexTranslateAPI is default URL base for requests
//
const YandexTranslateAPI = "https://translate.yandex.net/api/v1.5"

type tr struct {
	key    string
	client *http.Client
	apiURL string
}

//
// NewYandexTranslate creates service client
// w/ default HTTP client
//
func NewYandexTranslate(key string) *tr {
	return NewYandexTranslateWithClient(key, &http.Client{
		Timeout: DefaultClientTimeout,
	})
}

//
// NewYandexTranslateWithClient service client
// w/ associated key and optional HTTP client
// in case of using proxy or different timeouts
//
func NewYandexTranslateWithClient(key string, client *http.Client) *tr {
	p := &tr{
		key:    key,
		client: client,
		apiURL: YandexTranslateAPI,
	}

	if nil == p.client {
		p.client = &http.Client{
			Timeout: DefaultClientTimeout,
		}
	}

	return p
}

//
// NewYandexTranslateWithClientAndURL is a special ctor allows to specify
// custom client and apiURL
//
func NewYandexTranslateWithClientAndURL(key string, client *http.Client, apiURL string) *tr {

	p := NewYandexTranslateWithClient(key, client)
	p.apiURL = apiURL

	return p
}
