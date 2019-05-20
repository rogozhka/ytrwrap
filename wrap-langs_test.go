package ytrwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTr_Langs(t *testing.T) {

	tr := createTestClientFromEnv()

	ll, err := tr.Langs("en")
	assert.Nil(t, err, "err")
	assert.NotNil(t, ll, "res")
	assert.Equal(t, "English", ll["en"], "in map")
}
