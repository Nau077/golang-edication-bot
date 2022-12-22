package events

import (
	"golang-edication-bot/internal/presentation/events/consumer"
)

type TelegramBotStarter struct {
	consumer *consumer.TelegramConsumer
}

func NewTelegramBotStarter(consumer *consumer.TelegramConsumer) *TelegramBotStarter {
	return &TelegramBotStarter{
		consumer: consumer,
	}
}

func (t *TelegramBotStarter) Start() error {
	if err := t.consumer.Fetch(); err != nil {
		return err
	}
	return nil
}
