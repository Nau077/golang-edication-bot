package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Command) startCommand(chatId int64, text string) error {
	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Теория golang", "/getGolangInfo"),
			tgbotapi.NewInlineKeyboardButtonData("System-design party", "/getSytemDesignTopics"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Задачки", "/getTaskCommand"),
			tgbotapi.NewInlineKeyboardButtonData("Help", "/help"),
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
