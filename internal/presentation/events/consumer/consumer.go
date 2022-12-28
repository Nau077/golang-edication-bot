package consumer

import (
	"fmt"
	"golang-edication-bot/internal/services/events/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Consumer interface {
	Fetch()
}

type TelegramConsumer struct {
	processor *telegram.TelegramProcessor
}

func NewTelegramConsumer(processor *telegram.TelegramProcessor) *TelegramConsumer {

	return &TelegramConsumer{
		processor: processor,
	}
}

func (t *TelegramConsumer) Fetch() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := t.processor.Producer.TelegramClient.Client.GetUpdatesChan(u)
	fmt.Println("start fetch updates")

	for update := range updates {

		// if !update.Message.IsCommand() {
		// 	continue
		// }

		err := t.processor.Process(&update)

		if err != nil {
			return err
		}

	}

	return nil
}
