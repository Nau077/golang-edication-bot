package telegram

import (
	"context"
	"golang-edication-bot/internal/infrustructure/config"
	"golang-edication-bot/internal/presentation/client"
	"golang-edication-bot/internal/presentation/events"
	"golang-edication-bot/internal/presentation/events/consumer"
	"golang-edication-bot/internal/presentation/events/producer"
	"golang-edication-bot/internal/services/events/telegram"
	"log"
)

type Provider struct {
	// db         db.Client
	configPath string
	config     *config.Config
	commands   *telegram.Command
	processor  *telegram.TelegramProcessor
	client     *client.TelegramClient
	consumer   *consumer.TelegramConsumer
	producer   *producer.TelegramProducer
	botStarter *events.TelegramBotStarter
}

func NewProvider(configPath string) *Provider {
	return &Provider{
		configPath: configPath,
	}
}

func (p *Provider) GetConfig() *config.Config {
	if p.config == nil {
		cfg, err := config.NewConfig(p.configPath)
		if err != nil {
			log.Fatalf("failed to get config: %s", err.Error())
		}
		p.config = cfg
	}

	return p.config
}

func (p *Provider) GetClient() *client.TelegramClient {
	if p.client == nil {
		c, err := client.NewTelegramClient(p.GetConfig())

		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}
		p.client = c
	}

	return p.client
}

func (p *Provider) GetProducer() *producer.TelegramProducer {
	if p.producer == nil {
		pr := producer.NewTelegramProducer(p.GetClient())

		p.producer = pr
	}

	return p.producer
}

func (p *Provider) GetProcessor() *telegram.TelegramProcessor {
	if p.processor == nil {
		pr := telegram.NewTelegramProcessor(p.GetConfig(), p.GetProducer())

		p.processor = pr
	}

	return p.processor
}

func (p *Provider) GetConsumer() *consumer.TelegramConsumer {
	if p.consumer == nil {
		c := consumer.NewTelegramConsumer(p.GetProcessor(), p.GetClient())

		p.consumer = c
	}

	return p.consumer
}

func (p *Provider) GetBotStarter(ctx context.Context) *events.TelegramBotStarter {
	if p.botStarter == nil {
		b := events.NewTelegramBotStarter(p.GetConsumer())

		p.botStarter = b
	}

	return p.botStarter
}
