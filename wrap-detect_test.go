package ytrwrap

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTr_DetectRU(t *testing.T) {

	tr := createRealTestClientFromEnv()
	lc, err := tr.Detect("мама мыла раму", nil)
	assert.Nil(t, err, "err")
	assert.Equal(t, RU, lc, "lc")
}

func TestTr_DetectEN(t *testing.T) {

	tr := createRealTestClientFromEnv()
	lc, err := tr.Detect("the pony is eating grass", nil)
	assert.Nil(t, err, "err")
	assert.Equal(t, EN, lc, "lc")
}

func TestURLDetect(t *testing.T) {

	dummyKey := "trnsl.there.is.no.key"
	text2detect := "no-text"

	client := newVoidClient()
	tr := NewYandexTranslateWithClient(dummyKey, client)

	_, apierr := tr.Detect(text2detect, nil)
	assert.NotNil(t, apierr, "detect err")

	baseURL, err := url.Parse(YandexTranslateAPI)
	assert.Nil(t, err, "URL.Parse err")

	theURL, err := url.Parse(client.LastURL())
	assert.Nil(t, err, "URL.Parse err")
	assert.Equal(t, baseURL.Host, theURL.Host, "url")
	assert.Equal(t, baseURL.Path+"/tr.json/detect", theURL.Path, "url")

	values := theURL.Query()

	assert.Equal(t, dummyKey, values.Get("key"), "url")
	assert.Equal(t, "", values.Get("hint"), "url")
	assert.Equal(t, text2detect, values.Get("text"), "url")
}
