package telegram

import (
	"fmt"
	"golang-edication-bot/internal/infrustructure/config"
	"golang-edication-bot/internal/infrustructure/repositories"
	"golang-edication-bot/internal/presentation/events/producer"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ProcessProcessor interface {
	Process(msg *tgbotapi.Message)
}

type TelegramProcessor struct {
	config   *config.Config
	Producer *producer.TelegramProducer
	repo     repositories.InfoData
}

type ProcArgs struct {
	Config   *config.Config
	Producer *producer.TelegramProducer
	Repo     repositories.InfoData
}

func NewTelegramProcessor(procArgs *ProcArgs) *TelegramProcessor {
	return &TelegramProcessor{
		config:   procArgs.Config,
		Producer: procArgs.Producer,
		repo:     procArgs.Repo,
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

	var command = NewCommand(t.Producer, t.repo)

	if command.GetCommandsMap()[text] == nil {
		fmt.Println("no command")
		return nil
	}

	handler, ok := command.GetCommandsMap()[text]

	if !ok {
		fmt.Println("no command")
		return nil
	}

	err := handler(chatId, text)

	if err != nil {
		return err
	}

	return nil
}
