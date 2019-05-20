package ytrwrap

import (
    "fmt"
    "net/http"
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


func createTestClientFromEnv() *tr {
    return NewYandexTranslateWithClient(getKeyEnv(), &http.Client{
        Timeout: DefaultClientTimeout,
    })
}

