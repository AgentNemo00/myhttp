package http

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/AgentNemo00/myhttp/checksum"
	"io"
	"net/http"
	"testing"
)

type MockClient struct {
	Callback func(request *http.Request) (*http.Response, error)
}

func (m MockClient) Do(request *http.Request) (*http.Response, error) {
	return m.Callback(request)
}

func TestChecksum_Check(t *testing.T) {
	type fields struct {
		Url string
	}
	tests := []struct {
		name    string
		fields  fields
		want    checksum.Checksum
		wantErr bool
		client  Client
	}{
		{
			name:    "success",
			fields:  fields{Url: "www.example.com"},
			want:    checksum.Checksum(fmt.Sprintf("%x", md5.Sum([]byte("EXAMPLEBODY")))),
			wantErr: false,
			client: MockClient{Callback: func(request *http.Request) (*http.Response, error) {
				return &http.Response{Body: io.NopCloser(bytes.NewBufferString("EXAMPLEBODY"))}, nil
			}},
		},
		{
			name:    "error",
			fields:  fields{Url: "!invalid"},
			want:    "",
			wantErr: true,
			client: MockClient{Callback: func(request *http.Request) (*http.Response, error) {
				return &http.Response{}, fmt.Errorf("some error")
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Checksum{
				Url: tt.fields.Url,
			}
			DefaultClient = tt.client
			got, err := c.Check()
			if (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Check() got = %v, want %v", got, tt.want)
			}
		})
	}
}
