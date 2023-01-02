package commands

import (
	"context"
	"golang-edication-bot/internal/infrustructure/repositories"
	"golang-edication-bot/internal/presentation/events/producer"
	"math"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler func(chatId int64, text string) error

type Command struct {
	producer *producer.TelegramProducer
	repos    *repositories.Container
	ctx      context.Context
}

type CommandArgs struct {
}

const (
	maxTgMsgLen = 4096
)

var numericKeyboardStart = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Загрузить стартовое меню", "/start"),
	),
)

func NewCommand(producer *producer.TelegramProducer, repos *repositories.Container, ctx context.Context) *Command {

	return &Command{
		producer: producer,
		repos:    repos,
		ctx:      ctx,
	}
}

func (c *Command) GetCommandsMap() map[string]Handler {
	m := map[string]Handler{
		"/start":                  c.startCommand,
		"/help":                   c.helpCommand,
		"/getSytemDesignTopics":   c.getSytemDesignTopics,
		"/getTaskCommand":         c.getTaskCommand,
		"/getGolangInfo":          c.getGolangInfo,
		"/getDataTypes":           c.getDataTypesCommand,
		"/getStringsInfo":         c.getStringsInfoCommand,
		"/getNumbersInfoCommand":  c.getNumbersInfoCommand,
		"/getMapsInfoCommand":     c.getMapsInfoCommand,
		"/geStructureInfoCommand": c.geStructureInfoCommand,
		"/getLoopsInfoCommand":    c.getLoopsInfoCommand,
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
