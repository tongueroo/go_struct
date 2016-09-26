package main

import (
  "fmt"
)

type Location struct {
    x     *int64
    y     int
    valid bool
}

func main() {
  loc := Location{}
  loc.x = 10
  fmt.Println(loc.x)
}
