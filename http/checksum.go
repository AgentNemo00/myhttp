package http

import (
	"github.com/AgentNemo00/myhttp/checksum"
	"net/http"
)

type Checksum struct {
	Url string
}

// Check - calls the Url and returns the checksum
func (c Checksum) Check() (checksum.Checksum, error) {
	req, err := http.NewRequest(http.MethodGet, c.Url, nil)
	if err != nil {
		return "", err
	}
	resp, err := DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	return checksum.ByResponse(resp)
}
