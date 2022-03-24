package main

import (
	"fmt"

	"github.com/fnapoles/go-tutorial/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))

	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}