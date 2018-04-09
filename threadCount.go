package main

import (
  "fmt"
  "runtime"
)

func main() {
  runtime.GOMAXPROCS(100)
  fmt.Printf("Thread: %v\n",  runtime.GOMAXPROCS(-1))
}
