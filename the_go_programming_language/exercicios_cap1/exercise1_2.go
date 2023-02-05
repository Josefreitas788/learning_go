package main

import (
  "fmt"
  "os"
)

func main() {
  for i := 1; i < len(os.Args); i++ {
    fmt.Println("Index: ",i, " Argumento: ", os.Args[i])
  }
}
