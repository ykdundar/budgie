package api

import (
	"fmt"
	"github.com/spf13/viper"
	"net/url"
)

func BaseURL() *url.URL {
	return &url.URL{
		Scheme:     "http",
		Host:       "api.marketstack.com",
		Path:       "v1",
		ForceQuery: false,
		RawQuery:   fmt.Sprintf("access_key=%s", viper.GetString("token")),
	}
}
