package ytrwrap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeWrong(t *testing.T) {

	tr := NewYandexTranslate("wrong-key")

	_, err := tr.Langs("ru")
	assert.NotNil(t, err, "err")
	assert.Equal(t, KEY_WRONG, err.ErrorCode, fmt.Sprintf("%v", err))
}

func TestNewYandexTranslateWithClientAndURL(t *testing.T) {

	dummyKey := "trnsl.there.is.no.key"
	dummyAPI := "https://there-is-no-such-site.com/v1/"

	tr := NewYandexTranslateWithClientAndURL(dummyKey, nil, dummyAPI)
	assert.NotNil(t, tr.fetcher, "fetcher")
	assert.Equal(t, dummyAPI, tr.apiURL, "api")
	assert.Equal(t, dummyKey, tr.key, "key")
}
