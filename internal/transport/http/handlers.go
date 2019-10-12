package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ferruvich/go-exchange-rates-api/internal/rates/service"
)

// GetRatesHandler is the handler for GET /rates
func GetRatesHandler(s service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := s.CurrentGBPUSDRates()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error getting exchange rates",
			})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
