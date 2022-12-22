package client

import (
	"errors"
	"golang-edication-bot/internal/infrustructure/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Client interface {
	init() error
}

type TelegramClient struct {
	config *config.Config
	token  string
	Client *tgbotapi.BotAPI
}

func NewTelegramClient(config *config.Config) (*TelegramClient, error) {

	client, err := tgbotapi.NewBotAPI(config.TelegramBot.Token)
	client.Debug = true

	if err != nil {
		return nil, err
	}

	return &TelegramClient{
		token:  config.TelegramBot.Token,
		Client: client,
	}, nil
}

func (t *TelegramClient) GetTgClient() (*tgbotapi.BotAPI, error) {
	if t.Client == nil {
		return nil, errors.New("client field is empty ")
	}

	return t.Client, nil
}

func (t *TelegramClient) init() error {
	client, err := tgbotapi.NewBotAPI(t.token)

	if err != nil {
		return err
	}

	client.Debug = true

	t.Client = client

	return nil
}
