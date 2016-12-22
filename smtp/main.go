package main

import (
	"log"
	"net/smtp"
)

const (
	//	TO = "recipient@example.net"
	TO = "awisu2@gmail.com"
)

func main() {
	sendMail()

}

// gmailで送る
// https://support.google.com/a/answer/2956491

func sendMail() {
	log.Println("========== sendMail ==========")

	// Set up authentication information.
	// PlainAuth(identity, username, password, host string) Auth
	identity := ""
	//	username := "user@example.com"
	username := "awisu2@gmail.com"
	//	password := "password"
	password := "ggyuichi"
	//	host := "mail.example.com"
	host := "smtp.gmail.com"
	auth := smtp.PlainAuth(identity, username, password, host)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{TO}
	msg := []byte("To: " + TO + "\r\n" +
		"Subject: go smtp sample\r\n" +
		"\r\n" +
		"This is the email body.\r\n")

	//	addr string, a Auth, from string, to []string, msg []byte
	//	addr := "mail.example.com:25"
	addr := "smtp.gmail.com:587"
	from := "awisu2@gmail.com"
	err := smtp.SendMail(addr, auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("senMail End")
}
