package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	gomail "gopkg.in/gomail.v2"
)

func sendInvoice(w http.ResponseWriter, r *http.Request) {
	log.Println("Sending Invoice...")

	msg := gomail.NewMessage()
	msg.SetHeader("From", "magnus.andreas.holmen@cegal.com")
	msg.SetHeader("To", "lente147@gmail.com")
	msg.SetHeader("Subject", "Invoice")
	msg.SetBody("text/html", "You owe us money!!!!!!!!!")
	//msg.Attach("/home/User/cat.jpg")
	pass := os.Getenv("pass")

	// n := gomail.NewDialer("themagnus1208@gmail.com", 587, "<paste your gmail account here>", "<paste Google password or app password here>")

	n := gomail.NewDialer("smtp.gmail.com", 587, "themagnus1208@gmail.com", pass)

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	} else {
		log.Println("Nothing failed so it should have been sent...")
	}

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/Invoice", sendInvoice)
	log.Fatal(http.ListenAndServe(":8080", router))
}
