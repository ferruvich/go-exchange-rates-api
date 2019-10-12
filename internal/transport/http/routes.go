package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ferruvich/go-exchange-rates-api/internal/rates/service"
)

// Routes returns our application routes
func Routes(
	s service.Servicer,
) *gin.Engine {
	router := gin.Default()

	router.GET("/rates", GetRatesHandler(s))
	router.GET("/eur-rates/:currency", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{})
	})
	router.GET("/recommend/:currency", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{})
	})

	return router
}
