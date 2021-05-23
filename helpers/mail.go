package helpers

import (
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendMail(recipientEmail string, firstname string, url string) {

	from := mail.NewEmail("Developer-Justice","pistischaris494@gmail.com")
	subject := "Email Verification"
	to := mail.NewEmail(firstname, recipientEmail)
	htmlContent := "<p>Hi "+firstname+"</p>"+"<p>You recently signed up for GitBank's Online Banking.</p>"+"<p>Please click <b><a href="+url+">here</a></b> to verify your email address</p>"
	message := mail.NewSingleEmail(from, subject, to,"", htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	if err !=nil{
		log.Println(err)
	}

	log.Println(response.StatusCode)
	log.Println(response.Body)
}