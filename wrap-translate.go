package ytrwrap

import (
	"encoding/json"
	"fmt"
	"net/url"
)

//
// TextFormat - formatting mode
//
type TextFormat string

const (
	//
	// PlainText - default formatting mode
	//
	PlainText TextFormat = "plain"

	//
	// HTML means translation will preserve source tags
	//
	HTML TextFormat = "html"
)

//
// TranslateOpt - optional params for Translate method
//
type TranslateOpt struct {
	//
	// OutputFormat affects result formatting
	//
	OutputFormat TextFormat

	//
	// From if omitted, detection is used
	//
	From LC
}

//
// Translate makes the translation from LC to LC
// w/ output format plainText by default or HTML
// from arg is optional
//
func (p *tr) Translate(text string, to LC, opt *TranslateOpt) (string, *apiError) {

	if nil == opt {
		opt = &TranslateOpt{}
	}

	switch opt.OutputFormat {
	case PlainText:
		break
	case HTML:
		break
	default:
		opt.OutputFormat = PlainText
	}

	strLang := ""
	if len(opt.From) > 0 {
		strLang = fmt.Sprintf("%s-%s", opt.From, to)
	} else {
		strLang = string(to)
	}

	url := formatURLTranslate(p.apiURL,
		url.Values{
			"key":    {p.key},
			"text":   {text},
			"format": {string(opt.OutputFormat)},
			"lang":   {strLang},
		})

	type resp struct {
		GenericResponse
		LangDir LD       `json:"lang"`
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
