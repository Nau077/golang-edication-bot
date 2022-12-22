package internal

import (
	"context"
	"fmt"
	"golang-edication-bot/internal/infrustructure/providers/telegram"
)

type App struct {
	provider   *telegram.Provider
	pathConfig string
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
	bot := a.provider.GetBotStarter(ctx)
	var err error
	if bot.Start(); err != nil {
		fmt.Printf("error at initBot: %s", err)
	}

	return nil
}
