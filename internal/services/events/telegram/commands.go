package telegram

import (
	"context"
	"fmt"
	"golang-edication-bot/internal/infrustructure/repositories"
	"golang-edication-bot/internal/presentation/events/producer"
	"math"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler func(chatId int64, text string) error

type Command struct {
	producer   *producer.TelegramProducer
	goInfoRepo repositories.InfoData
	ctx        context.Context
}

func NewCommand(producer *producer.TelegramProducer, goInfoRepo repositories.InfoData, ctx context.Context) *Command {

	return &Command{
		producer:   producer,
		goInfoRepo: goInfoRepo,
		ctx:        ctx,
	}
}

func (c *Command) GetCommandsMap() map[string]Handler {
	m := map[string]Handler{
		"/start":                 c.startCommand,
		"/help":                  c.helpCommand,
		"/getDataTypes":          c.getDataTypesCommand,
		"/getStringsInfo":        c.getStringsInfoCommand,
		"/getNumbersInfoCommand": c.getNumbersInfoCommand,
		"/getMapsInfoCommand":    c.getMapsInfoCommand,
		"/getGolangInfo":         c.getGolangInfo,
	}

	return m
}

func (c *Command) SendNoCommandsMsg(chatId int64) error {
	message := tgbotapi.NewMessage(chatId, "Такой команды не существует")

	err := c.producer.Send(&message)
	if err != nil {
		return err
	}
	return nil
}

func (c *Command) startCommand(chatId int64, text string) error {
	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Теория golang", "/getGolangInfo"),
			tgbotapi.NewInlineKeyboardButtonData("System-design party", "/getDataTypes"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Задачки", "/getTasks"),
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

func (c *Command) getGolangInfo(chatId int64, text string) error {
	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Типы и структуры данных", "/getDataTypes"),
			tgbotapi.NewInlineKeyboardButtonData("Циклы", "/getLoopsInfo"),
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
			tgbotapi.NewInlineKeyboardButtonData("Map", "/getMapsInfoCommand"),
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
	data, err := c.goInfoRepo.GetData(c.ctx, "strings")
	if err != nil {
		return err
	}

	message := tgbotapi.NewMessage(chatId, data.Text)
	message.ParseMode = tgbotapi.ModeHTML
	err = c.producer.Send(&message)
	if err != nil {
		return err
	}

	return nil
}

func (c *Command) getNumbersInfoCommand(chatId int64, _ string) error {
	goInfo, err := c.goInfoRepo.GetData(c.ctx, "numbers")
	if err != nil {
		return err
	}

	message := tgbotapi.NewMessage(chatId, goInfo.Text)
	message.ParseMode = tgbotapi.ModeHTML
	err = c.producer.Send(&message)
	if err != nil {
		return err
	}

	return nil
}

func (c *Command) getMapsInfoCommand(chatId int64, _ string) error {
	fmt.Println("in getMapsInfoCommand")

	goInfo, err := c.goInfoRepo.GetData(c.ctx, "maps")
	if err != nil {
		return err
	}

	maxTgMsgLen := 4096

	if len(goInfo.Text) > maxTgMsgLen {

		subs := SplitSubN(goInfo.Text, maxTgMsgLen)

		for i := 0; i < len(subs); i++ {
			message := tgbotapi.NewMessage(chatId, subs[i])
			message.ParseMode = tgbotapi.ModeHTML
			err = c.producer.Send(&message)
			if err != nil {
				return err
			}
		}

		return nil
	}

	message := tgbotapi.NewMessage(chatId, goInfo.Text)
	message.ParseMode = tgbotapi.ModeHTML
	err = c.producer.Send(&message)
	if err != nil {
		return err
	}

	return nil
}

func SplitSubN(str string, size int) []string {
	strLength := len(str)
	splitedLength := int(math.Ceil(float64(strLength) / float64(size)))
	splited := make([]string, splitedLength)
	var start, stop int
	for i := 0; i < splitedLength; i += 1 {
		start = i * size
		stop = start + size
		if stop > strLength {
			stop = strLength
		}
		splited[i] = str[start:stop]
	}
	return splited
}
