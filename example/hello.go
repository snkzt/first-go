package main

import (
  "fmt"
  "github.com/snkzt/first-go-unit-test/morestrings"
  "github.com/google/go-cmp/cmp"
)

func main() {
  fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
  fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}

