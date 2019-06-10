//
// Example how to use github.com/rogozhka/ytrwrap
// with http proxy
//

package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/rogozhka/ytrwrap"
)

func main() {

	url, err := url.Parse("your-proxy-url")
	if err != nil {
		panic(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(url),
		},
		Timeout: time.Minute * 3,
	}

	tr := ytrwrap.NewYandexTranslateWithClient(
		"<your-api-key",
		ytrwrap.NewFetcher(client),
	)

	src := "the pony eat grass"

	out, err := tr.Translate(src, ytrwrap.FR, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", out)
	//
	// le poney, manger de l'herbe
	//

}
