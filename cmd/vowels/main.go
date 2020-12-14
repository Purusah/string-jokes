package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/purusah/string-jokes/pkg/vowels"
)

func main() {
	isReversed := flag.Bool("reverse", false, "set flag to decode string")
	flag.Parse()
	words := flag.Args()
	action := vowels.Encode
	if *isReversed {
		action = vowels.Decode
	}
	fmt.Println(action(strings.Join(words, " ")))
}
