package ytrwrap

import (
	"fmt"
	"io/ioutil"
)

//
// getRequest sends HTTP GET
// and reads result
//
func (p *tr) getRequest(url string) ([]byte, APICode, error) {

	var res []byte

	resp, err := p.client.Get(url)
	if resp != nil {
		defer resp.Body.Close()
		resp.Close = true
	}
	if err != nil {
		return res, WRAPPER_INTERNAL_ERROR, err
	}

	res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, WRAPPER_INTERNAL_ERROR, fmt.Errorf("cannot ReadAll | %v", err)
	}

	return res, APICode(resp.StatusCode), nil
}
