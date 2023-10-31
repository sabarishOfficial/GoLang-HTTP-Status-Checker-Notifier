package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"strconv"
)

// email funcations
func sendEmailNotification(subject, body string) {
	// email details
	from := ""     //sender mail id
	password := "" //send mail app password
	toList := ""   // to mail
	hostname := ""
	port := "587"
	msg := []byte("To: " + toList + "\r\n" + "Subject: " + subject + "\r\n" + "\r\n" +
		body + "\r\n")

	// login auth
	auth := smtp.PlainAuth("", from, password, hostname)

	// email sending
	err := smtp.SendMail(hostname+":"+port, auth, from, []string{toList}, []byte(msg))

	if err != nil {
		fmt.Printf("Email not sending %d\n", err)
		os.Exit(1)
	}
}

func main() {
	arg := os.Args

	// error handling
	if len(arg) < 2 {
		fmt.Printf("Usage: ./http-get <url>\n")
		os.Exit(1)
	}

	// uri verification
	if _, err := url.ParseRequestURI(arg[1]); err != nil {
		fmt.Printf("Url invaild format %d\n,", err)
		os.Exit(1)
	}

	// variable and argument traversal
	response, err := http.Get(arg[1])

	// error handling
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	// status code verification If the status code is not 200, an email will be sent to the server administrator.
	if response.StatusCode != 200 {
		fmt.Printf("somthing went wrong: %d\n", response.StatusCode)
		status_code := strconv.Itoa(response.StatusCode)
		sendEmailNotification("Server Status Code: "+status_code, "server problem")

	}

	// If it is 200, the server is up and functioning.
	if response.StatusCode == 200 {
		fmt.Printf("Http Status Code: %d\n", response.StatusCode)
	}

}
