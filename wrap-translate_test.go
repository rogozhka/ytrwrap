package ytrwrap

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTr_TranslateRU(t *testing.T) {

	tr := createRealTestClientFromEnv()
	res, err := tr.Translate("the pony eat grass", RU, nil)
	assert.Nil(t, err, "err")
	assert.Equal(t, "пони едят траву", res, "lc")
}

func TestTr_TranslateFR(t *testing.T) {

	tr := createRealTestClientFromEnv()
	res, err := tr.Translate("the pony eat grass", FR, nil)
	assert.Nil(t, err, "err")
	assert.Equal(t, "le poney, manger de l'herbe", res, "lc")
}

func TestTr_Translate2(t *testing.T) {

	tr := createRealTestClientFromEnv()

	trash := "asdsfkjshflkjsadf--"

	res, err := tr.Translate(trash, RU, nil)
	assert.Nil(t, err, "err")
	assert.Equal(t, trash, res, "lc")
}

func TestURLTranslateDefault(t *testing.T) {

	dummyKey := "trnsl.there.is.no.key"
	txt := "no-text-to translate"

	client := newVoidClient()
	tr := NewYandexTranslateWithClient(dummyKey, client)

	toLC := EN

	_, apierr := tr.Translate(txt, toLC, nil)
	assert.NotNil(t, apierr, "detect err")

	baseURL, err := url.Parse(YandexTranslateAPI)
	assert.Nil(t, err, "URL.Parse err")

	theURL, err := url.Parse(client.LastURL())
	assert.Nil(t, err, "URL.Parse err")
	assert.Equal(t, baseURL.Host, theURL.Host, "url")
	assert.Equal(t, baseURL.Path+"/tr.json/translate", theURL.Path, "url")

	values := theURL.Query()

	assert.Equal(t, dummyKey, values.Get("key"), "url")
	assert.Equal(t, txt, values.Get("text"), "url")
	assert.Equal(t, string(toLC), values.Get("lang"), "url")
	assert.Equal(t, "plain", values.Get("format"), "url")
}

func TestURLTranslateHTML(t *testing.T) {

	dummyKey := "trnsl.there.is.no.key"
	txt := "no-text-to translate"

	client := newVoidClient()
	tr := NewYandexTranslateWithClient(dummyKey, client)

	toLC := EN

	_, apierr := tr.Translate(txt, toLC, &TranslateOpt{
		OutputFormat: HTML,
	})
	assert.NotNil(t, apierr, "detect err")

	baseURL, err := url.Parse(YandexTranslateAPI)
	assert.Nil(t, err, "URL.Parse err")

	theURL, err := url.Parse(client.LastURL())
	assert.Nil(t, err, "URL.Parse err")
	assert.Equal(t, baseURL.Host, theURL.Host, "url")
	assert.Equal(t, baseURL.Path+"/tr.json/translate", theURL.Path, "url")

	values := theURL.Query()

	assert.Equal(t, dummyKey, values.Get("key"), "url")
	assert.Equal(t, txt, values.Get("text"), "url")
	assert.Equal(t, string(toLC), values.Get("lang"), "url")
	assert.Equal(t, "html", values.Get("format"), "url")
}

func TestURLTranslatePair(t *testing.T) {

	dummyKey := "trnsl.there.is.no.key"
	txt := "no-text-to translate"

	client := newVoidClient()
	tr := NewYandexTranslateWithClient(dummyKey, client)

	toLC := EN
	fromLC := ZH

	_, apierr := tr.Translate(txt, toLC, &TranslateOpt{
		OutputFormat: PlainText,
		From:         fromLC,
	})
	assert.NotNil(t, apierr, "detect err")

	baseURL, err := url.Parse(YandexTranslateAPI)
	assert.Nil(t, err, "URL.Parse err")

	theURL, err := url.Parse(client.LastURL())
	assert.Nil(t, err, "URL.Parse err")
	assert.Equal(t, baseURL.Host, theURL.Host, "url")
	assert.Equal(t, baseURL.Path+"/tr.json/translate", theURL.Path, "url")

	values := theURL.Query()

	assert.Equal(t, dummyKey, values.Get("key"), "url")
	assert.Equal(t, txt, values.Get("text"), "url")
	assert.Equal(t, string(fromLC)+"-"+string(toLC), values.Get("lang"), "url")
	assert.Equal(t, "plain", values.Get("format"), "url")
}

func TestTrTranslateErrUnmarshal(t *testing.T) {
	dummyKey := "trnsl.there.is.no.key"
	text := "no-text"
	lcTO := EN

	client := newVoidClient()
	tr := NewYandexTranslateWithClient(dummyKey, client)
	_, apierr := tr.Translate(text, lcTO, nil)
	assert.NotNil(t, apierr, "detect err")

	client.Set(client.LastURL(), []byte("broken-data"), http.StatusOK, nil)

	_, err := tr.Translate(text, lcTO, nil)
	assert.NotNil(t, err, "err")
	assert.Equal(t, WRAPPER_INTERNAL_ERROR, err.ErrorCode, "exp err")
	assert.Equal(t, "500 | Unmarshal | invalid character 'b' looking for beginning of value", err.Error(), "exp err")
}

func TestTrTranslateErrGET(t *testing.T) {
	dummyKey := "trnsl.there.is.no.key"
	text := "no-text"

	lcTO := EN

	client := newVoidClient()
	tr := NewYandexTranslateWithClient(dummyKey, client)
	_, apierr := tr.Translate(text, lcTO, nil)
	assert.NotNil(t, apierr, "detect err")

	client.Set(client.LastURL(), []byte(""), http.StatusNotFound, fmt.Errorf("offline"))

	_, err := tr.Translate(text, lcTO, nil)
	assert.NotNil(t, err, "err")
	assert.Equal(t, WRAPPER_INTERNAL_ERROR, err.ErrorCode, "exp err")
	assert.Equal(t, "500 | GET | offline", err.Error(), "exp err")
}

func TestTrTranslateErrZeroResult(t *testing.T) {
	dummyKey := "trnsl.there.is.no.key"
	text := "no-text"

	lcTO := EN

	client := newVoidClient()
	tr := NewYandexTranslateWithClient(dummyKey, client)
	_, apierr := tr.Translate(text, lcTO, nil)
	assert.NotNil(t, apierr, "detect err")

	client.Set(client.LastURL(), []byte(`{
  "code": 200,
  "lang": "ru-en",
  "text": [
    ""
  ]
}
`), http.StatusOK, nil)

	_, err := tr.Translate(text, lcTO, nil)
	assert.NotNil(t, err, "err")
	assert.Equal(t, CANNOT_TRANSLATE, err.ErrorCode, "exp err")
	assert.Equal(t, "422 | Empty result", err.Error(), "exp err")
}

func TestTrTranslateErrCodeNotOK(t *testing.T) {
	dummyKey := "trnsl.there.is.no.key"
	text := "no-text"

	lcTO := EN

	client := newVoidClient()
	tr := NewYandexTranslateWithClient(dummyKey, client)
	_, apierr := tr.Translate(text, lcTO, nil)
	assert.NotNil(t, apierr, "detect err")

	client.Set(client.LastURL(), []byte(`{
  "code": 200,
  "lang": "ru-en",
  "text": [
    ""
  ]
}
`), http.StatusInternalServerError, nil)

	_, err := tr.Translate(text, lcTO, nil)
	assert.NotNil(t, err, "err")
	assert.Equal(t, WRAPPER_INTERNAL_ERROR, err.ErrorCode, "exp err")
	assert.Equal(t, "500 | ", err.Error(), "exp err")
}
