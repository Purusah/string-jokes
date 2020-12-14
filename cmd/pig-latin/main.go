package main

import (
	"flag"
	"fmt"
	latin "github.com/purusah/string-jokes/pkg/pig-latin"
	"strings"
)

func main() {
	flag.Parse()
	words := flag.Args()
	var result []string
	for _, w := range words {
		result = append(result, latin.Encode(w))
	}
	fmt.Println(strings.Join(result, " "))
}
