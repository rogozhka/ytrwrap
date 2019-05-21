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

	if res, is := p.results[url]; !is {
		return nil, http.StatusInternalServerError, fmt.Errorf("result does not exist")
	} else {
		return res.bb, res.code, res.err
	}
}

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

func (p *voidClient) LastURL() string {
	return p.lastURL
}
