package main

import (
	"context"
	"flag"
	"golang-edication-bot/internal"
	"log"
)

var pathConfig string

func init() {
	flag.StringVar(&pathConfig, "config", "config/config.json", "Path to configuration file")
}

func main() {
	flag.Parse()
	ctx := context.Background()
	a, err := internal.NewApp(ctx, "")

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app %s", err.Error())
	}
}
