package service

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
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

func SendEmailforCustomerVerification(toEmail, id, name string) {
	sender := NewGmailSender("GURUAKAKSH SM", "guruakash.ec20@bitsathy.ac.in",  "snuk gatz ohoa agyt")
	subject := "A test Email"
	htmlTemplate := `
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Email Verification</title>
            <style>
                body {
                    font-family: 'Arial', sans-serif;
                    background-color: #f4f4f4;
                    margin: 0;
                    padding: 0;
                    text-align: center;
                }

                .container {
                    max-width: 600px;
                    margin: 20px auto;
                    background-color: #fff;
                    padding: 20px;
                    border-radius: 5px;
                    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
                }

                h1 {
                    color: #333;
                }

                p {
                    color: #666;
                    margin-bottom: 20px;
                }

                .otp-container {
                    background-color: #e0e0e0;
                    padding: 10px;
                    border-radius: 5px;
                }

                .otp {
                    font-size: 24px;
                    font-weight: bold;
                    color: #333;
                }

                .note {
                    color: #666;
                    font-size: 14px;
                    margin-top: 10px;
                }

                .footer {
                    margin-top: 20px;
                    color: #999;
                    font-size: 12px;
                }
            </style>
                </head>
            <body>
            <div class="container">
                <h1>Email Verification</h1>
                <p>Hello <strong>`
	htmlTemplate += name
	htmlTemplate += `</strong>,</p>
                <p>Your OTP for email verification is:</p>
                <div class="otp-container">
                    <span class="otp">`
	htmlTemplate += id
	htmlTemplate += `</span>
                </div>
                <p class="note">Please use this OTP to complete the verification process.</p>
                <p>If you did not request this verification, please ignore this email.</p>
                <p class="footer">&copy; 2024 Anon. All rights reserved.</p>
            </div>
        </body>
        </html>
        `
	to := []string{toEmail}

	err := sender.SendEmail(subject, htmlTemplate, to, nil, nil)

	if err != nil {
		fmt.Println("Error sending email:", err)
	} else {
		fmt.Println("Email sent successfully!")
	}

}

func SendThankYouEmail(toEmail, username string) {
	sender := NewGmailSender("GURUAKAKSH SM", "guruakash.ec20@bitsathy.ac.in", "snuk gatz ohoa agyt")
	subject := "Thank You for Creating an Account"
	htmlTemplate := `
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Thank You for Creating an Account</title>
            <style>
                body {
                    font-family: 'Arial', sans-serif;
                    background-color: #f4f4f4;
                    margin: 0;
                    padding: 0;
                    text-align: center;
                }

                .container {
                    max-width: 600px;
                    margin: 20px auto;
                    background-color: #fff;
                    padding: 20px;
                    border-radius: 5px;
                    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
                }

                h1 {
                    color: #333;
                }

                p {
                    color: #666;
                    margin-bottom: 20px;
                }

                .note {
                    color: #666;
                    font-size: 14px;
                    margin-top: 10px;
                }

                .footer {
                    margin-top: 20px;
                    color: #999;
                    font-size: 12px;
                }
            </style>
        </head>
        <body>
            <div class="container">
                <h1>Thank You for Creating an Account</h1>
                <p>Hello <strong>` + username + `</strong>,</p>
                <p>We are delighted to inform you that your account has been successfully created on our site.</p>
                <p class="note">Thank you for choosing our services.</p>
                <p class="footer">&copy; 2024 Anon. All rights reserved.</p>
            </div>
        </body>
        </html>
    `
	to := []string{toEmail}

	err := sender.SendEmail(subject, htmlTemplate, to, nil, nil)

	if err != nil {
		fmt.Println("Error sending email:", err)
	} else {
		fmt.Println("Email sent successfully!")
	}
}

