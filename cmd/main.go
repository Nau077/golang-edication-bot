package main

import (
	"context"
	"fmt"
	"golang-edication-bot/internal"
	"log"
	"os"
)

func main() {
	staticPath := os.Args[1]
	ctx := context.Background()
	fmt.Println(staticPath)
	a, err := internal.NewApp(ctx, staticPath)

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app %s", err.Error())
	}
}
