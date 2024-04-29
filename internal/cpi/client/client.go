package client

import (
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/clientcredentials"
)

func NewClient() *resty.Client {
	config := &clientcredentials.Config{
		TokenURL:     viper.GetString("tenant.token_url"),
		ClientID:     viper.GetString("tenant.client_id"),
		ClientSecret: viper.GetString("tenant.client_secret"),
	}

	httpClient := config.Client(context.Background())
	restyClient := resty.NewWithClient(httpClient).
		SetBaseURL(viper.GetString("tenant.base_url"))

	return restyClient
}
