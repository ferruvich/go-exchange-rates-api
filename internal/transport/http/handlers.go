package http

import (
	"fmt"
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

// GetEURValue is the handler for GET /value/:currency
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

// Recommend is the handler for GET /recommendation/:currency
func Recommend(s service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		currency := c.Param("currency")
		res, err := s.RecommendEURExchange(currency)
		if err != nil {
			if service.ErrInvalidParam == errors.Cause(err) {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "bad currency",
				})
				return
			}

			if service.ErrNotEnoughData == errors.Cause(err) {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "no enough data to make recommendations",
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error getting exchange rates",
			})
			return
		}

		if res {
			c.JSON(http.StatusOK, gin.H{
				"msg": fmt.Sprintf(
					"it is a good time to exchange %s for EUR", currency,
				),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf(
				"it is not a good time to exchange %s for EUR", currency,
			),
		})
	}
}
