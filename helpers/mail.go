package helpers

import (
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendMail(recipientEmail string, firstname string, url string) {

	from := mail.NewEmail("Developer-Justice", "pistischaris494@gmail.com")
	subject := "Email Verification"
	to := mail.NewEmail(firstname, recipientEmail)
	htmlContent := "<p>Hi " + firstname + "</p>" + "<p>You recently signed up for GitBank's Online Banking.</p>" + "<p>Please click <b><a href=" + url + ">here</a></b> to verify your email address</p>"
	message := mail.NewSingleEmail(from, subject, to, "", htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	if err != nil {
		log.Println(err)
	}

	log.Println(response.StatusCode)
	log.Println(response.Body)
}

func SendAccountOpenMail(recipientEmail string, firstname string, accountnumber string, accounttype string) {
	from := mail.NewEmail("Developer-Justice", "pistischaris494@gmail.com")
	subject := "Welcome to GitBank"
	to := mail.NewEmail(firstname, recipientEmail)
	htmlContent := "<p>Hello " + firstname + "</p>" + "<p>You recently opened a " + accounttype + " account with GitBank. We hope you will enjoy our services.</p>" + "<p>Your account number is " + "<b>" + accountnumber + "</b>" + "</p>" + "<p>Cheers,<br />GitBank Team</p>"
	message := mail.NewSingleEmail(from, subject, to, "", htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	if err != nil {
		log.Println(err)
	}

	log.Println(response.StatusCode)
	log.Println(response.Body)
}

func SendAccountCloseMail(recipientEmail string, firstname string) {
	from := mail.NewEmail("Developer-Justice", "pistischaris494@gmail.com")
	subject := "GitBank Account Closed"
	to := mail.NewEmail(firstname, recipientEmail)
	htmlContent := "<p>Hello " + firstname + "</p>" + "<p>Your GitBank account has been closed successfully. To reopen, please contact the management by sending a mail to this email</p>" + "<p>Cheers,<br />GitBank Team.</p>"
	message := mail.NewSingleEmail(from, subject, to, "", htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	if err != nil {
		log.Println(err)
	}

	log.Println(response.StatusCode)
	log.Println(response.Body)
}

func SendAccountReActivateMail(recipientEmail string, firstname string) {
	from := mail.NewEmail("Developer-Justice", "pistischaris494@gmail.com")
	subject := "GitBank Account Re-Activation"
	to := mail.NewEmail(firstname, recipientEmail)
	htmlContent := "<p>Hello " + firstname + "</p>" + "<p>Your GitBank account has been successfully re-activated!. We are glad to have you back!!</p>" + "<p>Cheers,<br />GitBank Team.</p>"
	message := mail.NewSingleEmail(from, subject, to, "", htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	if err != nil {
		log.Println(err)
	}

	log.Println(response.StatusCode)
	log.Println(response.Body)
}
