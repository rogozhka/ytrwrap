//
// Example how to handle custom errors
//

package main

import (
	"fmt"
	"log"

	"github.com/rogozhka/ytrwrap"
)

func main() {

	tr := ytrwrap.NewYandexTranslate("<your-api-key")

	src := "the pony eat grass"

	out, err := tr.Translate(src, ytrwrap.FR, nil)

	if err != nil {
		switch err.Code() {
		case ytrwrap.KEY_WRONG:
			log.Println("KEY_WRONG")
		case ytrwrap.KEY_BLOCKED:
			log.Println("KEY_BLOCKED")
		case ytrwrap.LIMIT_DAILY_EXCEEDED:
			log.Println("LIMIT_DAILY_EXCEEDED")
		case ytrwrap.LIMIT_TEXTSIZE_EXCEEDED:
			log.Println("LIMIT_TEXTSIZE_EXCEEDED")
		case ytrwrap.NOT_SUPPORTED_DIRECTION:
			log.Println("NOT_SUPPORTED_DIRECTION")
		default:
			log.Println(err.Message())
		}
	}

	fmt.Printf("%s", out)
	//
	// le poney, manger de l'herbe
	//

}
