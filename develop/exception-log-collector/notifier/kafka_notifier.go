package notifier

import (
    "github.com/Shopify/sarama"
)

type KafkaNotifier struct {
    producer sarama.SyncProducer
    topic    string
}

func NewKafkaNotifier(config map[string]interface{}) *KafkaNotifier {
    brokers := config["brokers"].([]interface{})
    brokersStr := make([]string, len(brokers))
    for i, b := range brokers {
        brokersStr[i] = b.(string)
    }
    topic := config["topic"].(string)
    producer, _ := sarama.NewSyncProducer(brokersStr, nil)
    return &KafkaNotifier{
        producer: producer,
        topic:    topic,
    }
}

func (kn *KafkaNotifier) Notify(message string) error {
    msg := &sarama.ProducerMessage{
        Topic: kn.topic,
        Value: sarama.StringEncoder(message),
    }
    _, _, err := kn.producer.SendMessage(msg)
    return err
}
