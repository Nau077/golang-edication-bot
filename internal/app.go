package internal

import (
	"context"
	"fmt"
	"golang-edication-bot/internal/infrustructure/providers/telegram"
	"golang-edication-bot/internal/presentation/events"
)

type App struct {
	provider   *telegram.Provider
	staticPath string
	bot        events.BotStarter
}

func NewApp(ctx context.Context, staticPath string) (*App, error) {
	a := &App{
		staticPath: staticPath,
	}
	err := a.initDeps(ctx)

	return a, err
}

func (a *App) Run() error {
	defer func() {
		a.provider.GetDB().Close()
	}()

	if a.bot != nil {
		if err := a.bot.Start(); err != nil {
			fmt.Printf("error at initBot: %s", err)
			return err
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

func (a *App) initProvider(ctx context.Context) error {
	a.provider = telegram.NewProvider(a.staticPath, ctx)

	return nil
}

func (a *App) initBot(ctx context.Context) error {
	a.bot = a.provider.GetBotStarter(ctx)

	return nil
}
