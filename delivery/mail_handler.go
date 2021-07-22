package delivery

import (
	"github.com/cooljar/go-mailer-service/domain"
)

type mailHandler struct {
	mailerUsecase domain.MailerUseCase
}

func NewMailHandler(mailerUsecase domain.MailerUseCase) *mailHandler {
	return &mailHandler{mailerUsecase: mailerUsecase}
}

func (mh *mailHandler) Send(mail domain.Mail) error {
	return mh.mailerUsecase.SendMail(mail)
}
