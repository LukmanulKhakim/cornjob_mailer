package helper

import (
	loo "log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/gommon/log"
	cron "github.com/robfig/cron/v3"
	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Corn Job <klukmanul33@gmail.com>"

func Corn(name, email, pin string) error {
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron.New(cron.WithLocation(jakartaTime))

	go scheduler.Start()
	scheduler.AddFunc("@every 10s", func() { sendEmail(name, email, pin) })
	defer scheduler.Stop()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	return nil
}

func sendEmail(name, email, pin string) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetAddressHeader("Cc", email, name)
	mailer.SetHeader("Subject", "Activate Account")
	mailer.SetBody("text/html", "Please !!! Activate your acount")
	mailer.SetBody("text/html", pin)

	dialer := gomail.NewDialer(CONFIG_SMTP_HOST, 587, "klukmanul33@gmail.com", os.Getenv("CONFIG_AUTH_PASSWORD"))

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Error(err.Error())
	}

	loo.Println("Mail sent!")
}
