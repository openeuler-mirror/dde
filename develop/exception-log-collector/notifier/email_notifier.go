package notifier

import (
    "strconv"

    "gopkg.in/gomail.v2"
)

type EmailNotifier struct {
    smtpServer string
    smtpPort   int
    username   string
    password   string
    from       string
    to         string
}

func NewEmailNotifier(config map[string]interface{}) *EmailNotifier {
    port, _ := strconv.Atoi(config["smtp_port"].(string))
    return &EmailNotifier{
        smtpServer: config["smtp_server"].(string),
        smtpPort:   port,
        username:   config["username"].(string),
        password:   config["password"].(string),
        from:       config["from"].(string),
        to:         config["to"].(string),
    }
}

func (en *EmailNotifier) Notify(message string) error {
    m := gomail.NewMessage()
    m.SetHeader("From", en.from)
    m.SetHeader("To", en.to)
    m.SetHeader("Subject", "Exception Log Alert")
    m.SetBody("text/plain", message)

    d := gomail.NewDialer(en.smtpServer, en.smtpPort, en.username, en.password)
    return d.DialAndSend(m)
}
