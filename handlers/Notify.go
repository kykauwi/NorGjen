package handlers

import (
	"log"
	"net/http"
	"os"

	gomail "gopkg.in/gomail.v2"
)

func SendNotification(w http.ResponseWriter, r *http.Request) {
	log.Println("Sending Invoice...")

	msg := gomail.NewMessage()
	msg.SetHeader("From", "magnus.andreas.holmen@cegal.com")
	msg.SetHeader("To", "lente147@gmail.com")
	msg.SetHeader("Subject", "Invoice")
	msg.SetBody("text/html", "You owe us money!!!!!!!!!")
	msg.Attach("Invoice.pdf")
	pass := os.Getenv("pass")

	n := gomail.NewDialer("smtp.gmail.com", 587, "themagnus1208@gmail.com", pass)

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	} else {
		log.Println("Nothing failed so it should have been sent...")
	}

}
