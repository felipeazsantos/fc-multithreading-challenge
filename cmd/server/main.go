package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/felipeazsantos/pos-goexpert/fc-multithread-challenge/configs"
	"github.com/felipeazsantos/pos-goexpert/fc-multithread-challenge/internal/api"
)

func main() {
	var cep string
	flag.StringVar(&cep, "cep", "01000000", "cep to search address on external apis")
	flag.Parse()

	apiResponse := make(chan map[string]any)
	apiResponseErr := make(chan error)

	conf := configs.InitLogs()
	cepAPI := api.NewCepAPIS(conf)

	conf.LInfo().Printf("start search address from cep: %s\n", cep)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*1))
	defer cancel()

	cepAPI.GetAddressByCep(ctx, cep, apiResponse, apiResponseErr)

	select {
	case apiResp := <-apiResponse:
		fmt.Printf("api url: %s\n", apiResp["url"])
		fmt.Printf("address found: %s\n", apiResp["response"])
	case apiErr := <-apiResponseErr:
		conf.LErr().Println(apiErr)
	case ctxErr := <-ctx.Done():
		conf.LErr().Printf("context error: %v", ctxErr)
	}

}
