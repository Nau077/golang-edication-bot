package telegram

import (
	"golang-edication-bot/internal/infrustructure/repositories"
	"golang-edication-bot/internal/presentation/events/producer"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler func(chatId int64, text string) error

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
		"/start":                 c.startCommand,
		"/help":                  c.helpCommand,
		"/getDataTypes":          c.getDataTypesCommand,
		"/getStringsInfo":        c.getStringsInfoCommand,
		"/getNumbersInfoCommand": c.getNumbersInfoCommand,
	}

	return m
}

func (c *Command) startCommand(chatId int64, text string) error {
	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(

		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Типы данных", "/getDataTypes"),
			tgbotapi.NewInlineKeyboardButtonData("3", "3"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("4", "4"),
			tgbotapi.NewInlineKeyboardButtonData("5", "5"),
			tgbotapi.NewInlineKeyboardButtonData("6", "6"),
		),
	)

	message := tgbotapi.NewMessage(chatId, text)
	message.ReplyMarkup = numericKeyboard

	err := c.producer.Send(&message)
	if err != nil {
		return err
	}

	return nil
}

func (c *Command) getDataTypesCommand(chatId int64, text string) error {
	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(

		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Строки", "/getStringsInfo"),
			tgbotapi.NewInlineKeyboardButtonData("Числа", "/getNumbersInfoCommand"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("4", "4"),
			tgbotapi.NewInlineKeyboardButtonData("5", "5"),
			tgbotapi.NewInlineKeyboardButtonData("6", "6"),
		),
	)

	message := tgbotapi.NewMessage(chatId, text)
	message.ReplyMarkup = numericKeyboard

	err := c.producer.Send(&message)
	if err != nil {
		return err
	}

	return nil
}

func (c *Command) helpCommand(chatId int64, text string) error {
	var err error
	if err != nil {
		return err
	}

	return nil
}

func (c *Command) getStringsInfoCommand(chatId int64, _ string) error {
	// stringsData, err := c.goInfoRepo.GetData("strings.json")
	// if err != nil {
	// 	return err
	// }

	// message := tgbotapi.NewMessage(chatId, string(stringsData))
	// message.ParseMode = tgbotapi.ModeMarkdown
	// err = c.producer.Send(&message)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (c *Command) getNumbersInfoCommand(chatId int64, _ string) error {
	stringsData, err := c.goInfoRepo.GetData("numbers.txt")
	if err != nil {
		return err
	}

	message := tgbotapi.NewMessage(chatId, string(stringsData))
	message.ParseMode = tgbotapi.ModeHTML
	err = c.producer.Send(&message)
	if err != nil {
		return err
	}

	return nil
}
