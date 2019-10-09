package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ferruvich/go-exchange-rates-api/internal/rates/service"
)

const (
	reqTimeoutSecs = 5
)

func TestService_DailyRates(t *testing.T) {
	// TODO
}

func TestNew(t *testing.T) {
	t.Run("should return new service", func(t *testing.T) {
		s := service.New(nil)
		assert.NotNil(t, s)
	})
}
