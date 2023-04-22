package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"

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
				if !slices.Contains(params, k) {
					params = append(params, k)
					fmt.Println(k)
				}
			}
		}

	}
}
