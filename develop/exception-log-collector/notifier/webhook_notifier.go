package notifier

import (
    "bytes"
    "encoding/json"
    "net/http"
)

type WebhookNotifier struct {
    url string
}

func NewWebhookNotifier(config map[string]interface{}) *WebhookNotifier {
    return &WebhookNotifier{
        url: config["url"].(string),
    }
}

func (wn *WebhookNotifier) Notify(message string) error {
    payload := map[string]string{"text": message}
    data, _ := json.Marshal(payload)
    _, err := http.Post(wn.url, "application/json", bytes.NewBuffer(data))
    return err
}
