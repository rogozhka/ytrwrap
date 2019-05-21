package ytrwrap

import (
	"fmt"
	"net/http"
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

func TestTrLangsErrUnmarshal(t *testing.T) {
	dummyKey := "trnsl.there.is.no.key"

	client := newVoidClient()
	tr := NewYandexTranslateWithClient(dummyKey, client)
	_, apierr := tr.Languages(EN)
	assert.NotNil(t, apierr, "detect err")

	client.Set(client.LastURL(), []byte("broken-data"), http.StatusOK, nil)

	_, err := tr.Languages(EN)
	assert.NotNil(t, err, "err")
	assert.Equal(t, WRAPPER_INTERNAL_ERROR, err.ErrorCode, "exp err")
	assert.Equal(t, "500 | Unmarshal | invalid character 'b' looking for beginning of value", err.Error(), "exp err")
}

func TestTrLangsErrGET(t *testing.T) {
	dummyKey := "trnsl.there.is.no.key"

	client := newVoidClient()
	tr := NewYandexTranslateWithClient(dummyKey, client)
	_, apierr := tr.Languages(EN)
	assert.NotNil(t, apierr, "detect err")

	msgOffline := "net offline"
	client.Set(client.LastURL(), []byte(""), http.StatusNotFound, fmt.Errorf(msgOffline))

	_, err := tr.Languages(EN)
	assert.NotNil(t, err, "err")
	assert.Equal(t, WRAPPER_INTERNAL_ERROR, err.ErrorCode, "exp err")
	assert.Equal(t, "500 | GET | "+msgOffline, err.Error(), "exp err")
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
