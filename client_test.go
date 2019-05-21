package ytrwrap

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientGet(t *testing.T) {
	ff := NewFetcher(nil)
	_, code, err := ff.Get("http://localhost:1234")
	assert.NotNil(t, err, "err")
	assert.Equal(t, http.StatusInternalServerError, code, "code")
}
