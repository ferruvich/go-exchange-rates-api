package http_test

import (
	"errors"
	gohttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"gopkg.in/go-playground/assert.v1"

	"github.com/ferruvich/go-exchange-rates-api/internal/rates"
	service_mock "github.com/ferruvich/go-exchange-rates-api/internal/rates/service/mock"
	"github.com/ferruvich/go-exchange-rates-api/internal/transport/http"
)

func TestGetRatesHandler(t *testing.T) {
	t.Run("should return 500 INTERNAL SERVER ERROR due to service error", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockSvc := service_mock.NewMockServicer(controller)
		mockSvc.EXPECT().CurrentGBPUSDRates().Return(nil, errors.New("error"))

		router := http.Routes(mockSvc)

		r := httptest.NewRequest(
			gohttp.MethodGet, "/rates", nil,
		)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)
		assert.Equal(t, gohttp.StatusInternalServerError, w.Code)
	})

	t.Run("should return 200 OK", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockSvc := service_mock.NewMockServicer(controller)
		mockSvc.EXPECT().CurrentGBPUSDRates().Return(
			[]*rates.BasedRates{}, nil,
		)

		router := http.Routes(mockSvc)

		r := httptest.NewRequest(
			gohttp.MethodGet, "/rates", nil,
		)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)
		assert.Equal(t, gohttp.StatusOK, w.Code)
	})
}
