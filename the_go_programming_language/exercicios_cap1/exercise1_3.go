package main

import (
  "fmt"
  "time"
  "os"
  "strings"
)

func main() {
  s, sep := "", ""
  start := time.Now()
  for i := 1; i<len(os.Args); i++ {
    s += sep + os.Args[i]
    sep = " "
  }
  fmt.Println(s)
  stop := time.Now()
  fmt.Println("Tempo do echo1: " ,stop.Sub(start))

  start = time.Now()
  s, sep = "", ""
  for _, arg := range os.Args[1:] {
    s += sep + arg
    sep = " "
  }
  fmt.Println(s)
  stop = time.Now()
  fmt.Println("Tempo do echo2: " ,stop.Sub(start))

  start = time.Now()
  fmt.Println(strings.Join(os.Args[1:], " "))
  stop = time.Now()
  fmt.Println("Tempo do echo3: " ,stop.Sub(start))

}
