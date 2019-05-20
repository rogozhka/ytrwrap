package ytrwrap

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

//
// Detect used to detect text language
// hints are useful to suggest supposed LC
//
func (p *tr) Detect(text string, hints []LC) (LC, *apiError) {

	arr := []string{}

	for _, t := range hints {
		str := strings.TrimSpace(string(t))
		if len(str) < 2 {
			continue
		}
		arr = append(arr, str)
	}

	strHints := strings.Join(arr, ",")

	url := formatURLDetect(p.apiURL, url.Values{"key": {p.key}, "hint": {strHints}, "text": {text}})

	type resp struct {
		GenericResponse
		Lang LC `json:"lang"`
	}

	dataResp := resp{}

	bb, code, err := p.getRequest(url)
	if err != nil {
		return "", apiErrorf("GET | %v", err)
	}

	if err := json.Unmarshal(bb, &dataResp); err != nil {
		return "", apiErrorf("Unmarshal | %v", err)
	}

	if code != OK {
		dataResp.ErrorCode = code
		return "", newError(dataResp.GenericResponse)
	}

	return dataResp.Lang, nil
}

func formatURLDetect(api string, values url.Values) string {
	return fmt.Sprintf("%s/tr.json/detect?%s", api, values.Encode())
}
