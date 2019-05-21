[![GoDoc](https://godoc.org/github.com/rogozhka/ytrwrap?status.svg)](https://godoc.org/github.com/rogozhka/ytrwrap)
[![Travis](https://travis-ci.org/rogozhka/ytrwrap.svg?branch=master)](https://travis-ci.org/rogozhka/ytrwrap)
[![Go Report Card](https://goreportcard.com/badge/github.com/rogozhka/ytrwrap)](https://goreportcard.com/report/github.com/rogozhka/ytrwrap)
[![codecov](https://codecov.io/gh/rogozhka/ytrwrap/branch/master/graph/badge.svg)](https://codecov.io/gh/rogozhka/ytrwrap)

**ytrwrap** is a wrapper for [Yandex.Translate API](https://tech.yandex.com/translate/). Free API key is required to use machine translation service. Supports more than 90 languages and can translate separate words or complete texts. API consists of 3 simple methods: Languages, Translate, Detect.

## Usage example

```go
package main

import (
	"fmt"

	"github.com/rogozhka/ytrwrap"
)

func main() {

	tr := ytrwrap.NewYandexTranslate("<your-api-key")

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
```



## Contribution

Welcome feedback Issues and PR.



## Licence

Released under the [MIT License](https://github.com/rogozhka/ytrwrap/blob/master/LICENSE).