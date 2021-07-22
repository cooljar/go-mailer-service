package domain

import "encoding/json"

type Mail struct {
	SenderName string `json:"sender_name"`
	To string `json:"to"`
	Subject string `json:"subject"`
	Body string `json:"body"`
}

func (mail *Mail) FromJSON(msg []byte) error {
	return json.Unmarshal(msg, mail)
}

func (mail *Mail) ToJSON() []byte {
	str, _ := json.Marshal(mail)
	return str
}

// MailerUseCase represent the mailer's use cases
type MailerUseCase interface {
	SendMail(mail Mail) error
}
