package service

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"time"
)

func StartSending() {

	go func() {
		for {

			rate, err := GetRate()
			if err != nil {
				log.Println("Failed to get rate:", err)
				return
			}

			dsn := os.Getenv("DSN")
			db, err := sqlx.Open("mysql", dsn)
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			repo := NewRepository(db)
			selectEmails(repo, rate)
			time.Sleep(24 * time.Hour)
		}
	}()
}

func selectEmails(repo *Repository, rate float64) {
	emails, err := repo.GetAllSubscribers()
	if err != nil {
		log.Printf("Failed to get emails: %v", err)
		return
	}

	if err := sendEmail(emails, rate); err != nil {
		log.Printf("Failed to send email %v", err)
	}

}

func sendEmail(to []string, rate float64) error {
	m := gomail.NewMessage()
	hostEmail := os.Getenv("HOST_EMAIL")
	serviceEmail := os.Getenv("SERVICE_EMAIL")
	serviceEmailPasswd := os.Getenv("SERVICE_EMAIL_PASSWD")
	m.SetHeader("From", serviceEmail)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", "Today's exchange rate news")
	m.SetBody("text/plain", "Today's exchange rate: "+fmt.Sprintf("%.4f", rate))

	d := gomail.NewDialer(hostEmail, 587, serviceEmail, serviceEmailPasswd)

	return d.DialAndSend(m)
}
