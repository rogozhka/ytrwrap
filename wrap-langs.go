package ytrwrap

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Langs map[LC]string

func (p *tr) Langs(uiLang string) (Langs, *apiError) {

	url := formatURLGetLangs(p.apiURL, url.Values{"key": {p.key}, "ui": {uiLang}})

	type resp struct {
		GenericResponse
		Langs Langs `json: langs`
	}

	dataResp := resp{
		Langs: make(map[LC]string),
	}

	bb, code, err := p.getRequest(url)
	if err != nil {
		return nil, apiErrorf("GET | %v", err)
	}

	if err := json.Unmarshal(bb, &dataResp); err != nil {
		return nil, apiErrorf("Unmarshal | %v", err)
	}

	if code != OK {
		dataResp.ErrorCode = code
		return nil, newError(dataResp.GenericResponse)
	}

	return dataResp.Langs, nil
}

func formatURLGetLangs(api string, values url.Values) string {
	return fmt.Sprintf("%s/tr.json/getLangs?%s", api, values.Encode())
}
