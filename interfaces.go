package sendgridproxy

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Mailer is an interface to send emails from dynamic templates, through Sendgrid.
type Mailer interface {
	Send(ctx context.Context, recipient *mail.Email, templateID string, data map[string]interface{}) error
}

// NewMailer returns a new implementation of Mailer.
//
// APIKey is a secret key to access Sendgrid API. It must not be pushed to git.
// Sandbox flag logs mail in the terminal rather than sending them, for local development.
func NewMailer(apiKey string, sender *mail.Email, sandbox bool, logger zerolog.Logger) Mailer {
	return &mailerImpl{
		apiKey:  apiKey,
		sandbox: sandbox,
		from:    sender,
		logger:  logger,
	}
}
