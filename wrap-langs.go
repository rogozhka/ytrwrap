package ytrwrap

import (
	"encoding/json"
	"fmt"
	"net/url"
)

//
// Languages enumerates all the supported LC
//

func (p *tr) Languages(uiLang LC) (map[LC]string, *apiError) {
	return p.Langs(uiLang)
}

//
// Langs is shortcut to Languages used to list all the supported LC
//
func (p *tr) Langs(uiLang LC) (map[LC]string, *apiError) {

	url := formatURLGetLangs(p.apiURL, url.Values{"key": {p.key}, "ui": {string(uiLang)}})

	type resp struct {
		GenericResponse
		Langs map[LC]string `json: langs`
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
