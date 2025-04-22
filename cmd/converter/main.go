package main

import (
	"context"
	"converter/internal/service"
	"converter/internal/usecase/convert"
	"fmt"
	"os"
	"time"
)

func main() {
	token := os.Getenv("COINMARKET_API_TOKEN")
	if token == "" {
		fmt.Println("missing CoinMarket API token in COINMARKET_API_TOKEN")
		os.Exit(1)
	}

	if len(os.Args) != 4 {
		fmt.Printf("usage: %s <amount> <from> <to>\n", os.Args[0])
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	service := service.NewConverterService(token)
	usecase := convert.New(service)
	err := usecase.Execute(ctx, os.Args[1:])
	if err != nil {
		fmt.Printf("execution error: %v\n", err)
		os.Exit(1)
	}
}
