package telegram

import (
	"context"
	"golang-edication-bot/internal/infrustructure/config"
	"golang-edication-bot/internal/infrustructure/libs/db"
	"golang-edication-bot/internal/infrustructure/repositories"
	"golang-edication-bot/internal/presentation/client"
	"golang-edication-bot/internal/presentation/events"
	"golang-edication-bot/internal/presentation/events/consumer"
	"golang-edication-bot/internal/presentation/events/producer"
	"golang-edication-bot/internal/services/events/telegram"
	"log"
)

type Provider struct {
	staticPath string
	config     *config.Config
	commands   *telegram.Command
	processor  *telegram.TelegramProcessor
	client     *client.TelegramClient
	consumer   *consumer.TelegramConsumer
	producer   *producer.TelegramProducer
	botStarter *events.TelegramBotStarter
	goInfoRepo repositories.InfoData
	db         db.Client
	ctx        context.Context
}

func NewProvider(staticPath string, ctx context.Context) *Provider {
	return &Provider{
		staticPath: staticPath,
		ctx:        ctx,
	}
}

func (p *Provider) GetConfig() *config.Config {
	if p.config == nil {
		configPath := p.staticPath + "/config/config.json"
		cfg, err := config.NewConfig(configPath)
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

func (p *Provider) GetDB() db.Client {
	if p.db == nil {
		cfg, err := p.GetConfig().GetDBConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}

		dbc, err := db.NewClient(p.ctx, cfg)
		if err != nil {
			log.Fatalf("cant connect to db err %s", err.Error())
		}
		p.db = dbc
	}

	return p.db
}

func (p Provider) GetGoInfoRepo() repositories.InfoData {
	if p.goInfoRepo == nil {
		g := repositories.NewGoInfoPgRepo(p.staticPath, p.GetDB())

		p.goInfoRepo = g
	}

	return p.goInfoRepo
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
		pr := telegram.NewTelegramProcessor(
			&telegram.ProcArgs{
				Config:   p.GetConfig(),
				Producer: p.GetProducer(),
				Repo:     p.GetGoInfoRepo(),
				Ctx:      p.ctx,
			},
		)

		p.processor = pr
	}

	return p.processor
}

func (p *Provider) GetConsumer() *consumer.TelegramConsumer {
	if p.consumer == nil {
		c := consumer.NewTelegramConsumer(p.GetProcessor())

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
