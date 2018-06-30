package main

import (
  "flag"

  "os"
  "fmt"
  "bufio"
)

func main () {
  // get arg
  flag.Parse()
  fmt.Println(flag.Args()[0])

  // wait
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
      fmt.Println(scanner.Text())
  }
  if scanner.Err() != nil {
      fmt.Println(scanner.Text())
  }
}