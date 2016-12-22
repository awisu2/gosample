package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"
	"strings"
)

func main() {

	parseAddressList()
	readMessage()
}

//アドレスを解析
// ========== parseAddressList ==========
// Alice alice@example.com
// Bob bob@example.com
// Eve eve@example.com
func parseAddressList() {
	fmt.Println("========== parseAddressList ==========")
	const list = "Alice <alice@example.com>, Bob <bob@example.com>, Eve <eve@example.com>"
	emails, err := mail.ParseAddressList(list)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range emails {
		fmt.Println(v.Name, v.Address)
	}
}

//メールを読み込んで解析
// ========== readMessage ==========
// Date: Mon, 23 Jun 2015 11:40:36 -0400
// From: Gopher <from@example.com>
// To: Another Gopher <to@example.com>
// Subject: Gophers at Gophercon
// Message body
func readMessage() {
	fmt.Println("========== readMessage ==========")
	msg := `Date: Mon, 23 Jun 2015 11:40:36 -0400
From: Gopher <from@example.com>
To: Another Gopher <to@example.com>
Subject: Gophers at Gophercon

Message body
`

	r := strings.NewReader(msg)
	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatal(err)
	}

	header := m.Header
	fmt.Println("Date:", header.Get("Date"))
	fmt.Println("From:", header.Get("From"))
	fmt.Println("To:", header.Get("To"))
	fmt.Println("Subject:", header.Get("Subject"))

	body, err := ioutil.ReadAll(m.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
