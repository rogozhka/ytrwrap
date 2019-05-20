package ytrwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTr_TranslateRU(t *testing.T) {

	tr := createTestClientFromEnv()
	res, err := tr.Translate("the pony eat grass", RU, nil)
	assert.Nil(t, err, "err")
	assert.Equal(t, "пони едят траву", res, "lc")
}

func TestTr_TranslateFR(t *testing.T) {

	tr := createTestClientFromEnv()
	res, err := tr.Translate("the pony eat grass", FR, nil)
	assert.Nil(t, err, "err")
	assert.Equal(t, "le poney, manger de l'herbe", res, "lc")
}

func TestTr_Translate2(t *testing.T) {

	tr := createTestClientFromEnv()

	trash := "asdsfkjshflkjsadf--"

	res, err := tr.Translate(trash, RU, nil)
	assert.Nil(t, err, "err")
	assert.Equal(t, trash, res, "lc")
}
