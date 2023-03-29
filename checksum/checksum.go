package checksum

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
)

var (
	errNoBody = fmt.Errorf("no body")
)

// Checksum definition
type Checksum string

// ByResponse - returns the checksum for the response body
func ByResponse(response *http.Response) (Checksum, error) {
	if response.Body == nil {
		return "", errNoBody
	}
	bodyRaw, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return Checksum(fmt.Sprintf("%x", md5.Sum(bodyRaw))), err
}
