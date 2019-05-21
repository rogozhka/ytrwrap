package ytrwrap

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//
// DefaultClientTimeout is waiting for result interval
// used by default in default client
//
const DefaultClientTimeout = time.Second * 50

type fetcher struct {
	client *http.Client
}

//
// NewFetcher creates required fetcherInterface impl
// w/ optional http.Client (leave nil for defaults)
//
func NewFetcher(client *http.Client) *fetcher {
	p := &fetcher{
		client: client,
	}
	if p.client == nil {
		p.client = &http.Client{
			Timeout: DefaultClientTimeout,
		}
	}

	return p
}

//
// Get makes HTTP GET request
//
func (p *fetcher) Get(url string) ([]byte, int, error) {

	var res []byte

	resp, err := p.client.Get(url)
	if resp != nil {
		defer resp.Body.Close()
		resp.Close = true
	}
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, http.StatusInternalServerError, fmt.Errorf("ReadAll | %v", err)
	}

	return res, resp.StatusCode, nil
}
