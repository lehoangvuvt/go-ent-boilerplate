package smtpmail

import mailports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/mail"

type SMTPMailService struct {
	host string
	port int
	user string
	pass string
}

var _ mailports.MailService = (*SMTPMailService)(nil)

func NewSMTPMailService(host string, port int, user, pass string) *SMTPMailService {
	return &SMTPMailService{host, port, user, pass}
}

func (s *SMTPMailService) SendPlain(to, subject, body string) error { return nil }

func (s *SMTPMailService) SendHTML(to, subject, html string) error { return nil }

func (s *SMTPMailService) SendTemplate(to, templateName string, data any) error { return nil }
