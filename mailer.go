package sendgridproxy

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type mailerImpl struct {
	apiKey  string
	sandbox bool
	from    *mail.Email
	logger  zerolog.Logger
}

func (mailer *mailerImpl) Send(ctx context.Context, recipient *mail.Email, templateID string, data map[string]interface{}) error {
	message := mail.NewV3Mail()
	personalization := mail.NewPersonalization()

	// Set senders and recipients.
	message.SetFrom(mailer.from)

	// NOTE: for now, mailer is not tailored for email campaigns. The only emails sent are for account management (such
	// as validating email or resetting password), so delivery is an absolute priority. It is thus acceptable to
	// bypass all inbox firewalls, especially since this method is only able to send one email at a time (which is
	// greatly inconvenient for mailing campaigns).
	message.MailSettings = &mail.MailSettings{BypassListManagement: mail.NewSetting(true)}
	message.MailSettings.SandboxMode = mail.NewSetting(mailer.sandbox)

	// Set email content.
	message.SetTemplateID(templateID)
	for k, v := range data {
		personalization.SetDynamicTemplateData(k, v)
	}
	message.AddPersonalizations(personalization)

	sendClient := sendgrid.NewSendClient(mailer.apiKey)
	// If sandbox mode is enabled, sendgrid will just parse the template, without actually sending the email.
	response, err := sendClient.SendWithContext(ctx, message)

	// Print email in sandbox mode, for testing and debugging.
	if mailer.sandbox {
		mailer.logger.Debug().
			Int("status", response.StatusCode).
			Str("recipient", recipient.Address).
			Str("alias", recipient.Name).
			Msgf(string(mail.GetRequestBody(message)))
	}

	return err
}
