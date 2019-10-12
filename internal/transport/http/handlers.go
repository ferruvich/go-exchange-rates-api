package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

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

// GetEURValue returns the value of the currency
// in path parameters, expressed in EUR
func GetEURValue(s service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		currency := c.Param("currency")
		res, err := s.CurrentEURRate(currency)
		if err != nil {
			if service.ErrInvalidParam == errors.Cause(err) {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "bad currency",
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error getting exchange rates",
			})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
