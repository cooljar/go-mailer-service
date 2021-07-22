package main

import (
	"fmt"
	"github.com/cooljar/go-mailer-service/delivery"
	"github.com/cooljar/go-mailer-service/domain"
	"github.com/cooljar/go-mailer-service/usecase"
	"github.com/streadway/amqp"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

var err error
var smtpHost,smtpPort,smtpEmailAddress,smtpEmailPassword string
var rabbitMqDial, rabbitMqQueue string
var smtpPortInt int
var dialer *gomail.Dialer
var conn *amqp.Connection
var ch *amqp.Channel

func init()  {
	//SMTP env
	smtpHost = os.Getenv("SMTP_HOST")
	if smtpHost == "" {exitf("SMTP_HOST config is required")}
	smtpPort = os.Getenv("SMTP_PORT")
	if smtpPort == "" {exitf("SMTP_PORT config is required")}
	smtpPortInt,err = strconv.Atoi(smtpPort)
	if err != nil {
		exitf("error converting : ", err)
	}
	smtpEmailAddress = os.Getenv("SMTP_EMAIL_ADDRESS")
	if smtpEmailAddress == "" {exitf("SMTP_EMAIL_ADDRESS config is required")}
	smtpEmailPassword = os.Getenv("SMTP_EMAIL_PASSWORD")
	if smtpEmailPassword == "" {exitf("SMTP_EMAIL_PASSWORD config is required")}
	dialer = gomail.NewDialer(
		smtpHost,
		smtpPortInt,
		smtpEmailAddress,
		smtpEmailPassword,
	)

	//RabbitMQ env
	rabbitMqDial = os.Getenv("RABBIT_MQ_DIAL")
	if rabbitMqDial == "" {exitf("RABBIT_MQ_DIAL config is required")}
	rabbitMqQueue = os.Getenv("RABBIT_MQ_QUEUE")
	if rabbitMqQueue == "" {exitf("RABBIT_MQ_QUEUE config is required")}
}

func main()  {
	conn, err = amqp.Dial(rabbitMqDial)
	if err != nil {
		exitf("Failed Initializing Broker Connection: ", err)
	}

	ch, err = conn.Channel()
	if err != nil {
		exitf("Failed connecting to a channel: ", err)
	}
	defer ch.Close()

	// Consuming messages
	msgs, err := ch.Consume(
		rabbitMqQueue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		exitf("Failed consuming to a channel: ", err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)

			var mail domain.Mail
			err := mail.FromJSON(d.Body)
			if err != nil {
				fmt.Printf("Failed unmarshall mail: %s\n", err)
			}

			//kirim email
			mailerUsecase := usecase.NewMailerUsecase(dialer)
			mailHandler := delivery.NewMailHandler(mailerUsecase)
			err = mailHandler.Send(mail)
			if err != nil {
				fmt.Printf("Failed to send mail: %s\n", err)
			}

			if err == nil {
				fmt.Println("Mail successfully sent..!")
			}
		}
	}()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")
	fmt.Println(" [*] - Waiting for messages")
	<-forever
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

