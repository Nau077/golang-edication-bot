package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Command) getGolangInfo(chatId int64, _ string) error {
	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Типы и структуры данных", "/getDataTypes"),
			tgbotapi.NewInlineKeyboardButtonData("Циклы", "/getLoopsInfoCommand"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Условные выражения", "/getConditionsInfo"),
			tgbotapi.NewInlineKeyboardButtonData("Работы с файлами", "/getFilesWorkInfo"),
		),

		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Функции", "/getFunctionsInfo"),
			tgbotapi.NewInlineKeyboardButtonData("Работа с json", "/getJsonWorkInfo"),
			tgbotapi.NewInlineKeyboardButtonData("Указатели", "/getPointersInfo"),
		),

		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Интерфейсы", "/getInterfacesInfo"),
			tgbotapi.NewInlineKeyboardButtonData("Горутины", "/getGoroutinsInfo"),
			tgbotapi.NewInlineKeyboardButtonData("Каналы", "/getChannelsInfo"),
		),
	)

	message := tgbotapi.NewMessage(chatId, "Теория go")
	message.ReplyMarkup = numericKeyboard

	err := c.producer.Send(&message)
	if err != nil {
		return err
	}
	return nil
}
