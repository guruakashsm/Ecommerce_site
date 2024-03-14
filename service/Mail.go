package service

import (
	"ecommerce/models"
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
	sender := NewGmailSender("GURUAKAKSH SM", "guruakash.ec20@bitsathy.ac.in", "snuk gatz ohoa agyt")
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

func SendOrderConformation(toEmail,price,totalAmount,dateofdelivey,id,noofitems string,address models.Address) {
	sender := NewGmailSender("GURUAKAKSH SM", "guruakash.ec20@bitsathy.ac.in", "snuk gatz ohoa agyt")
	subject := "Order Conformation"
	htmlTemplate := `
    <!DOCTYPE html>
    <html>
    <head>
    <title></title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <style type="text/css">
    
    body, table, td, a { -webkit-text-size-adjust: 100%; -ms-text-size-adjust: 100%; }
    table, td { mso-table-lspace: 0pt; mso-table-rspace: 0pt; }
    img { -ms-interpolation-mode: bicubic; }
    
    img { border: 0; height: auto; line-height: 100%; outline: none; text-decoration: none; }
    table { border-collapse: collapse !important; }
    body { height: 100% !important; margin: 0 !important; padding: 0 !important; width: 100% !important; }
    
    
    a[x-apple-data-detectors] {
        color: inherit !important;
        text-decoration: none !important;
        font-size: inherit !important;
        font-family: inherit !important;
        font-weight: inherit !important;
        line-height: inherit !important;
    }
    
    @media screen and (max-width: 480px) {
        .mobile-hide {
            display: none !important;
        }
        .mobile-center {
            text-align: center !important;
        }
    }
    div[style*="margin: 16px 0;"] { margin: 0 !important; }
    </style>
    <body style="margin: 0 !important; padding: 0 !important; background-color: #eeeeee;" bgcolor="#eeeeee">
    
    
    <div style="display: none; font-size: 1px; color: #fefefe; line-height: 1px; font-family: Open Sans, Helvetica, Arial, sans-serif; max-height: 0px; max-width: 0px; opacity: 0; overflow: hidden;">
    For what reason would it be advisable for me to think about business content? That might be little bit risky to have crew member like them. 
    </div>
    
    <table border="0" cellpadding="0" cellspacing="0" width="100%">
        <tr>
            <td align="center" style="background-color: #eeeeee;" bgcolor="#eeeeee">
            
            <table align="center" border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width:600px;">
                <tr>
                    <td align="center" valign="top" style="font-size:0; padding: 35px;" bgcolor="#F44336">
                   
                    <div style="display:inline-block; max-width:50%; min-width:100px; vertical-align:top; width:100%;">
                        <table align="left" border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width:300px;">
                            <tr>
                                <td align="left" valign="top" style="font-family: Open Sans, Helvetica, Arial, sans-serif; font-size: 36px; font-weight: 800; line-height: 48px;" class="mobile-center">
                                    <h1 style="font-size: 36px; font-weight: 800; margin: 0; color: #ffffff;">Anon</h1>
                                </td>
                            </tr>
                        </table>
                    </div>
                    
                    <div style="display:inline-block; max-width:50%; min-width:100px; vertical-align:top; width:100%;" class="mobile-hide">
                        <table align="left" border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width:300px;">
                            <tr>
                                <td align="right" valign="top" style="font-family: Open Sans, Helvetica, Arial, sans-serif; font-size: 48px; font-weight: 400; line-height: 48px;">
                                    <table cellspacing="0" cellpadding="0" border="0" align="right">
                                        <tr>
                                            <td style="font-family: Open Sans, Helvetica, Arial, sans-serif; font-size: 18px; font-weight: 400;">
                                                <p style="font-size: 18px; font-weight: 400; margin: 0; color: #ffffff;"><a href="#" target="_blank" style="color: #ffffff; text-decoration: none;">Shop &nbsp;</a></p>
                                            </td>
                                            <td style="font-family: Open Sans, Helvetica, Arial, sans-serif; font-size: 18px; font-weight: 400; line-height: 24px;">
                                                <a href="#" target="_blank" style="color: #ffffff; text-decoration: none;"><img src="https://img.icons8.com/color/48/000000/small-business.png" width="27" height="23" style="display: block; border: 0px;"/></a>
                                            </td>
                                        </tr>
                                    </table>
                                </td>
                            </tr>
                        </table>
                    </div>
                  
                    </td>
                </tr>
                <tr>
                    <td align="center" style="padding: 35px 35px 20px 35px; background-color: #ffffff;" bgcolor="#ffffff">
                    <table align="center" border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width:600px;">
                        <tr>
                            <td align="center" style="font-family: Open Sans, Helvetica, Arial, sans-serif; font-size: 16px; font-weight: 400; line-height: 24px; padding-top: 25px;">
                                <img src="https://img.icons8.com/carbon-copy/100/000000/checked-checkbox.png" width="125" height="120" style="display: block; border: 0px;" /><br>
                                <h2 style="font-size: 30px; font-weight: 800; line-height: 36px; color: #333333; margin: 0;">
                                    Thank You For Your Order!
                                </h2>
                            </td>
                        </tr>
                        <tr>
                            <td align="left" style="font-family: Open Sans, Helvetica, Arial, sans-serif; font-size: 16px; font-weight: 400; line-height: 24px; padding-top: 10px;">
                                <p style="font-size: 16px; font-weight: 400; line-height: 24px; color: #777777;">
                                    Thank you for ordering in our site. We will make sure your order reaches you in right time.
                                </p>
                            </td>
                        </tr>
                        <tr>
                            <td align="left" style="padding-top: 20px;">
                                <table cellspacing="0" cellpadding="0" border="0" width="100%">
                                    <tr>
                                        <td width="75%" align="left" bgcolor="#eeeeee" style="font-family: Open Sans, Helvetica, Arial, sans-serif; font-size: 16px; font-weight: 800; line-height: 24px; padding: 10px;">
                                            Order Confirmation #
                                        </td>
                                        <td width="25%" align="left" bgcolor="#eeeeee" style="font-family: Open Sans, Helvetica, Arial, sans-serif; font-size: 16px; font-weight: 800; line-height: 24px; padding: 10px;">
                                            `+id+`
                                        </td>
                                    </tr>
                                    <tr>
                                        <td width="75%" align="left" style="font-family: Open Sans, Helvetica, Arial, sans-serif; font-size: 16px; font-weight: 400; line-height: 24px; padding: 15px 10px 5px 10px;">
                                            Purchased Item (`+noofitems+`)
                                        </td>
                                        <td width="25%" align="left" style="font-family: Open Sans, Helvetica, Arial, sans-serif; font-size: 16px; font-weight: 400; line-height: 24px; padding: 15px 10px 5px 10px;">
                                            ₹`+price+`
                                        </td>
                                    </tr>
                                    <tr>
                                        <td width="75%" align="left" style="font-family: Open Sans, Helvetica, Arial, sans-serif; font-size: 16px; font-weight: 400; line-height: 24px; padding: 5px 10px;">
                                            Shipping + Handling
                                        </td>
                                        <td width="25%" align="left" style="font-family: Open Sans, Helvetica, Arial, sans-serif; font-size: 16px; font-weight: 400; line-height: 24px; padding: 5px 10px;">
                                            ₹50.00
                                        </td>
                                    </tr>
                                </table>
                            </td>
                        </tr>
                        <tr>
                            <td align="left" style="padding-top: 20px;">
                                <table cellspacing="0" cellpadding="0" border="0" width="100%">
                                    <tr>
                                        <td width="75%" align="left" style="font-family: Open Sans, Helvetica, Arial, sans-serif; font-size: 16px; font-weight: 800; line-height: 24px; padding: 10px; border-top: 3px solid #eeeeee; border-bottom: 3px solid #eeeeee;">
                                            TOTAL
                                        </td>
                                        <td width="25%" align="left" style="font-family: Open Sans, Helvetica, Arial, sans-serif; font-size: 16px; font-weight: 800; line-height: 24px; padding: 10px; border-top: 3px solid #eeeeee; border-bottom: 3px solid #eeeeee;">
                                            ₹`+totalAmount+`
                                        </td>
                                    </tr>
                                </table>
                            </td>
                        </tr>
                    </table>
                    
                    </td>
                </tr>
                 <tr>
                    <td align="center" height="100%" valign="top" width="100%" style="padding: 0 35px 35px 35px; background-color: #ffffff;" bgcolor="#ffffff">
                    <table align="center" border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width:660px;">
                        <tr>
                            <td align="center" valign="top" style="font-size:0;">
                                <div style="display:inline-block; max-width:50%; min-width:240px; vertical-align:top; width:100%;">
    
                                    <table align="left" border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width:300px;">
                                        <tr>
                                            <td align="left" valign="top" style="font-family: Open Sans, Helvetica, Arial, sans-serif; font-size: 16px; font-weight: 400; line-height: 24px;">
                                                <p style="font-weight: 800;">Delivery Address</p>
                                                <p>`+address.FirstName + address.LastName +`<br>`+address.Street_Name+`<br>`+address.City + string(address.Pincode)+`</p>
                                            </td>
                                        </tr>
                                    </table>
                                </div>
                                <div style="display:inline-block; max-width:50%; min-width:240px; vertical-align:top; width:100%;">
                                    <table align="left" border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width:300px;">
                                        <tr>
                                            <td align="left" valign="top" style="font-family: Open Sans, Helvetica, Arial, sans-serif; font-size: 16px; font-weight: 400; line-height: 24px;">
                                                <p style="font-weight: 800;">Estimated Delivery Date</p>
                                                <p>`+dateofdelivey+`</p>
                                            </td>
                                        </tr>
                                    </table>
                                </div>
                            </td>
                        </tr>
                    </table>
                    </td>
                </tr>
                <tr>
                    <td align="center" style=" padding: 35px; background-color: #ff7361;" bgcolor="#1b9ba3">
                    <table align="center" border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width:600px;">
                        <tr>
                            <td align="center" style="font-family: Open Sans, Helvetica, Arial, sans-serif; font-size: 16px; font-weight: 400; line-height: 24px; padding-top: 25px;">
                                <h2 style="font-size: 24px; font-weight: 800; line-height: 30px; color: #ffffff; margin: 0;">
                                    Buy more before it gets out of stock.
                                </h2>
                            </td>
                        </tr>
                        <tr>
                            <td align="center" style="padding: 25px 0 15px 0;">
                                <table border="0" cellspacing="0" cellpadding="0">
                                    <tr>
                                        <td align="center" style="border-radius: 5px;" bgcolor="#66b3b7">
                                          <a href="#" target="_blank" style="font-size: 18px; font-family: Open Sans, Helvetica, Arial, sans-serif; color: #ffffff; text-decoration: none; border-radius: 5px; background-color: #F44336; padding: 15px 30px; border: 1px solid #F44336; display: block;">Shop Again</a>
                                        </td>
                                    </tr>
                                </table>
                            </td>
                        </tr>
                    </table>
                    </td>
                </tr>
                <tr>
                    <td align="center" style="padding: 35px; background-color: #ffffff;" bgcolor="#ffffff">
                    <table align="center" border="0" cellpadding="0" cellspacing="0" width="100%" style="max-width:600px;">
                        <tr>
                            <td align="center">
                                <img src="logo-footer.png" width="37" height="37" style="display: block; border: 0px;"/>
                            </td>
                        </tr>
                        <tr>
                            <td align="center" style="font-family: Open Sans, Helvetica, Arial, sans-serif; font-size: 14px; font-weight: 400; line-height: 24px; padding: 5px 0 10px 0;">
                                <p style="font-size: 14px; font-weight: 800; line-height: 18px; color: #333333;">
                                    Anon<br>
                                    Sathy, Erode - 638402
                                </p>
                            </td>
                        </tr>
                        <tr>
                            <td align="left" style="font-family: Open Sans, Helvetica, Arial, sans-serif; font-size: 14px; font-weight: 400; line-height: 24px;">
                                <p style="font-size: 14px; font-weight: 400; line-height: 20px; color: #777777;">
                                    If you didn't create an account using this email address, please ignore this email.
                                </p>
                            </td>
                        </tr>
                    </table>
                    </td>
                </tr>
            </table>
            </td>
        </tr>
    </table>
        
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
