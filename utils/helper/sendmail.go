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

type Data struct {
	name  string
	email string
	pin   string
}

func Corn(data Data) {
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron.New(cron.WithLocation(jakartaTime))

	defer scheduler.Stop()
	scheduler.AddFunc("@every 10s", func() { sendEmail(data) })
	go scheduler.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

}

func sendEmail(data Data) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetAddressHeader("Cc", data.email, data.name)
	mailer.SetHeader("Subject", "Activate Account")
	mailer.SetBody("text/html", "Please !!! Activate your acount")
	mailer.SetBody("text/html", data.pin)

	dialer := gomail.NewDialer(CONFIG_SMTP_HOST, 587, "klukmanul33@gmail.com", os.Getenv("CONFIG_AUTH_PASSWORD"))

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Error(err.Error())
	}

	loo.Println("Mail sent!")
}
