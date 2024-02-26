package email

import (
	"context"

	"movido-media/controllers/billing"

	"gopkg.in/gomail.v2"
)

type emailController struct {
	mailer *gomail.Message
}

type EmailController interface {
	Sender(ctx context.Context, data billing.ContractDetail, sender string) error
}

func NewEmailController() EmailController {
	return &emailController{
		mailer: gomail.NewMessage(),
	}
}
