package pool

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/AgentNemo00/myhttp/checksum"
	http2 "github.com/AgentNemo00/myhttp/http"
	"io"
	"net/http"
	"reflect"
	"testing"
)

type MockClient struct {
	Callback func(request *http.Request) (*http.Response, error)
}

func (m MockClient) Do(request *http.Request) (*http.Response, error) {
	return m.Callback(request)
}

func TestWorkerByURl(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name   string
		args   args
		want   result
		client http2.Client
	}{
		{
			name: "worker",
			args: args{url: "http://example.com"},
			want: result{
				Name:     "http://example.com",
				Checksum: checksum.Checksum(fmt.Sprintf("%x", md5.Sum([]byte("EXAMPLEBODY")))),
				Error:    nil,
			},
			client: MockClient{Callback: func(request *http.Request) (*http.Response, error) {
				return &http.Response{Body: io.NopCloser(bytes.NewBufferString("EXAMPLEBODY"))}, nil
			}},
		},
		{
			name: "error",
			args: args{url: "http://example.com"},
			want: result{
				Name:     "http://example.com",
				Checksum: "",
				Error:    fmt.Errorf("some error"),
			},
			client: MockClient{Callback: func(request *http.Request) (*http.Response, error) {
				return &http.Response{}, fmt.Errorf("some error")
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			http2.DefaultClient = tt.client
			if got := WorkerByURl(tt.args.url)(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WorkerByURl() = %v, want %v", got, tt.want)
			}

		})
	}
}
