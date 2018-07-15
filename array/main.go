/**
 * array の動作サンプル
 */
package main

import (
  "log"
)

type Hash struct {
  id int
  name string
}

func main () {
  // append
  arrAppend()

  // hash log
  hashlog(Hash{})
  hashlog(Hash{name: "myname"})
  hashlog(Hash{id: 99, name: "next name"})
}

/**
 * append
 * @return {[type]} [description]
 */
func arrAppend () {
  arr := []string{"a", "b", "c"}

  // append
  arr2 := append(arr, "d", "e")

  // append by arr
  arr3 := append(arr, []string{"f", "g"}...)

  log.Println(arr, arr2, arr3)
}

func hashlog (h Hash) {
  log.Println(h)
}

