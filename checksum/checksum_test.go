package checksum

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestByResponse(t *testing.T) {
	type args struct {
		response *http.Response
	}
	tests := []struct {
		name    string
		args    args
		want    Checksum
		wantErr bool
	}{
		{
			name:    "success",
			args:    args{response: &http.Response{Body: io.NopCloser(bytes.NewBufferString("EXAMPLERESPONSE"))}},
			want:    Checksum(fmt.Sprintf("%x", md5.Sum([]byte("EXAMPLERESPONSE")))),
			wantErr: false,
		},
		{
			name:    "no body",
			args:    args{response: &http.Response{Body: nil}},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ByResponse(tt.args.response)
			if (err != nil) != tt.wantErr {
				t.Errorf("ByResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ByResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
