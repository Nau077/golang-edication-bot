package commands

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Command) getTaskCommand(chatId int64, _ string) error {

	task, err := c.repos.GoTasksPgRepo.GetTask(c.ctx)
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

	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Ответ на задачу", "/TasksSolution:{"+strconv.FormatInt(task.Id, 10)+"}"),
		),
	)

	if len(task.TasksText) > maxTgMsgLen {

		subs := SplitSubN(task.TasksText, maxTgMsgLen)

		for i := 0; i < len(subs); i++ {
			message := tgbotapi.NewMessage(chatId, subs[i])
			message.ParseMode = tgbotapi.ModeHTML
			message.ReplyMarkup = numericKeyboard
			err = c.producer.Send(&message)
			if err != nil {
				return err
			}
		}

		return nil
	}

	message := tgbotapi.NewMessage(chatId, task.TasksText)
	message.ReplyMarkup = numericKeyboard
	message.ParseMode = tgbotapi.ModeHTML
	err = c.producer.Send(&message)
	if err != nil {
		return err
	}

	return nil
}

func (c *Command) GetTaskSolution(chatId int64, id int64) error {
	solution, err := c.repos.GoTasksPgRepo.GetTasksSolution(c.ctx, id)
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

	message := tgbotapi.NewMessage(chatId, solution)
	message.ParseMode = tgbotapi.ModeHTML
	message.ReplyMarkup = numericKeyboardStart
	err = c.producer.Send(&message)
	if err != nil {
		return err
	}

	return nil
}
