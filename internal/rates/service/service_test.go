package service_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/ferruvich/go-exchange-rates-api/internal/rates"
	"github.com/ferruvich/go-exchange-rates-api/internal/rates/repository"
	repository_mock "github.com/ferruvich/go-exchange-rates-api/internal/rates/repository/mock"
	"github.com/ferruvich/go-exchange-rates-api/internal/rates/service"
)

const (
	gbp = "GBP"
	usd = "USD"
	eur = "EUR"
)

func TestService_CurrentGBPUSDRates(t *testing.T) {
	t.Run("should return error due to problem on CurrentRates first call", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepo := repository_mock.NewMockRepositorer(controller)
		mockRepo.EXPECT().CurrentRates(gbp).Return(nil, errors.New("error"))

		s := service.New(mockRepo)
		res, err := s.CurrentGBPUSDRates()
		assert.Error(t, err)
		assert.Nil(t, res)
		assert.Equal(t, service.ErrRepo, errors.Cause(err))
	})

	t.Run("should return error due to problem on CurrentRates second call", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepo := repository_mock.NewMockRepositorer(controller)
		mockRepo.EXPECT().CurrentRates(gbp).Return(new(rates.BasedRates), nil)
		mockRepo.EXPECT().CurrentRates(usd).Return(nil, errors.New("error"))

		s := service.New(mockRepo)
		res, err := s.CurrentGBPUSDRates()
		assert.Error(t, err)
		assert.Nil(t, res)
		assert.Equal(t, service.ErrRepo, errors.Cause(err))
	})

	t.Run("should return rates successfully", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepo := repository_mock.NewMockRepositorer(controller)
		mockRepo.EXPECT().CurrentRates(gbp).Return(new(rates.BasedRates), nil)
		mockRepo.EXPECT().CurrentRates(usd).Return(new(rates.BasedRates), nil)

		s := service.New(mockRepo)
		res, err := s.CurrentGBPUSDRates()
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Len(t, res, 2)
	})
}

func TestService_CurrentEURRate(t *testing.T) {
	t.Run("should return error due to problem on CurrentRates", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepo := repository_mock.NewMockRepositorer(controller)
		mockRepo.EXPECT().SpecificRates(
			"", []string{eur},
		).Return(nil, errors.Wrap(repository.ErrInvalidParam, "error"))

		s := service.New(mockRepo)
		res, err := s.CurrentEURRate("")
		assert.Error(t, err)
		assert.Nil(t, res)
		assert.Equal(t, service.ErrInvalidParam, errors.Cause(err))
	})

	t.Run("should return error due to problem on CurrentRates", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepo := repository_mock.NewMockRepositorer(controller)
		mockRepo.EXPECT().SpecificRates(
			gbp, []string{eur},
		).Return(nil, errors.New("error"))

		s := service.New(mockRepo)
		res, err := s.CurrentEURRate(gbp)
		assert.Error(t, err)
		assert.Nil(t, res)
		assert.Equal(t, service.ErrRepo, errors.Cause(err))
	})

	t.Run("should return rates successfully", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepo := repository_mock.NewMockRepositorer(controller)
		mockRepo.EXPECT().SpecificRates(
			gbp, []string{eur},
		).Return(new(rates.BasedRates), nil)

		s := service.New(mockRepo)
		res, err := s.CurrentEURRate(gbp)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
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
