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
