package mailerService

type IMailerService interface {
	SendEmail(emailAddress, token string) (bool)
}