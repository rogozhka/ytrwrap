package ytrwrap

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTr_Langs(t *testing.T) {

	tr := createRealTestClientFromEnv()

	ll, err := tr.Langs("en")
	assert.Nil(t, err, "err")
	assert.NotNil(t, ll, "res")
	assert.Equal(t, "English", ll["en"], "in map")
}

func TestURLLangs(t *testing.T) {

	dummyKey := "trnsl.there.is.no.key"

	client := newVoidClient()
	tr := NewYandexTranslateWithClient(dummyKey, client)

	uiLC := AF

	_, apierr := tr.Langs(uiLC)
	assert.NotNil(t, apierr, "Langs err")

	baseURL, err := url.Parse(YandexTranslateAPI)
	assert.Nil(t, err, "URL.Parse err")

	theURL, err := url.Parse(client.LastURL())
	assert.Nil(t, err, "URL.Parse err")
	assert.Equal(t, baseURL.Host, theURL.Host, "url")
	assert.Equal(t, baseURL.Path+"/tr.json/getLangs", theURL.Path, "url")

	values := theURL.Query()
	assert.Equal(t, dummyKey, values.Get("key"), "url")
	assert.Equal(t, string(uiLC), values.Get("ui"), "url")
}
