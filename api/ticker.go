package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Ticker struct {
	Data []struct {
		Name          string      `json:"name"`
		Symbol        string      `json:"symbol"`
		HasIntraday   bool        `json:"has_intraday"`
		HasEod        bool        `json:"has_eod"`
		Country       interface{} `json:"country"`
		StockExchange struct {
			Name        string `json:"name"`
			Acronym     string `json:"acronym"`
			Mic         string `json:"mic"`
			Country     string `json:"country"`
			CountryCode string `json:"country_code"`
			City        string `json:"city"`
			Website     string `json:"website"`
		} `json:"stock_exchange"`
	} `json:"data"`
}

func TickerRequest(name string) (Ticker, error) {
	const endpointURL string = "/tickers"
	baseURL := BaseURL()
	values := baseURL.Query()

	values.Add("search", name)
	baseURL.RawQuery = values.Encode()
	baseURL.Path = baseURL.Path + endpointURL

	response, err := http.Get(baseURL.String())

	if err != nil {
		return Ticker{}, errors.New("the HTTP request has failed with an error")
	}

	data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	ticker := Ticker{}
	err = json.Unmarshal(data, &ticker)
	if err != nil {
		return Ticker{}, err
	}
	return ticker, nil
}
