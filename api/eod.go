package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type EndOfDay struct {
	Data []struct {
		Open        float64     `json:"open"`
		High        float64     `json:"high"`
		Low         float64     `json:"low"`
		Close       float64     `json:"close"`
		Volume      float64     `json:"volume"`
		AdjHigh     interface{} `json:"adj_high"`
		AdjLow      interface{} `json:"adj_low"`
		AdjClose    float64     `json:"adj_close"`
		AdjOpen     interface{} `json:"adj_open"`
		AdjVolume   interface{} `json:"adj_volume"`
		SplitFactor float64     `json:"split_factor"`
		Dividend    float64     `json:"dividend"`
		Symbol      string      `json:"symbol"`
		Exchange    string      `json:"exchange"`
		Date        string      `json:"date"`
	} `json:"data"`
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func EndOfDayRequest(symbols string, date string) (EndOfDay, error) {
	var endpointURL string = "/eod/" + date

	baseURL := BaseURL()
	values := baseURL.Query()

	values.Add("symbols", symbols)

	baseURL.RawQuery = values.Encode()
	baseURL.Path = baseURL.Path + endpointURL

	response, err := http.Get(baseURL.String())

	if err != nil {
		return EndOfDay{}, errors.New("the HTTP request has failed with an error")
	}

	data, _ := ioutil.ReadAll(response.Body)
	endOfDay := EndOfDay{}
	err = json.Unmarshal(data, &endOfDay)
	if err != nil {
		return EndOfDay{}, err
	}

	if endOfDay.Error.Code != "" {
		return EndOfDay{}, errors.New(endOfDay.Error.Message)
	}

	return endOfDay, nil
}