func SendSellerInvitation(toEmail, name, password, siteURL string) {
	sender := NewGmailSender("GURUAKAKSH SM", "guruakash.ec20@bitsathy.ac.in", "snuk gatz ohoa agyt")
	subject := "Invitation to Join Our Platform"
	htmlTemplate := `
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Invitation to Join Our Platform</title>
            <style>
                body {
                    font-family: 'Arial', sans-serif;
                    background-color: #f4f4f4;
                    margin: 0;
                    padding: 0;
                    text-align: center;
                }

                .container {
                    max-width: 600px;
                    margin: 20px auto;
                    background-color: #fff;
                    padding: 20px;
                    border-radius: 5px;
                    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
                }

                h1 {
                    color: #333;
                }

                p {
                    color: #666;
                    margin-bottom: 20px;
                }

                .credentials {
                    background-color: #e0e0e0;
                    padding: 10px;
                    border-radius: 5px;
                    margin-top: 20px;
                }

                .label {
                    font-weight: bold;
                    color: #333;
                }

                .value {
                    color: #666;
                }

                .note {
                    color: #666;
                    font-size: 14px;
                    margin-top: 10px;
                }

                .footer {
                    margin-top: 20px;
                    color: #999;
                    font-size: 12px;
                }
            </style>
        </head>
        <body>
            <div class="container">
                <h1>Invitation to Join Our Platform</h1>
                <p>Hello <strong>` + name + `</strong>,</p>
                <p>Congratulations! You have been invited to join our platform as a seller.</p>
                <div class="credentials">
                    <p class="label">Email:</p>
                    <p class="value">` + toEmail + `</p>
                    <p class="label">Password:</p>
                    <p class="value">` + password + `</p>
                    <p class="label">Visit Our Site:</p>
                    <p class="value"><a href="` + siteURL + `" target="_blank">` + siteURL + `</a></p>
                </div>
                <p class="note">Please use the provided credentials to log in and start using our services.</p>
                <p class="footer">&copy; 2024 Anon. All rights reserved.</p>
            </div>
        </body>
        </html>
    `
	to := []string{toEmail}

	err := sender.SendEmail(subject, htmlTemplate, to, nil, nil)

	if err != nil {
		fmt.Println("Error sending email:", err)
	} else {
		fmt.Println("Email sent successfully!")
	}
}

func SendAdminInvitation(toEmail, name, password, siteURL, ipAddress, totpSecret string) {
    sender := NewGmailSender("GURUAKAKSH SM", "guruakash.ec20@bitsathy.ac.in", "snuk gatz ohoa agyt")
    subject := "Invitation to Join Our Platform"
    htmlTemplate := `
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Invitation to Join Our Platform</title>
            <style>
                body {
                    font-family: 'Arial', sans-serif;
                    background-color: #f4f4f4;
                    margin: 0;
                    padding: 0;
                    text-align: center;
                }

                .container {
                    max-width: 600px;
                    margin: 20px auto;
                    background-color: #fff;
                    padding: 20px;
                    border-radius: 5px;
                    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
                }

                h1 {
                    color: #333;
                }

                p {
                    color: #666;
                    margin-bottom: 20px;
                }

                .credentials {
                    background-color: #e0e0e0;
                    padding: 10px;
                    border-radius: 5px;
                    margin-top: 20px;
                }

                .label {
                    font-weight: bold;
                    color: #333;
                }

                .value {
                    color: #666;
                }

                .note {
                    color: #666;
                    font-size: 14px;
                    margin-top: 10px;
                }

                .footer {
                    margin-top: 20px;
                    color: #999;
                    font-size: 12px;
                }

                .disclaimer {
                    margin-top: 20px;
                    color: #ff0000;
                    font-size: 14px;
                    font-weight: bold;
                }
            </style>
        </head>
        <body>
            <div class="container">
                <h1>Invitation to Join Our Platform</h1>
                <p>Hello <strong>` + name + `</strong>,</p>
                <p>Congratulations! You have been invited to join our platform as a Admin.</p>
                <div class="credentials">
                    <p class="label">Email:</p>
                    <p class="value">` + toEmail + `</p>
                    <p class="label">Password:</p>
                    <p class="value">` + password + `</p>
                    <p class="label">IP Address:</p>
                    <p class="value">` + ipAddress + `</p>
                    <p class="label">Login URL:</p>
                    <p class="value"><a href="` + siteURL + `" target="_blank">` + siteURL + `</a></p>
                    <p class="label">TOTP Secret:</p>
                    <p class="value">` + totpSecret + `</p>
                </div>
                <p class="note">Please use the provided credentials to log in and start using our services.</p>
                <div class="disclaimer">
                    <p>**Disclaimer: Do not forward or share this email as it contains sensitive information.**</p>
                </div>
                <p class="footer">&copy; 2024 Anon. All rights reserved.</p>
            </div>
        </body>
        </html>
    `
    to := []string{toEmail}

    err := sender.SendEmail(subject, htmlTemplate, to, nil, nil)

    if err != nil {
        fmt.Println("Error sending email:", err)
    } else {
        fmt.Println("Email sent successfully!")
    }
}


