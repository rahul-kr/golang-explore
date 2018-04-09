package main

import (
  "fmt"
  "runtime"
)

func main() {
  fmt.Printf("Thread: %v\n",  runtime.GOMAXPROCS(-1))
}
