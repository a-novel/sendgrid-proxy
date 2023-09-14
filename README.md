# Sendgrid proxy

Sendgrid abstraction layer for the A-Novel apis.

```bash
go get github.com/a-novel/sendgrid-proxy
```

## Usage

```go
package main

import (
    "github.com/a-novel/sendgrid-proxy"
    "github.com/rs/zerolog"
    "github.com/sendgrid/sendgrid-go/helpers/mail"
    "os"
)

func main() {
    ctx := context.Background()
    sendgridKey := os.Getenv("SENDGRID_API_KEY")
    logger := zerolog.New(os.Stdout)
    sender := mail.NewEmail("My App", "noreply@my-app.com")
    
    proxy := sendgridproxy.NewMailer(sendgridKey, sender, false, logger)
    
    user := GetUser()
    userEmail := mail.NewEmail(user.Name, user.Email)
    
    // Send an email
    err := proxy.Send(ctx, userEmail, mailTemplateID, mailTemplateData)
}
```
