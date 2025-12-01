package queuehandlers

import (
	"context"
	"encoding/json"
	"fmt"

	mailports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/mail"
	queueports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/queue"
)

type RegisterEmailHandler struct {
	mailService mailports.MailService
}

var _ queueports.QueueHandler = (*RegisterEmailHandler)(nil)

func NewRegisterEmailHandler(mailService mailports.MailService) *RegisterEmailHandler {
	return &RegisterEmailHandler{mailService: mailService}
}

type RegisterEmailTask struct {
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func (h *RegisterEmailHandler) HandleMessage(ctx context.Context, msg queueports.QueueMessage) error {
	if msg.Headers["type"] != "register_email" {
		return nil
	}

	var task RegisterEmailTask
	if err := json.Unmarshal(msg.Body, &task); err != nil {
		return fmt.Errorf("invalid register_email payload: %w", err)
	}

	return h.mailService.SendPlain(task.Email, task.Subject, task.Body)
}
