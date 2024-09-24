package notifier

import (
    "exception_log_collector/config"
)

type Notifier interface {
    Notify(message string) error
}

func NewNotifier(alertConfig config.AlertConfig) Notifier {
    switch alertConfig.Type {
    case "webhook":
        return NewWebhookNotifier(alertConfig.Config)
    case "email":
        return NewEmailNotifier(alertConfig.Config)
    case "kafka":
        return NewKafkaNotifier(alertConfig.Config)
    default:
        return nil
    }
}
