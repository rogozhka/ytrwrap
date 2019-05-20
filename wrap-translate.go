package ytrwrap

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type TextFormat string

const (
	PlainText = "plain"
	HTML      = "html"
)

func (p *tr) Translate(text string, to, from LC, format TextFormat) (string, *apiError) {

	optFormat := "plain"

	switch format {
	case PlainText:
		fallthrough
	case HTML:
		optFormat = string(format)
	}

	strLang := ""
	if len(from) > 0 {
		strLang = fmt.Sprintf("%s-%s", from, to)
	} else {
		strLang = string(to)
	}

	url := formatURLTranslate(p.apiURL,
		url.Values{
			"key":    {p.key},
			"text":   {text},
			"format": {optFormat},
			"lang":   {strLang},
		})

	type resp struct {
		GenericResponse
		LangDir LD       `json:lang`
		Text    []string `json:"text"`
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

	if len(dataResp.Text) < 1 {
		dataResp.ErrorCode = CANNOT_TRANSLATE
		return "", newError(dataResp.GenericResponse)
	}

	return dataResp.Text[0], nil
}

func formatURLTranslate(api string, values url.Values) string {
	return fmt.Sprintf("%s/tr.json/translate?%s", api, values.Encode())
}
