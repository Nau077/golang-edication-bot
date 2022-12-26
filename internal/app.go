package internal

import (
	"context"
	"fmt"
	"golang-edication-bot/internal/infrustructure/providers/telegram"
	"golang-edication-bot/internal/presentation/events"
)

type App struct {
	provider   *telegram.Provider
	pathConfig string
	bot        events.BotStarter
}

func NewApp(ctx context.Context, pathConfig string) (*App, error) {
	a := &App{
		pathConfig: pathConfig,
	}
	err := a.initDeps(ctx)

	return a, err
}

func (a *App) Run() error {
	// defer func() {
	// 	a.provider.db.Close()
	// }()
	if a.bot != nil {
		if err := a.bot.Start(); err != nil {
			fmt.Printf("error at initBot: %s", err)
		}
	}

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initProvider,
		a.initBot,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initProvider(_ context.Context) error {
	const configPath = "/home/roman/web/golang-projects/go-userfind-bot/config/config.json"
	a.provider = telegram.NewProvider(configPath)

	return nil
}

func (a *App) initBot(ctx context.Context) error {
	a.bot = a.provider.GetBotStarter(ctx)

	return nil
}
