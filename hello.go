package main

import (
	"fmt"

	"abhinayak.com/user/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
