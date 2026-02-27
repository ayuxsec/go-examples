package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, err := url.Parse("https://user:pass@example.com:8080/path?q=go+lang#section")
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)   // https
	fmt.Println(u.Host)     // example.com:8080
	fmt.Println(u.Path)     // /path
	fmt.Println(u.RawQuery) // q=go+lang
	fmt.Println(u.Fragment) // section
}
