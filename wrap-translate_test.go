package ytrwrap

import (
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
