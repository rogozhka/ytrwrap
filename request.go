package ytrwrap

//
// getRequest sends HTTP GET
// and reads result
//
func (p *tr) getRequest(url string) ([]byte, APICode, error) {

	bb, code, err := p.fetcher.Get(url)

	return bb, APICode(code), err
}
