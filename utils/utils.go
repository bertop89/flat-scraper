package utils

import (
	"flag"
	"log"
	"net/smtp"
	"time"
)


type Configuration struct {
  Email			string 
  Pass			string
  Areas			string
}

func LoadConfig() Configuration {

	var config Configuration

	flag.StringVar(&config.Email,"email","","Gmail address to send the results to")
	flag.StringVar(&config.Pass,"pass","","Password for that gmail account")
	flag.StringVar(&config.Areas,"areas","tetuan","Areas to scrape")

	flag.Parse()
	
	if config.Email == "" || config.Pass == "" {
		log.Fatal("Invalid email or password parameter")
	}
	return config
}

func EmailSend(c Configuration, body string) {
  from := c.Email
  pass := c.Pass
  to := c.Email

  msg := "From: " + from + "\n" +
    "To: " + to + "\n" +
    "Subject: "+ time.Now().String()+"\n\n" +
    body

  err := smtp.SendMail("smtp.gmail.com:587",
    smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
    from, []string{to}, []byte(msg))

  if err != nil {
    log.Printf("smtp error: %s", err)
    return
  }
  
}