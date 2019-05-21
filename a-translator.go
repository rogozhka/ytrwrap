package ytrwrap

//
// YandexTranslateAPI is default URL base for requests
//
const YandexTranslateAPI = "https://translate.yandex.net/api/v1.5"

type tr struct {
	key string

	fetcher fetcherInterface

	apiURL string
}

//
// NewYandexTranslate creates service client
// w/ default HTTP client
//
func NewYandexTranslate(key string) *tr {
	return NewYandexTranslateWithClient(key, nil)
}

//
// NewYandexTranslateWithClient service client
// w/ associated key and optional HTTP client
// in case of using proxy or different timeouts
//
func NewYandexTranslateWithClient(key string, client fetcherInterface) *tr {
	p := &tr{
		key:     key,
		fetcher: client,
		apiURL:  YandexTranslateAPI,
	}

	if nil == p.fetcher {
		p.fetcher = NewFetcher(nil)
	}

	return p
}

//
// NewYandexTranslateWithClientAndURL is a special ctor allows to specify
// custom client and apiURL
//
func NewYandexTranslateWithClientAndURL(key string, fetcher fetcherInterface, apiURL string) *tr {

	p := NewYandexTranslateWithClient(key, fetcher)
	p.apiURL = apiURL

	return p
}
