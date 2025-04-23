package convert

import (
	"context"
	"converter/internal/model"
	"fmt"
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

func (u *convertUsecase) Execute(ctx context.Context, args *model.ConversionArgs) (*model.ConversionResult, error) {
	converted, err := u.service.Convert(ctx, args.From, args.To, args.Amount)
	if err != nil {
		return nil, fmt.Errorf("conversion service error: %w", err)
	}

	return &model.ConversionResult{
		Amount: converted,
	}, nil
}
