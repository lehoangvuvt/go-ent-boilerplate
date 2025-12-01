package resendmail

import (
	mailports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/mail"
	"github.com/resend/resend-go/v2"
)

type ResendMailService struct {
	client *resend.Client
}

var _ mailports.MailService = (*ResendMailService)(nil)

func NewResendMailService(apiKey string) *ResendMailService {
	client := resend.NewClient(apiKey)
	return &ResendMailService{client}
}

func (s *ResendMailService) SendPlain(to, subject, body string) error {
	params := &resend.SendEmailRequest{
		From:    "onboarding@resend.dev",
		To:      []string{to},
		Subject: subject,
		Html:    body,
	}
	_, err := s.client.Emails.Send(params)
	if err != nil {
		return err
	}
	return nil
}

func (s *ResendMailService) SendHTML(to, subject, html string) error { return nil }

func (s *ResendMailService) SendTemplate(to, templateName string, data any) error { return nil }
