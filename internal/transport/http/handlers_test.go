package http_test

import (
	gohttp "net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/assert.v1"

	"github.com/ferruvich/go-exchange-rates-api/internal/rates"
	"github.com/ferruvich/go-exchange-rates-api/internal/rates/service"
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

func TestGetEURValue(t *testing.T) {
	currency := "GBP"

	t.Run("should return 400 BAD REQUEST due to bad currency", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockSvc := service_mock.NewMockServicer(controller)
		mockSvc.EXPECT().CurrentEURRate(currency).Return(
			nil, errors.Wrap(service.ErrInvalidParam, "error"),
		)

		router := http.Routes(mockSvc)

		route := strings.Join([]string{
			"/value", currency,
		}, "/")
		r := httptest.NewRequest(
			gohttp.MethodGet, route, nil,
		)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)
		assert.Equal(t, gohttp.StatusBadRequest, w.Code)
	})

	t.Run("should return 500 INTERNAL SERVER ERROR due to service error", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockSvc := service_mock.NewMockServicer(controller)
		mockSvc.EXPECT().CurrentEURRate(currency).Return(nil, errors.New("error"))

		router := http.Routes(mockSvc)

		route := strings.Join([]string{
			"/value", currency,
		}, "/")
		r := httptest.NewRequest(
			gohttp.MethodGet, route, nil,
		)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)
		assert.Equal(t, gohttp.StatusInternalServerError, w.Code)
	})

	t.Run("should return 200 OK", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockSvc := service_mock.NewMockServicer(controller)
		mockSvc.EXPECT().CurrentEURRate(currency).Return(
			new(rates.BasedRates), nil,
		)

		router := http.Routes(mockSvc)

		route := strings.Join([]string{
			"/value", currency,
		}, "/")
		r := httptest.NewRequest(
			gohttp.MethodGet, route, nil,
		)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)
		assert.Equal(t, gohttp.StatusOK, w.Code)
	})
}