func SendBlockingNotification(toEmail, name, blockingReason string) {
    sender := NewGmailSender("GURUAKAKSH SM", "guruakash.ec20@bitsathy.ac.in", "snuk gatz ohoa agyt")
    subject := "Account Blocked Notification"
    htmlTemplate := `
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Account Blocked Notification</title>
            <style>
                body {
                    font-family: 'Arial', sans-serif;
                    background-color: #f4f4f4;
                    margin: 0;
                    padding: 0;
                    text-align: center;
                }

                .container {
                    max-width: 600px;
                    margin: 20px auto;
                    background-color: #fff;
                    padding: 20px;
                    border-radius: 5px;
                    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
                }

                h1 {
                    color: #333;
                }

                p {
                    color: #666;
                    margin-bottom: 20px;
                }

                .notification {
                    background-color: #e0e0e0;
                    padding: 10px;
                    border-radius: 5px;
                    margin-top: 20px;
                }

                .label {
                    font-weight: bold;
                    color: #333;
                }

                .value {
                    color: #666;
                }

                .note {
                    color: #666;
                    font-size: 14px;
                    margin-top: 10px;
                }

                .footer {
                    margin-top: 20px;
                    color: #999;
                    font-size: 12px;
                }
            </style>
        </head>
        <body>
            <div class="container">
                <h1>Account Blocked Notification</h1>
                <p>Hello <strong>` + name + `</strong>,</p>
                <div class="notification">
                    <p class="label">Account Status:</p>
                    <p class="value">Your account has been blocked.</p>
                    <p class="label">Reason:</p>
                    <p class="value">` + blockingReason + `</p>
                </div>
                <p class="note">If you believe this is an error or have any questions, please contact our support team.</p>
                <p class="footer">&copy; 2024 Anon. All rights reserved.</p>
            </div>
        </body>
        </html>
    `
    to := []string{toEmail}

    err := sender.SendEmail(subject, htmlTemplate, to, nil, nil)

    if err != nil {
        fmt.Println("Error sending email:", err)
    } else {
        fmt.Println("Email sent successfully!")
    }
}


func SendEditDataNotification(toEmail, fieldUpdated, newValue string) {
    sender := NewGmailSender("GURUAKAKSH SM", "guruakash.ec20@bitsathy.ac.in", "snuk gatz ohoa agyt")
    subject := "Data Edit Notification"
    htmlTemplate := `
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Data Edit Notification</title>
            <style>
                body {
                    font-family: 'Arial', sans-serif;
                    background-color: #f4f4f4;
                    margin: 0;
                    padding: 0;
                    text-align: center;
                }

                .container {
                    max-width: 600px;
                    margin: 20px auto;
                    background-color: #fff;
                    padding: 20px;
                    border-radius: 5px;
                    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
                }

                h1 {
                    color: #333;
                }

                p {
                    color: #666;
                    margin-bottom: 20px;
                }

                .notification {
                    background-color: #e0e0e0;
                    padding: 10px;
                    border-radius: 5px;
                    margin-top: 20px;
                }

                .label {
                    font-weight: bold;
                    color: #333;
                }

                .value {
                    color: #666;
                }

                .note {
                    color: #666;
                    font-size: 14px;
                    margin-top: 10px;
                }

                .footer {
                    margin-top: 20px;
                    color: #999;
                    font-size: 12px;
                }
            </style>
        </head>
        <body>
            <div class="container">
                <h1>Data Edit Notification</h1>
                <p>Hello,</p>
                <div class="notification">
                    <p class="label">Field Updated:</p>
                    <p class="value">` + fieldUpdated + `</p>
                    <p class="label">New Value:</p>
                    <p class="value">` + newValue + `</p>
                </div>
                <p class="note">This is a notification regarding the recent update of data.</p>
                <p class="footer">&copy; 2024 Anon. All rights reserved.</p>
            </div>
        </body>
        </html>
    `
    to := []string{toEmail}

    err := sender.SendEmail(subject, htmlTemplate, to, nil, nil)

    if err != nil {
        fmt.Println("Error sending email:", err)
    } else {
        fmt.Println("Email sent successfully!")
    }
}



