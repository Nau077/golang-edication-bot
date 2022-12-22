package telegram

import (
	"fmt"
	"golang-edication-bot/internal/presentation/events/producer"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler func(msg *tgbotapi.Message) error

type Command struct {
	producer    *producer.TelegramProducer
	commandsMap map[string]Handler
}

func NewComand(producer *producer.TelegramProducer) *Command {
	m := map[string]Handler{
		"/start": newStartCommand().handle,
		"/help":  newHelpCommand().handle,
	}

	return &Command{
		producer:    producer,
		commandsMap: m,
	}
}

type startCommand struct {
	command Command
}

type helpCommand struct {
	command Command
}

func newStartCommand() *startCommand {
	return &startCommand{}
}

func newHelpCommand() *helpCommand {
	return &helpCommand{}
}

type CommandHandle interface {
	handle()
}

func (s *startCommand) handle(msg *tgbotapi.Message) error {
	fmt.Println("in start")
	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(

		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
			tgbotapi.NewInlineKeyboardButtonData("2", "2"),
			tgbotapi.NewInlineKeyboardButtonData("3", "3"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("4", "4"),
			tgbotapi.NewInlineKeyboardButtonData("5", "5"),
			tgbotapi.NewInlineKeyboardButtonData("6", "6"),
		),
	)

	message := tgbotapi.NewMessage(msg.Chat.ID, msg.Text)
	message.ReplyMarkup = numericKeyboard

	err := s.command.producer.Send(&message)
	if err != nil {
		return err
	}

	return nil
}

func (h *helpCommand) handle(msg *tgbotapi.Message) error {
	var err error
	if err != nil {
		return err
	}

	return nil
}
