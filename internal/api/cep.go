package api

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/felipeazsantos/pos-goexpert/fc-multithread-challenge/configs"
)

const (
	viacepURL    = "https://viacep.com.br/ws/%s/json/"
	brasilApiURL = "https://brasilapi.com.br/api/cep/v1/%s"
)

var apisUrl = []string{
	brasilApiURL,
	viacepURL,
}

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

func (a *cepAPI) GetAddressByCep(ctx context.Context, cep string, responseSuccess chan map[string]any, responseErr chan error) {
	for _, u := range apisUrl {
		go func(apiURL string) {
			url := fmt.Sprintf(apiURL, cep)
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

			if response.StatusCode != http.StatusOK {
				a.cfg.LErr().Printf("error status code %d from url %s", response.StatusCode, url)
				responseErr <- fmt.Errorf("error status code %d from url %s", response.StatusCode, url)
				return
			}

			body, err := io.ReadAll(response.Body)
			if err != nil {
				a.cfg.LErr().Printf("error parse body from url %s : %v", url, err)
				responseErr <- err
				return
			}

			result := map[string]any{
				"url":      url,
				"response": string(body),
			}

			responseSuccess <- result
		}(u)
	}
}
