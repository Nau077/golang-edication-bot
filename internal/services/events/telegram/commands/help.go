package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Command) helpCommand(chatId int64, text string) error {
	message := tgbotapi.NewMessage(chatId, `
		Спасибо за использование golang-edication-bot.

		1) Команда /start активирует главное меню
		2) В меню "задачки" генерируется рандомная задача, кнопка с ответом прилагается
		3_Предложить задачи и контент для тем можете в issues здесь https://github.com/Nau077/golang-edication-bot
		Про баги писать там же

		4)Ставьте звёзды репозиторию и добавляйтесь сюда https://www.linkedin.com/in/roman-tolokontsev-b2b504179/
		Успехов в учёбе! 😌
	`)
	message.ReplyMarkup = numericKeyboardStart
	err := c.producer.Send(&message)

	if err != nil {
		return err
	}

	return nil
}
