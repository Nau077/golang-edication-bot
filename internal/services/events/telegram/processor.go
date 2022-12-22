package telegram

import (
	"fmt"
	"golang-edication-bot/internal/infrustructure/config"
	"golang-edication-bot/internal/presentation/events/producer"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ProcessProcessor interface {
	Process(msg *tgbotapi.Message)
}

type TelegramProcessor struct {
	config   *config.Config
	producer *producer.TelegramProducer
}

func NewTelegramProcessor(config *config.Config, producer *producer.TelegramProducer) *TelegramProcessor {
	return &TelegramProcessor{
		config:   config,
		producer: producer,
	}
}

func (t *TelegramProcessor) Process(msg *tgbotapi.Message) error {
	text := strings.ToLower(msg.Text)

	var command = NewComand(t.producer)

	if command.commandsMap[text] == nil {
		fmt.Println("no command")
		return nil
	}

	err := command.commandsMap[text](msg)
	if err != nil {
		return err
	}

	return nil
}
