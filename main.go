package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
)

type EmailRequest struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func sendEmail(subject, body string) error {
	from := os.Getenv("GMAIL_USER")
	password := os.Getenv("GMAIL_PASSWORD")
	to := "romain.eyquem@esme.fr" // Remplacez par le destinataire souhaité

	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, password, "smtp.gmail.com"),
		from, []string{to}, msg)
	return err
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Ajout des en-têtes CORS
	w.Header().Set("Access-Control-Allow-Origin", "*") // Autoriser toutes les origines (à adapter en production)
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		// Répondre aux requêtes préflight CORS
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var emailReq EmailRequest
	err := json.NewDecoder(r.Body).Decode(&emailReq)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := sendEmail(emailReq.Subject, emailReq.Body); err != nil {
		http.Error(w, "Failed to send email: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Email sent successfully!")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
