# Mailer Service
Package cooljar/go-mailer-service Implements the mailing service using Go and RabbitMQ message broker.

## Getting Started
These instructions will get you a copy of the project up and running on docker container and on your local machine.

### Prerequisites
Prequisites package:
* [Docker](https://www.docker.com/get-started) - for developing, shipping, and running applications (Application Containerization).
* [Go](https://golang.org/) - Go Programming Language
* [Gomail](https://github.com/go-gomail/gomail) - simple and efficient package to send emails
* [RabbitMQ](https://www.rabbitmq.com) - one of the most popular open source message brokers
* [Make](https://golang.org/) - Automated Execution using Makefile
* Email for SMTP. If you use gmail, make sure [Less Secure App Setting](https://support.google.com/a/answer/6260879) is enabled.

Optional package:
* [gosec](https://github.com/securego/gosec) Golang Security Checker. Inspects source code for security problems by scanning the Go AST

### Setup Message Broker (RabbitMQ)
First things first, we’ll need an instance of RabbitMQ that we can interact with and play about with. 
The quickest approach is to use the docker run command and specifying the image name that we want to run:
```bash
$ docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
```
This will kick off our local RabbitMQ instance which we can manage through the UI which is available at 
[http://localhost:15672](http://localhost:15672) using the username guest and password guest. See [https://hub.docker.com/_/rabbitmq](https://hub.docker.com/_/rabbitmq) for more detailed information.
<br>
Create new Queue under `Queues` menu. I create a new queue called `q-mailer` for this package, 
you must include it on your `Makefile` and `run.sh` file.
<br>
See [https://github.com/cooljar/go-rabbitmq](https://github.com/cooljar/go-rabbitmq) for more detailed instruction about RabbitMQ implementation in go.

### Changing Email Template
* Open email layout template in `assets/html/layout/mail.html` and fill it with your setting.
* Open email content template in `assets/html/mail_content.html` and fill it with your setting.

### Running On Local Machine
Below is the instructions to run this project on your local machine:
1. Rename `run.sh.example` to `run.sh` and fill it with your environment values.
2. Open new `terminal`.
3. Set `run.sh` file permission.
```bash
$ chmod +x ./run.sh
```
4. Run application from terminal by using following command:
```bash
$ ./run.sh
```

### Running On Docker Container
1. Rename `Makefile.example` to `Makefile` and fill it with your make setting.
2. Run project by using following command:
```bash
$ make run

# Process:
#   - Build and run Docker containers Cooljar Mailer
```
Stop application by using following command:
```bash
$ make stop

# Process:
#   - Stop and remove Cooljar Mailer container
#   - remove Cooljar Mailer image
```

### Sending Email
* Open RabbitMQ management GUI at `http://localhost:15672` using the username `guest` and password `guest`.
* Go to `Queues` tab, then send your email inside `Publish Message` menu.
    * Let the `Headers` field empty.
    * Fill `Properties` key=>value field with `content_type=>text/plain`
    * Under the `Payload` field, fill it with your email message. Format must in Json format `{"sender_name":"Mailer App","to":"some@mail.com","subject":"Test Email","body":"Hi there, this is email body."}`.

## Testing
- Inspects source code for security problems using [gosec](https://github.com/securego/gosec). You need to install it first.
- Execute unit test by using following command:
```bash
$ make test
```

## Built With
* [Go](https://golang.org/) - Go Programming Languange
* [Go Modules](https://github.com/golang/go/wiki/Modules) - Go Dependency Management System
* [Make](https://www.gnu.org/software/make/) - GNU Make Automated Execution
* [Docker](https://www.docker.com/) - Application Containerization

## Authors
* **Fajar Rizky** - *Initial Work* - [cooljar](https://github.com/cooljar)

## More
-------