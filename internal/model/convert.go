package model

type ConversionArgs struct {
	Amount float64
	From   string
	To     string
}

type ConversionResult struct {
	Amount float64
}
