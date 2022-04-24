package functions

import "github.com/ykdundar/budgie/api"

func GetLatestPrice(ticker string) float64 {
	currentPrice, _ := api.IntradayRequest([]string{ticker})
	lastPriceFromIntraday := currentPrice.Data[0].Last

	if lastPriceFromIntraday == 0 {
		lastPriceFromEod, _ := api.EndOfDayRequest([]string{ticker}, "latest")
		return lastPriceFromEod.Data[0].Close
	}

	return lastPriceFromIntraday
}
