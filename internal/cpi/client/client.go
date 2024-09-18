package client

import (
	"github.com/go-resty/resty/v2"
	"github.com/vadimklimov/cpi-navigator/internal/config"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/clientcredentials"
)

func NewClient() *resty.Client {
	oauthConfig := &clientcredentials.Config{
		TokenURL:     config.TenantTokenURL().String(),
		ClientID:     config.TenantClientID(),
		ClientSecret: config.TenantClientSecret(),
	}

	httpClient := oauthConfig.Client(context.Background())
	restyClient := resty.NewWithClient(httpClient).
		SetBaseURL(config.TenantBaseURL().String())

	return restyClient
}
