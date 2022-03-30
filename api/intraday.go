package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type Intraday struct {
	Data []struct {
		Open     float64 `json:"open"`
		High     float64 `json:"high"`
		Low      float64 `json:"low"`
		Last     float64 `json:"last"`
		Close    float64 `json:"close"`
		Volume   float64 `json:"volume"`
		Date     string  `json:"date"`
		Symbol   string  `json:"symbol"`
		Exchange string  `json:"exchange"`
	} `json:"data"`
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}


func IntradayRequest(symbols []string) (Intraday, error) {
	const endpointURL string = "/intraday/latest"

	baseURL := BaseURL()
	values := baseURL.Query()

	values.Add("symbols", strings.Join(symbols, ","))

	baseURL.RawQuery = values.Encode()
	baseURL.Path = baseURL.Path + endpointURL

	response, err := http.Get(baseURL.String())

	if err != nil {
		return Intraday{}, errors.New("the HTTP request has failed with an error")
	}

	data, _ := ioutil.ReadAll(response.Body)
	intraday := Intraday{}
	err = json.Unmarshal(data, &intraday)
	if err != nil {
		return Intraday{}, err
	}

	if intraday.Error.Code != "" {
		return Intraday{}, errors.New(intraday.Error.Message)
	}

	return intraday, nil
}
