package cli

import (
	"context"
	"converter/internal/model"
	"converter/internal/service"
	"converter/internal/usecase/convert"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Run() error {
	token := os.Getenv("COINMARKET_API_TOKEN")
	if token == "" {
		return errors.New("missing CoinMarket API token")
	}

	args, err := parseArgs(os.Args[1:])
	if err != nil {
		return fmt.Errorf("error processing arguments: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	service := service.NewConverterService(token)
	usecase := convert.New(service)
	result, err := usecase.Execute(ctx, args)
	if err != nil {
		return fmt.Errorf("execution error: %w", err)
	}

	printResult(result)

	return nil
}

func parseArgs(args []string) (*model.ConversionArgs, error) {
	if len(args) != 3 {
		return nil, errors.New("wrong number of arguments")
	}

	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return nil, fmt.Errorf("cannot parse amount: %w", err)
	}

	return &model.ConversionArgs{
		Amount: amount,
		From:   strings.ToUpper(args[1]),
		To:     strings.ToUpper(args[2]),
	}, nil
}

func printResult(result *model.ConversionResult) {
	fmt.Println(result.Amount)
}
