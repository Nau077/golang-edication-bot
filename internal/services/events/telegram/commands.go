package telegram

import (
	"fmt"
	"golang-edication-bot/internal/infrustructure/repositories"
	"golang-edication-bot/internal/presentation/events/producer"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler func(msg *tgbotapi.Message) error

type Command struct {
	producer   *producer.TelegramProducer
	goInfoRepo repositories.InfoData
}

func NewCommand(producer *producer.TelegramProducer, goInfoRepo repositories.InfoData) *Command {

	return &Command{
		producer:   producer,
		goInfoRepo: goInfoRepo,
	}
}

func (c *Command) GetCommandsMap() map[string]Handler {
	m := map[string]Handler{
		"/start":        c.startCommand,
		"/help":         c.helpCommand,
		"/getDataTypes": c.getDataTypesCommand,
	}

	return m
}

func (c *Command) startCommand(msg *tgbotapi.Message) error {
	fmt.Println("in start")
	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(

		tgbotapi.NewInlineKeyboardRow(
			//	tgbotapi.NewInlineKeyboardButtonURL("Типы данных", "/start"),
			tgbotapi.NewInlineKeyboardButtonData("Типы данных", "/start"),
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

	err := c.producer.Send(&message)
	if err != nil {
		return err
	}

	return nil
}

func (c *Command) getDataTypesCommand(msg *tgbotapi.Message) error {
	fmt.Println("in getDataTypesCommand")

	data, err := c.goInfoRepo.GetData(msg.Text)
	if err != nil {
		return err
	}

	message := tgbotapi.NewMessage(msg.Chat.ID, data.Data)

	err = c.producer.Send(&message)
	if err != nil {
		return err
	}

	return nil

}

func (c *Command) helpCommand(msg *tgbotapi.Message) error {
	var err error
	if err != nil {
		return err
	}

	return nil
}
