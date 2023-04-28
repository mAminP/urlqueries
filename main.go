package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	var urls []string
	var params []string

	sa := bufio.NewScanner(os.Stdin)
	for sa.Scan() {
		urls = append(urls, sa.Text())
	}

	if sa.Err() != nil {
		// Handle error.
	}
	for _, val := range urls {
		u, err := url.Parse(val)
		if err != nil {
			continue
		}
		if u.RawQuery != "" {
			m, _ := url.ParseQuery(u.RawQuery)
			for k := range m {
				key := strings.TrimSpace(k)
				if !strings.Contains(key, ".") && !strings.Contains(key, ">") && !strings.Contains(key, "<") && !strings.Contains(key, ",") && !strings.Contains(key, "/") && !strings.Contains(key, "\\") && !strings.Contains(key, "\"") && !strings.Contains(key, "'") {
					if !slices.Contains(params, key) {
						params = append(params, key)
						fmt.Println(k)
					}
				}
			}
		}

	}
}
