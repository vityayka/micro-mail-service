package main

import (
	"log"
	"net/http"
)

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestPayload mailMessage
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		log.Printf("parsing request json error: %v", err)
		app.errorJson(w, err)
		return
	}

	message := Message{
		From:    requestPayload.From,
		To:      requestPayload.To,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
	}
	err = app.Mail.SendSMTPMessage(message)
	if err != nil {
		log.Printf("smtp sending fail: %v", err)
		app.errorJson(w, err)
		return
	}

	rsp := jsonResponse{
		Error:   false,
		Message: "mail sent",
	}

	app.writeJSON(w, http.StatusAccepted, rsp)
}
