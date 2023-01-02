package commands

import (
	"fmt"
	"golang-edication-bot/internal/services/events/telegram/models"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Command) getSytemDesignTopics(chatId int64, _ string) error {
	goInfoList, err := c.repos.GoInfoRepo.GetDataListType(c.ctx, "system_design")
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
	var stRecords []models.SystemDesignMsg

	for _, record := range goInfoList {
		stRecords = append(stRecords, models.SystemDesignMsg{
			Title: record.Title,
			Text:  record.Text,
		})
	}
	var text string

	for i := 0; i < len(stRecords); i++ {
		text = text + strconv.Itoa(i) + ") " +
			stRecords[i].Title + " " + "<b>" +
			stRecords[i].Text + "</b>" + "\n"
	}

	message := tgbotapi.NewMessage(chatId, text)
	message.ParseMode = tgbotapi.ModeHTML
	message.ReplyMarkup = numericKeyboardStart
	err = c.producer.Send(&message)
	if err != nil {
		return err
	}

	return nil
}
