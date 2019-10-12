package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/ferruvich/go-exchange-rates-api/internal/rates/service"
)

// ErrBody is the standard error response body
type ErrBody struct {
	Err string `json:"error"`
}

// RecommendBody is the response body for successfull
// recommendation requests
type RecommendBody struct {
	Msg string `json:"msg"`
}

// GetRatesHandler is the handler for GET /rates
// @Summary Get Rates
// @Description returns GBP and USD exchange rates
// @Produce  json
// @Success 200 {array} rates.BasedRates base rate object
// @Failure 500 {object} ErrBody
// @Router /rates [get]
func GetRatesHandler(s service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := s.CurrentGBPUSDRates()
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrBody{
				Err: "error getting exchange rates",
			})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

// GetEURValue is the handler for GET /value/:currency
// @Summary Get EUR Values
// @Description returns the EUR value for the given currency
// @Param currency path string true "Currency"
// @Produce json
// @Success 200 {array} rates.BasedRates base rate object
// @Failure 400 {object} ErrBody
// @Failure 500 {object} ErrBody
// @Router /value/{currency} [get]
func GetEURValue(s service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		currency := c.Param("currency")
		res, err := s.CurrentEURRate(currency)
		if err != nil {
			if service.ErrInvalidParam == errors.Cause(err) {
				c.JSON(http.StatusBadRequest, ErrBody{
					Err: "bad currency",
				})
				return
			}

			c.JSON(http.StatusInternalServerError, ErrBody{
				Err: "error getting exchange rates",
			})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

// Recommend is the handler for GET /recommendation/:currency
// @Summary Get Exchange Recommendation
// @Description returns a recommendation as to whether this is a good time to exchange money or not
// @Param currency path string true "Currency"
// @Produce json
// @Success 200 {object} RecommendBody
// @Failure 400 {object} ErrBody
// @Failure 500 {object} ErrBody
// @Router /recommendation/{currency} [get]
func Recommend(s service.Servicer) gin.HandlerFunc {
	return func(c *gin.Context) {
		currency := c.Param("currency")
		res, err := s.RecommendEURExchange(currency)
		if err != nil {
			if service.ErrInvalidParam == errors.Cause(err) {
				c.JSON(http.StatusBadRequest, ErrBody{
					Err: "bad currency",
				})
				return
			}

			if service.ErrNotEnoughData == errors.Cause(err) {
				c.JSON(http.StatusInternalServerError, ErrBody{
					Err: "no enough data to make recommendations",
				})
				return
			}

			c.JSON(http.StatusInternalServerError, ErrBody{
				Err: "error getting exchange rates",
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
