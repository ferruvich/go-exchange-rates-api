package http_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ferruvich/go-exchange-rates-api/internal/transport/http"
)

func TestRoutes(t *testing.T) {
	t.Run("should return routes", func(t *testing.T) {
		r := http.Routes("", nil)

		assert.NotNil(t, r)
	})
}
