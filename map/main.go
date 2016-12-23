package main

import (
	"fmt"
	"strconv"
	"time"
)

var maps = map[string]int{}

type val struct {
	key   string
	value int
}

var arrays = []val{}

const NUM = 1000000

func main() {
	loadAverage()
}

//mapとスライスでの負荷調査
//値登録はarrayのほうが早いがほぼ誤差範囲、取得時はmapのほうが高速
// set map time :    571288700
// set array time :  305243900
// map time :          0 , value: 1000000
// array time :  4902200 , value: 1000000
func loadAverage() {
	key := "key" + strconv.Itoa(NUM)

	before := now()
	for i := 1; i <= NUM; i++ {
		maps["key"+strconv.Itoa(i)] = i
	}
	after := now()
	fmt.Println("set map time : ", after-before)

	before = now()
	for i := 1; i <= NUM; i++ {
		arrays = append(arrays, val{"key" + strconv.Itoa(i), i})
	}
	after = now()
	fmt.Println("set array time : ", after-before)

	before = now()
	_v, ok := maps[key]
	after = now()
	if !ok {
		fmt.Println("no data")
	}
	fmt.Println("map time : ", after-before, ", value:", _v)

	ok = false
	_v = -1
	before = now()
	for _, v := range arrays {
		if v.key == key {
			ok = true
			_v = v.value
			break
		}
	}
	after = now()
	if !ok {
		fmt.Println("no data")
	}
	fmt.Println("array time : ", after-before, ", value:", _v)
}

func now() int64 {
	return time.Now().UnixNano()
}
