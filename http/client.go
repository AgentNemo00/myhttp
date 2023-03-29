package http

import "net/http"

var (
	DefaultClient Client = &http.Client{}
)

type Client interface {
	Do(request *http.Request) (*http.Response, error)
}
