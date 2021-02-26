package main

import (
	"net/url"
	"testing"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		url   string
		want  []string
		count int
	}{
		{
			url: "https://www.google.com/path/to/file/a.php",
			want: []string{
				"https://www.google.com/path/to/file",
				"https://www.google.com/path/to",
				"https://www.google.com/path",
				"https://www.google.com/",
			},
			count: 4,
		},
	}
	for _, tt := range tests {
		u, _ := url.Parse(tt.url)
		got := format(u)
		if len(tt.want) != tt.count {
			t.Errorf("Format(%v) = %v (%d), want %v (%d)", tt.url, got, len(got), tt.want, tt.count)
		}
	}
}
