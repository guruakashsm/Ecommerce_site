package service

import (
    "fmt"
    "github.com/jordan-wright/email"
    "net/smtp"
)

const (
    sntpAuthAddress  = "smtp.gmail.com"
    sntpSeverAddress = "smtp.gmail.com:587"
)

type EmailSender interface {
    SendEmail(
        subject string,
        content string,
        to []string,
        cc []string,
        bcc []string,
    ) error
}

type GmailSender struct {
    name              string
    fromEmailAddress  string
    fromEmailPassword string
}

func NewGmailSender(name string, fromEmailAddress string, fromEmailPassword string) EmailSender {
    return &GmailSender{
        name:              name,
        fromEmailAddress:  fromEmailAddress,
        fromEmailPassword: fromEmailPassword,
    }
}

func (sender *GmailSender) SendEmail(
    subject string,
    content string,
    to []string,
    cc []string,
    bcc []string,
) error {
    e := email.NewEmail()
    e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress)
    e.Subject = subject
    e.HTML = []byte(content)
    e.To = to
    e.Cc = cc
    e.Bcc = bcc

    sntpAuth := smtp.PlainAuth("", sender.fromEmailAddress, sender.fromEmailPassword, sntpAuthAddress)
    return e.Send(sntpSeverAddress, sntpAuth)
}

func SendEmailforCustomerVerification(toEmail,id,name string ) {
        sender := NewGmailSender("GURUAKAKSH SM", "guruakash.ec20@bitsathy.ac.in", "")
        subject := "A test Email"
        content := `<h1>Hello</h1><p>`
		content +=id
		content +=`</p>`
        to := []string{toEmail}
    
        err := sender.SendEmail(subject, content, to, nil, nil)
    
        if err != nil {
            fmt.Println("Error sending email:", err)
        } else {
            fmt.Println("Email sent successfully!")
        }

}