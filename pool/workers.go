package pool

import (
	"github.com/AgentNemo00/myhttp/checksum"
	"github.com/AgentNemo00/myhttp/http"
)

// result - worker result
// name of the worker
// checksum of the body
// error if some error occurred
type result struct {
	Name     string
	Checksum checksum.Checksum
	Error    error
}

// Worker - definition
type Worker func() result

// WorkerByURl - returns a Worker for the given url
func WorkerByURl(url string) Worker {
	return func() result {
		check := http.Checksum{Url: url}
		c, err := check.Check()
		return result{
			Checksum: c,
			Error:    err,
			Name:     url,
		}
	}
}
