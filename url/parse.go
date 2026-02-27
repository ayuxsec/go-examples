package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawURL := "https://user:pass@example.com:8080/path?q=go+lang#section"

	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme:", u.Scheme)        // Scheme: https
	fmt.Println("User:", u.User)            // User: user:pass
	fmt.Println("Username:", u.User.Username()) // Username: user
	password, _ := u.User.Password()
	fmt.Println("Password:", password)      // Password: pass
	fmt.Println("Host:", u.Host)            // Host: example.com:8080
	fmt.Println("Hostname:", u.Hostname())  // Hostname: example.com
	fmt.Println("Port:", u.Port())          // Port: 8080
	fmt.Println("Path:", u.Path)            // Path: /path
	fmt.Println("RawQuery:", u.RawQuery)    // RawQuery: q=go+lang
	fmt.Println("Query param q:", u.Query().Get("q")) // Query param q: go lang
	fmt.Println("Fragment:", u.Fragment)    // Fragment: section
}
