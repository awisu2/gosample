/**
 * regexp
 */
package main

import (
	"fmt"
	"regexp"
	"strings"
	"strconv"
)

func main() {
	match()
	matchString()
	matchReader()
	regexpTests()
	allMatch()
}

func matchString() {
	fmt.Println("========== matchString ==========")

	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.
	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

	//true
	//true
	//false
	//false
	fmt.Println(validID.MatchString("adam[23]"))
	fmt.Println(validID.MatchString("eve[7]"))
	fmt.Println(validID.MatchString("Job[48]"))
	fmt.Println(validID.MatchString("snakey"))
}

func match() {
	fmt.Println("========== match ==========")

	pattern := `abc`

	// true
	matched, err := regexp.Match(pattern, []byte("123abcdef"))
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(matched)
	}

	// false
	matched, err = regexp.Match(pattern, []byte("ZZZZ"))
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(matched)
	}

	// patternが正規表現に当てはまらない場合エラー
	// Error :  error parsing regexp: missing closing ]: `[123`
	matched, err = regexp.Match(`[123`, []byte("ZZZZ"))
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(matched)
	}
}

func matchReader() {
	fmt.Println("========== matchReader ==========")

	// true
	r := strings.NewReader("123abcdef")
	matched, err := regexp.MatchReader(`def`, r)
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println(matched)
	}
}

func regexpTests() {
	fmt.Println("========== regexpTests ==========")

	s := "this is the sample : Id123ABC id234DEF Id456GHI Id789JKL"

	// Id[0-9]+
	reg := regexp.MustCompile(`Id[0-9]+`)
	fmt.Println(reg.String())

	// [Id123 Id456 Id789]
	hitNum := -1 // all
	finds := reg.FindAllString(s, hitNum)
	fmt.Println(finds)

	// [Id123 Id456]
	hitNum = 2 // 2 stop
	finds = reg.FindAllString(s, hitNum)
	fmt.Println(finds)

	// this is the sample :
	// ABC id234DEF
	// GHI
	// JKL
	finds = reg.Split(s, -1)
	for _, v := range finds {
		fmt.Println(v)
	}

	// [[Id123 123] [Id456 456]]
	reg = regexp.MustCompile(`Id([0-9]+)`)
	findses := reg.FindAllStringSubmatch(s, -1)
	fmt.Println(findses)

	// [[Id123ABC 123 ABC] [Id456GHI 456 GHI] [Id789JKL 789 JKL]]
	reg = regexp.MustCompile(`Id([0-9]+)([A-Z]+)`)
	findses = reg.FindAllStringSubmatch(s, -1)
	fmt.Println(findses)

	// multi string
	// [Id123番]
	reg = regexp.MustCompile(`Id[0-9]+番`)
	finds = reg.FindAllString("Id123番", -1)
	fmt.Println(finds)


	// 全部を同じ長さの0フィルにする
	// a_0123_0123_3456
	repStr := zeroFillInStr("a_123_123_3456", 4)
	fmt.Println(repStr)
	// a_123_123_3456
	repStr = zeroFillInStr("a_123_123_3456", 2)
	fmt.Println(repStr)
}

/**
 * 全部を同じ長さの0フィルにする
 * 一致した文字列を変換する関数はなく
 * 同じ文字が連続した場合などに対処するため直指定での変換を採用
 */
func zeroFillInStr(str string, length int) string {
	// 文字列から数値を抜粋
	reg := regexp.MustCompile(`[0-9]+`)
	finds := reg.FindAllString(str, -1)

	// Replacer用に変換配列を作成
	replaces := []string{}
	format := `%0` + strconv.Itoa(length) + `d`
	for _, find := range finds {
		replaces = append(replaces, find)
		i, _ := strconv.Atoi(find)
		replaces = append(replaces, fmt.Sprintf(format, i))
	}
	r := strings.NewReplacer(replaces...)

	// 変換
	return r.Replace(str)
}

func allMatch() {
	fmt.Println("========== allMatch(comma int value) ==========")

	reg := regexp.MustCompile(`^([0-9]+,)*[0-9]$`)
	// true
	printBool(reg.Match([]byte("1,2,3")))
	// false
	printBool(reg.Match([]byte(",1,2,3")))
	// false
	printBool(reg.Match([]byte("1,2,3,")))
	// true
	printBool(reg.Match([]byte("1")))
	// false
	printBool(reg.Match([]byte("")))
}

func printBool(b bool) {
	if b {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
