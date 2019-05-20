package ytrwrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTr_DetectRU(t *testing.T) {

	tr := createTestClientFromEnv()
	lc, err := tr.Detect("мама мыла раму", nil)
	assert.Nil(t, err, "err")
	assert.Equal(t, RU, lc, "lc")
}

func TestTr_DetectEN(t *testing.T) {

	tr := createTestClientFromEnv()
	lc, err := tr.Detect("the pony is eating grass", nil)
	assert.Nil(t, err, "err")
	assert.Equal(t, EN, lc, "lc")
}
