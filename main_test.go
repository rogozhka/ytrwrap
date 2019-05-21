package ytrwrap

import (
	"fmt"
	"os"
)

func getKeyEnv() string {
	return requireEnv("YTR_KEY")
}

func requireEnv(name string) string {
	v := os.Getenv(name)
	if len(v) < 1 {
		panic(fmt.Sprintf("env not set | %s", name))
	}

	return v
}

func createRealTestClientFromEnv() *tr {
	return NewYandexTranslate(getKeyEnv())
}
