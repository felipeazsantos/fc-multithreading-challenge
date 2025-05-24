package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/felipeazsantos/pos-goexpert/fc-multithread-challenge/configs"
	"github.com/felipeazsantos/pos-goexpert/fc-multithread-challenge/internal/api"
)

func main() {
	var cep string
	flag.StringVar(&cep, "cep", "01000000", "cep to search address on external apis")
	flag.Parse()

	brasilApiResponse := make(chan bool)
	viaCepResponse := make(chan bool)
	apiResponseErr := make(chan error)

	conf := configs.InitLogs()
	cepAPI := api.NewCepAPIS(conf)

	conf.LInfo().Printf("start search address from cep: %s\n", cep)
	ctx := context.Background()

	go cepAPI.GetAddressFromBrasilCepAPI(ctx, cep, brasilApiResponse, apiResponseErr)
	go cepAPI.GetAddressFromViaCepAPI(ctx, cep, viaCepResponse, apiResponseErr)

	select {
	case brasilApi := <-brasilApiResponse:
		fmt.Println(brasilApi)
	case viaCepApi := <-viaCepResponse:
		fmt.Println(viaCepApi)
	case apiErr := <-apiResponseErr:
		fmt.Println(apiErr)
	case ctxErr := <-ctx.Done():
		fmt.Printf("context error: %v", ctxErr)
	}

}
