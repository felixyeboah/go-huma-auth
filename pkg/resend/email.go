package resend

import (
	"fmt"
	"github.com/resend/resend-go/v2"
	"huma-auth/config"
	"log"
)

func SendVerificationEmail(email, userID, verificationToken, path string) error {
	// env config
	env, err := config.LoadConfig("../../")
	if err != nil {
		log.Fatal(err)
	}

	client := resend.NewClient(env.ResendApiKey)

	url := env.VerificationLink
	verificationLink := fmt.Sprintf(url+"/"+path+"?user_id=%s&token=%s", userID, verificationToken)

	subject := "Email Verification"
	body := fmt.Sprintf("Please verify your email using this link: %s", verificationLink)

	from := "noreply@misscookieghana.com"
	to := email

	emailParams := &resend.SendEmailRequest{
		From:    from,
		To:      []string{to},
		Subject: subject,
		Text:    body,
	}

	_, err = client.Emails.Send(emailParams)
	if err != nil {
		return err
	}

	return nil
}
