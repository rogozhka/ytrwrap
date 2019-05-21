package ytrwrap

import (
	"fmt"
	"net/http"
)

type voidResult struct {
	bb   []byte
	code int
	err  error
}

type voidClient struct {
	results map[string]voidResult
	lastURL string
}

func newVoidClient() *voidClient {
	p := &voidClient{}
	return p
}

func (p *voidClient) Get(url string) ([]byte, int, error) {

	p.lastURL = url

	res, is := p.results[url]
	if !is {
		return nil, http.StatusInternalServerError, fmt.Errorf("result does not exist")
	}

	return res.bb, res.code, res.err
}

//
// Set allows to store result for future Get calls with that url
//
func (p *voidClient) Set(url string, bb []byte, code int, err error) {
	if nil == p.results {
		p.results = make(map[string]voidResult)
	}

	p.results[url] = voidResult{
		bb:   bb,
		err:  err,
		code: code,
	}
}

//
// LastURL returns last used url w/ Get() call
//
func (p *voidClient) LastURL() string {
	return p.lastURL
}
