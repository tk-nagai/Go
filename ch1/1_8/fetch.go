// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		isSucceeded := fetch(url)
		if !isSucceeded {
			os.Exit(1)
		}
	}
}

func fetch(url string) bool {
	resp, err := http.Get(validateUrl(url))
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return false
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		return false
	}

	return true
}

func validateUrl(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url
	}

	return "http://" + url
}
