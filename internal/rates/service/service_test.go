package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ferruvich/go-exchange-rates-api/internal/rates/service"
)

func TestService_CurrentGBPUSDRates(t *testing.T) {
	// TODO
}

func TestService_CurrentEURRate(t *testing.T) {
	// TODO
}

func TestService_RecommendEURExchange(t *testing.T) {
	// TODO
}

func TestNew(t *testing.T) {
	t.Run("should return new repository", func(t *testing.T) {
		s := service.New(nil)
		assert.NotNil(t, s)
		assert.Implements(t, (*service.Servicer)(nil), s)
	})
}
