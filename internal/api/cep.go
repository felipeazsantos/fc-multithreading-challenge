package api

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/felipeazsantos/pos-goexpert/fc-multithread-challenge/configs"
)

const (
	viacepURL    = "https://brasilapi.com.br/api/cep/v1/%s"
	brasilApiURL = "https://brasilapi.com.br/api/cep/v1/%s"
)

type cepAPI struct {
	cfg    configs.IConfig
	client *http.Client
}

func NewCepAPIS(cfg configs.IConfig) *cepAPI {
	return &cepAPI{
		cfg:    cfg,
		client: http.DefaultClient,
	}
}

func (a *cepAPI) GetAddressFromBrasilCepAPI(ctx context.Context, cep string, responseSuccess chan bool, responseErr chan error) {
	url := fmt.Sprintf(brasilApiURL, cep)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		a.cfg.LErr().Printf("error when trying create request with context: %v", err)
		responseErr <- err
		return
	}

	response, err := a.client.Do(req)
	if err != nil {
		a.cfg.LErr().Printf("error call url %s : %v", url, err)
		responseErr <- err
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		a.cfg.LErr().Printf("error parse body from url %s : %v", url, err)
		responseErr <- err
		return
	}
}

func (a *cepAPI) GetAddressFromViaCepAPI(ctx context.Context, cep string, responseSuccess chan bool, responseErr chan error) {
	url := fmt.Sprintf(viacepURL, cep)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		a.cfg.LErr().Printf("error when trying create request with context: %v", err)
	}

	response, err := a.client.Do(req)
	if err != nil {
		a.cfg.LErr().Printf("error call url %s : %v", url, err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		a.cfg.LErr().Printf("error parse body from url %s : %v", url, err)
	}
}
