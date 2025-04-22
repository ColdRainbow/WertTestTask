package convert

import (
	"context"
	"errors"
	"fmt"
	"strconv"
)

type ConverterService interface {
	Convert(ctx context.Context, from, to string, amount float64) (float64, error)
}

type convertUsecase struct {
	service ConverterService
}

func New(service ConverterService) *convertUsecase {
	return &convertUsecase{
		service: service,
	}
}

func (u *convertUsecase) Execute(ctx context.Context, args []string) error {
	if len(args) != 3 {
		return errors.New("wrong number of arguments")
	}

	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return fmt.Errorf("cannot parse amount: %w", err)
	}

	converted, err := u.service.Convert(ctx, args[1], args[2], amount)
	if err != nil {
		return fmt.Errorf("conversion service error: %w", err)
	}

	fmt.Println(converted)

	return nil
}
