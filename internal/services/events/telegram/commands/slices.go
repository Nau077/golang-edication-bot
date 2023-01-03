package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Command) getSlicesInfo(chatId int64, _ string) error {
	goInfo, err := c.repos.GoInfoRepo.GetData(c.ctx, "slices")
	if err != nil {
		if fmt.Sprint(err) == "scanning one: no rows in result set" {
			message := tgbotapi.NewMessage(chatId, "Нет данных")
			message.ReplyMarkup = numericKeyboardStart
			err = c.producer.Send(&message)
			if err != nil {
				return err
			}

			return nil
		}
		return err
	}

	if len(goInfo.Text) > maxTgMsgLen {

		subs := SplitSubN(goInfo.Text, maxTgMsgLen)

		for i := 0; i < len(subs); i++ {
			message := tgbotapi.NewMessage(chatId, subs[i])
			message.ParseMode = tgbotapi.ModeHTML
			if len(subs)-1 == i {
				message.ReplyMarkup = numericKeyboardStart
			}
			err = c.producer.Send(&message)
			if err != nil {
				return err
			}
		}

		return nil
	}

	message := tgbotapi.NewMessage(chatId, goInfo.Text)
	message.ParseMode = tgbotapi.ModeHTML
	message.ReplyMarkup = numericKeyboardStart
	err = c.producer.Send(&message)
	if err != nil {
		return err
	}

	return nil
}
