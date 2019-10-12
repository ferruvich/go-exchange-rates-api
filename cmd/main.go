package main

import (
	gohttp "net/http"

	"github.com/ferruvich/go-exchange-rates-api/internal/transport/http"
	http_service "github.com/ferruvich/go-exchange-rates-api/internal/transport/http/service"

	rates_repo "github.com/ferruvich/go-exchange-rates-api/internal/rates/repository"
	rates_service "github.com/ferruvich/go-exchange-rates-api/internal/rates/service"
)

const (
	defaultAddr = ":8000"
)

// @title Exchange rates API
// @version 1.0
// @description API to retrieve exchange rates, and make recommendations

// @host localhost:8000
// @BasePath /
func main() {

	ratesSvc := rates_service.New(
		rates_repo.New(
			http_service.New(&gohttp.Client{}),
		),
	)

	router := http.Routes(defaultAddr, ratesSvc)
	router.Run(defaultAddr)
}
