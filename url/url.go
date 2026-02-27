package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawURL := "https://example.com/?test=testvalue&test=testvalue2&search=hello"
	parsedURL, _ := url.Parse(rawURL)
	for k, v := range parsedURL.Query() {
		fmt.Println(k, "=", v)
	}
}

/*
Output:
  test = [testvalue testvalue1]
  search = [hello]
*/
