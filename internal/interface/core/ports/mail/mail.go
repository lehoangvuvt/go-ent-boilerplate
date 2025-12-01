package mailports

type MailService interface {
	SendPlain(to, subject, body string) error
	SendHTML(to, subject, html string) error
	SendTemplate(to, templateName string, data any) error
}
