package ytrwrap

import (
	"fmt"
	"net/http"
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

func TestTr_DetectErrUnmarshal(t *testing.T) {
	dummyKey := "trnsl.there.is.no.key"
	text2detect := "no-text"

	client := newVoidClient()
	tr := NewYandexTranslateWithClient(dummyKey, client)
	_, apierr := tr.Detect(text2detect, nil)
	assert.NotNil(t, apierr, "detect err")

	client.Set(client.LastURL(), []byte("broken-data"), http.StatusOK, nil)

	_, err := tr.Detect(text2detect, nil)
	assert.NotNil(t, err, "err")
	assert.Equal(t, WRAPPER_INTERNAL_ERROR, err.ErrorCode, "exp err")
	assert.Equal(t, "500 | Unmarshal | invalid character 'b' looking for beginning of value", err.Error(), "exp err")
}

func TestTr_DetectErrGET(t *testing.T) {
	dummyKey := "trnsl.there.is.no.key"
	text2detect := "no-text"

	client := newVoidClient()
	tr := NewYandexTranslateWithClient(dummyKey, client)
	_, apierr := tr.Detect(text2detect, nil)
	assert.NotNil(t, apierr, "detect err")

	client.Set(client.LastURL(), []byte(""), http.StatusNotFound, fmt.Errorf("offline"))

	_, err := tr.Detect(text2detect, nil)
	assert.NotNil(t, err, "err")
	assert.Equal(t, WRAPPER_INTERNAL_ERROR, err.ErrorCode, "exp err")
	assert.Equal(t, "500 | GET | offline", err.Error(), "exp err")
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

func TestURLDetectHints(t *testing.T) {

	dummyKey := "trnsl.there.is.no.key"
	text2detect := "no-text"

	client := newVoidClient()
	tr := NewYandexTranslateWithClient(dummyKey, client)

	_, apierr := tr.Detect(text2detect, []LC{EN, DE})
	assert.NotNil(t, apierr, "detect err")

	baseURL, err := url.Parse(YandexTranslateAPI)
	assert.Nil(t, err, "URL.Parse err")

	theURL, err := url.Parse(client.LastURL())
	assert.Nil(t, err, "URL.Parse err")
	assert.Equal(t, baseURL.Host, theURL.Host, "url")
	assert.Equal(t, baseURL.Path+"/tr.json/detect", theURL.Path, "url")

	values := theURL.Query()

	assert.Equal(t, dummyKey, values.Get("key"), "url")
	assert.Equal(t, string(EN)+","+string(DE), values.Get("hint"), "url")
	assert.Equal(t, text2detect, values.Get("text"), "url")
}

func TestURLDetectHintsRare(t *testing.T) {

	dummyKey := "trnsl.there.is.no.key"
	text2detect := "no-text"

	client := newVoidClient()
	tr := NewYandexTranslateWithClient(dummyKey, client)

	_, apierr := tr.Detect(text2detect, []LC{EN, "", "", FR})
	assert.NotNil(t, apierr, "detect err")

	baseURL, err := url.Parse(YandexTranslateAPI)
	assert.Nil(t, err, "URL.Parse err")

	theURL, err := url.Parse(client.LastURL())
	assert.Nil(t, err, "URL.Parse err")
	assert.Equal(t, baseURL.Host, theURL.Host, "url")
	assert.Equal(t, baseURL.Path+"/tr.json/detect", theURL.Path, "url")

	values := theURL.Query()

	assert.Equal(t, dummyKey, values.Get("key"), "url")
	assert.Equal(t, string(EN)+","+string(FR), values.Get("hint"), "url")
	assert.Equal(t, text2detect, values.Get("text"), "url")
}

func TestTrDetectErrCodeNotOK(t *testing.T) {
	dummyKey := "trnsl.there.is.no.key"
	text2detect := "no-text"

	client := newVoidClient()
	tr := NewYandexTranslateWithClient(dummyKey, client)

	_, apierr := tr.Detect(text2detect, []LC{EN})
	assert.NotNil(t, apierr, "detect err")

	client.Set(client.LastURL(), []byte(`{
  "code": 200,
  "lang": "en"
}
`), http.StatusInternalServerError, nil)

	_, err := tr.Detect(text2detect, []LC{EN})
	assert.NotNil(t, err, "err")
	assert.Equal(t, WRAPPER_INTERNAL_ERROR, err.ErrorCode, "exp err")
	assert.Equal(t, "500 | ", err.Error(), "exp err")
}
