package producer

import (
	"golang-edication-bot/internal/presentation/client"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Producer interface {
	send(message *tgbotapi.Message) error
}

type Type int

const (
	Unknown Type = iota
	Message
)

type Event struct {
	Type Type
	Text string
	Meta interface{}
}

type TelegramProducer struct {
	TelegramClient *client.TelegramClient
}

func NewTelegramProducer(telegramClient *client.TelegramClient) *TelegramProducer {

	return &TelegramProducer{
		TelegramClient: telegramClient,
	}
}

func (t *TelegramProducer) Send(msg *tgbotapi.MessageConfig) error {

	if _, err := t.TelegramClient.Client.Send(msg); err != nil {
		return err
	}

	return nil
}
