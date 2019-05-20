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
