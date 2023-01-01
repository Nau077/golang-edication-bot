package telegram

import (
	"context"
	"golang-edication-bot/internal/infrustructure/config"
	"golang-edication-bot/internal/infrustructure/repositories"
	"golang-edication-bot/internal/presentation/events/producer"
	"regexp"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ProcessProcessor interface {
	Process(msg *tgbotapi.Message)
}

type TelegramProcessor struct {
	config   *config.Config
	Producer *producer.TelegramProducer
	repos    *repositories.Container
	ctx      context.Context
}

type ProcArgs struct {
	Config   *config.Config
	Producer *producer.TelegramProducer
	Repos    *repositories.Container
	Ctx      context.Context
}

func NewTelegramProcessor(procArgs *ProcArgs) *TelegramProcessor {
	return &TelegramProcessor{
		config:   procArgs.Config,
		Producer: procArgs.Producer,
		repos:    procArgs.Repos,
		ctx:      procArgs.Ctx,
	}
}

func (t *TelegramProcessor) Process(update *tgbotapi.Update) error {
	var text string
	var chatId int64

	if update.Message == nil && update.CallbackQuery != nil {
		chatId = update.CallbackQuery.Message.Chat.ID
		text = update.CallbackQuery.Data
	}

	if update.Message != nil {
		chatId = update.Message.Chat.ID
		text = update.Message.Text
	}

	var command = NewCommand(t.Producer, t.repos, t.ctx)

	if strings.Contains(text, "{") {
		re := regexp.MustCompile("[0-9]+")
		result := re.FindAllString(text, -1)
		i, err := strconv.ParseInt(result[0], 10, 64)
		if err != nil {
			return err
		}

		command.GetTaskSolution(chatId, i)
		return nil
	}

	if command.GetCommandsMap()[text] == nil {
		command.SendNoCommandsMsg(chatId)
		return nil
	}

	handler, ok := command.GetCommandsMap()[text]

	if !ok {
		command.SendNoCommandsMsg(chatId)
		return nil
	}

	err := handler(chatId, text)

	if err != nil {
		return err
	}

	return nil
}
