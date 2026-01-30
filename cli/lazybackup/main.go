package main

import (
	"log"
)

func main() {
	if err := GetRootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}
