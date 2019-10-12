package http

import (
	"github.com/gin-gonic/gin"

	"github.com/ferruvich/go-exchange-rates-api/internal/rates/service"
)

// Routes returns our application routes
func Routes(
	s service.Servicer,
) *gin.Engine {
	router := gin.Default()

	router.GET("/rates", GetRatesHandler(s))
	router.GET("/value/:currency", GetEURValue(s))
	router.GET("/recommendation/:currency", Recommend(s))

	return router
}
