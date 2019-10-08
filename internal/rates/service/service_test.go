package service_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ferruvich/go-exchange-rates-api/internal/rates/service"
)

const (
	reqTimeoutSecs = 5
)

func TestService_DailyRates(t *testing.T) {
	t.Run("should take rates successfully", func(t *testing.T) {
		c := &http.Client{
			Timeout: time.Duration(reqTimeoutSecs * time.Second),
		}
		s := service.New(c)
		assert.NotNil(t, s)

		r, err := s.DailyRates("EUR")
		assert.NoError(t, err)
		assert.NotNil(t, r)
	})
}

func TestNew(t *testing.T) {
	t.Run("should return new service", func(t *testing.T) {
		s := service.New(nil)
		assert.NotNil(t, s)
	})
}
