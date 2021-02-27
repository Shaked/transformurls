package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		url   string
		want  []string
		count int
	}{
		{
			url: "https://www.google.com/path/to/file/a.php?abc=def&efd=123#bbb=1",
			want: []string{
				"https://www.google.com",
				"https://www.google.com/path",
				"https://www.google.com/path/to",
				"https://www.google.com/path/to/file",
			},
			count: 4,
		},
	}
	for _, tt := range tests {
		u, _ := url.Parse(tt.url)
		got := format(u)
		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("Format(%v) = %v (%d), want %v (%d)", tt.url, got, len(got), tt.want, tt.count)
		}
	}
}
