package splitter

import (
	"reflect"
	"testing"
)

func TestRequestSplitter(t *testing.T) {
	tests := []struct {
		name string
		file string
		want []string
	}{
		{
			name: "delimiter with description",
			file: `### Get Weather
GET /weather
Host: example.com`,
			want: []string{
				"GET /weather\nHost: example.com",
			},
		},
		{
			name: "delimiter alone",
			file: `###
GET /weather
Host: example.com`,
			want: []string{
				"GET /weather\nHost: example.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RequestSplitter([]byte(tt.file))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequestSplitter() = %q, want %q", got, tt.want)
			}
		})
	}
}
