package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	seen := make(map[string]bool)
	for sc.Scan() {
		u, err := url.Parse(sc.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "parse failure: %s\n", err)
			continue
		}

		l := format(u)

		for i := 0; i < len(l); i++ {
			key := l[i]
			// Only output each host + path + params combination once
			if _, exists := seen[key]; exists {
				continue
			}
			seen[key] = true
			fmt.Println(key)
		}
	}
}

func format(u *url.URL) []string {
	formatted := []string{}
	dir, _ := filepath.Split(u.Path)
	pathsFromUrl := strings.Split(dir, "/")
	path := ""
	for t := 0; t < len(pathsFromUrl)-1; t++ {
		if pathsFromUrl[t] != "" {
			path = fmt.Sprintf("%s/%s", path, pathsFromUrl[t])
		}
		u.Path = path
		u.RawQuery = ""
		u.Fragment = ""
		formatted = append(formatted, u.String())
	}

	return formatted
}
