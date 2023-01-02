package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Command) getDataTypesCommand(chatId int64, _ string) error {
	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(

		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Строки", "/getStringsInfo"),
			tgbotapi.NewInlineKeyboardButtonData("Числа", "/getNumbersInfoCommand"),
			tgbotapi.NewInlineKeyboardButtonData("Map", "/getMapsInfoCommand"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Структуры", "/geStructureInfoCommand"),
		),
	)

	message := tgbotapi.NewMessage(chatId, "Типы и структуры данных")
	message.ReplyMarkup = numericKeyboard

	err := c.producer.Send(&message)
	if err != nil {
		return err
	}

	return nil
}
