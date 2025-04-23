package service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const apiUrl = "https://pro-api.coinmarketcap.com"

type converterService struct {
	token   string
	baseUrl *url.URL
	client  http.Client
}

type conversionResponse struct {
	Status responseStatus `json:"status"`
	Data   []struct {
		Quote map[string]struct {
			Price float64 `json:"price"`
		} `json:"quote"`
	} `json:"data"`
}

type responseStatus struct {
	ErrorMessage string `json:"error_message"`
}

func NewConverterService(token string) *converterService {
	url, _ := url.Parse(apiUrl)
	return &converterService{
		token:   token,
		baseUrl: url,
		client: http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func (s *converterService) Convert(ctx context.Context, from, to string, amount float64) (float64, error) {
	url := s.baseUrl.JoinPath("v2", "tools", "price-conversion")
	values := url.Query()
	values.Add("symbol", from)
	values.Add("convert", to)
	values.Add("amount", strconv.FormatFloat(amount, 'f', -1, 64))
	url.RawQuery = values.Encode()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	req.Header.Add("X-CMC_PRO_API_KEY", s.token)

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var conversionResp conversionResponse
	if err := json.NewDecoder(resp.Body).Decode(&conversionResp); err != nil {
		return 0, err
	}

	if resp.StatusCode != 200 {
		return 0, errors.New(conversionResp.Status.ErrorMessage)
	}

	if len(conversionResp.Data) == 0 {
		return 0, errors.New("no data")
	}

	if _, ok := conversionResp.Data[0].Quote[to]; !ok {
		return 0, errors.New("no quote")
	}

	return conversionResp.Data[0].Quote[to].Price, nil
}
