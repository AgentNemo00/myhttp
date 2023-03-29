package configuration

import (
	"testing"
)

func TestConfiguration_Validate(t *testing.T) {
	type fields struct {
		Parallel int
		Urls     []string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Parallel: 3,
				Urls:     []string{"google.de", "http://google.de", "http://www.google.de", "www.google.de"},
			},
			wantErr: false,
		},
		{
			name: "parallel zero",
			fields: fields{
				Parallel: 0,
				Urls:     []string{"google.de"},
			},
			wantErr: true,
		},
		{
			name: "parallel negative",
			fields: fields{
				Parallel: -2,
				Urls:     []string{"google.de"},
			},
			wantErr: true,
		},
		{
			name: "no urls",
			fields: fields{
				Parallel: 3,
				Urls:     []string{},
			},
			wantErr: true,
		},
		{
			name: "invalid url",
			fields: fields{
				Parallel: 3,
				Urls:     []string{"!csd:invalid"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Configuration{
				Parallel: tt.fields.Parallel,
				Urls:     tt.fields.Urls,
			}
			if err := c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
