package main

import (
	"flag"
	gohttp "net/http"

	"github.com/ferruvich/go-exchange-rates-api/internal/transport/http"
	http_service "github.com/ferruvich/go-exchange-rates-api/internal/transport/http/service"

	rates_repo "github.com/ferruvich/go-exchange-rates-api/internal/rates/repository"
	rates_service "github.com/ferruvich/go-exchange-rates-api/internal/rates/service"
)

const (
	defaultAddr = ":8000"
)

func main() {
	addr := flag.String("addr", defaultAddr, "address to run server")
	flag.Parse()

	ratesSvc := rates_service.New(
		rates_repo.New(
			http_service.New(&gohttp.Client{}),
		),
	)

	router := http.Routes(ratesSvc)
	router.Run(*addr)
}
