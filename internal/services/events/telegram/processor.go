package telegram

import (
	"fmt"
	"golang-edication-bot/internal/infrustructure/config"
	"golang-edication-bot/internal/infrustructure/repositories"
	"golang-edication-bot/internal/presentation/events/producer"
	"strings"

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

func (t *TelegramProcessor) Process(msg *tgbotapi.Message) error {
	text := strings.ToLower(msg.Text)

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

	err := handler(msg)

	if err != nil {
		return err
	}

	return nil
}
