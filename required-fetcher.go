package ytrwrap

type fetcherInterface interface {
	Get(url string) ([]byte, int, error)
}
