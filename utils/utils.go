package utils

import (
	"flag"
	"log"
	"net/smtp"
  "strings"
)


type Configuration struct {
  Email			string 
  Pass			string
  Areas			[]string
  Size      int
  Price     int
}



func LoadConfig() Configuration {

	var config Configuration

  var tempAreas string

	flag.StringVar(&config.Email,"email","","Gmail address to send the results to")
	flag.StringVar(&config.Pass,"pass","","Password for that gmail account")
	flag.StringVar(&tempAreas,"areas","tetuan","Areas to scrape")
  flag.IntVar(&config.Size,"size",40,"Size of the flat")
  flag.IntVar(&config.Price,"price",650,"Max price of the rent")

	flag.Parse()

  config.Areas = strings.Split(tempAreas,",")
	
	if config.Email == "" || config.Pass == "" {
		log.Fatal("Invalid email or password parameter")
	}
	return config
}

func EmailSend(c Configuration, len string, body string) {
  from := c.Email
  pass := c.Pass
  to := c.Email

  msg := "From: " + from + "\n" +
    "To: " + to + "\n" +
    "Subject: "+len+" results found \n\n" +
    body

  err := smtp.SendMail("smtp.gmail.com:587",
    smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
    from, []string{to}, []byte(msg))

  if err != nil {
    log.Printf("smtp error: %s", err)
    return
  }
  
}