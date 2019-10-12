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
	t.Run("should return error due to invalid param passed to SpecificRates", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepo := repository_mock.NewMockRepositorer(controller)
		mockRepo.EXPECT().CurrentSpecificRates(
			"", []string{eur},
		).Return(nil, errors.Wrap(repository.ErrInvalidParam, "error"))

		s := service.New(mockRepo)
		res, err := s.CurrentEURRate("")
		assert.Error(t, err)
		assert.Nil(t, res)
		assert.Equal(t, service.ErrInvalidParam, errors.Cause(err))
	})

	t.Run("should return error due to problem on SpecificRates", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepo := repository_mock.NewMockRepositorer(controller)
		mockRepo.EXPECT().CurrentSpecificRates(
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
		mockRepo.EXPECT().CurrentSpecificRates(
			gbp, []string{eur},
		).Return(new(rates.BasedRates), nil)

		s := service.New(mockRepo)
		res, err := s.CurrentEURRate(gbp)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
}

func TestService_RecommendEURExchange(t *testing.T) {
	t.Run("should return error due to invalid param passed to HistoricalSpecificRates", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepo := repository_mock.NewMockRepositorer(controller)
		mockRepo.EXPECT().HistoricalSpecificRates(
			gbp, gomock.Any(), gomock.Any(), []string{eur},
		).Return(nil, errors.Wrap(repository.ErrInvalidParam, "error"))

		s := service.New(mockRepo)
		res, err := s.RecommendEURExchange(gbp)
		assert.Error(t, err)
		assert.False(t, res)
		assert.Equal(t, service.ErrInvalidParam, errors.Cause(err))
	})

	t.Run("should return error due to problem on HistoricalSpecificRates", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepo := repository_mock.NewMockRepositorer(controller)
		mockRepo.EXPECT().HistoricalSpecificRates(
			gbp, gomock.Any(), gomock.Any(), []string{eur},
		).Return(nil, errors.New("error"))

		s := service.New(mockRepo)
		res, err := s.RecommendEURExchange(gbp)
		assert.Error(t, err)
		assert.False(t, res)
		assert.Equal(t, service.ErrRepo, errors.Cause(err))
	})

	t.Run("should return error due to not enough data", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockRepo := repository_mock.NewMockRepositorer(controller)
		mockRepo.EXPECT().HistoricalSpecificRates(
			gbp, gomock.Any(), gomock.Any(), []string{eur},
		).Return(new(rates.HistoricalRates), nil)

		s := service.New(mockRepo)
		res, err := s.RecommendEURExchange(gbp)
		assert.Error(t, err)
		assert.False(t, res)
		assert.Equal(t, service.ErrNotEnoughData, errors.Cause(err))
	})

	t.Run("should return true correctly", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		rates := &rates.HistoricalRates{
			Rates: map[string]map[string]float64{
				"2019-10-09": {
					eur: 1.1129660545,
				},
				"2019-10-10": {
					eur: 1.1092008208,
				},
				"2019-10-08": {
					eur: 1.1136477532,
				},
			},
			StartAt: "2019-10-05",
			Base:    gbp,
			EndAt:   "2019-10-12",
		}

		mockRepo := repository_mock.NewMockRepositorer(controller)
		mockRepo.EXPECT().HistoricalSpecificRates(
			gbp, gomock.Any(), gomock.Any(), []string{eur},
		).Return(rates, nil)

		s := service.New(mockRepo)
		res, err := s.RecommendEURExchange(gbp)
		assert.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("should return false correctly", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		rates := &rates.HistoricalRates{
			Rates: map[string]map[string]float64{
				"2019-10-09": {
					eur: 1.1129660545,
				},
				"2019-10-08": {
					eur: 1.1092008208,
				},
				"2019-10-10": {
					eur: 1.1136477532,
				},
			},
			StartAt: "2019-10-05",
			Base:    gbp,
			EndAt:   "2019-10-12",
		}

		mockRepo := repository_mock.NewMockRepositorer(controller)
		mockRepo.EXPECT().HistoricalSpecificRates(
			gbp, gomock.Any(), gomock.Any(), []string{eur},
		).Return(rates, nil)

		s := service.New(mockRepo)
		res, err := s.RecommendEURExchange(gbp)
		assert.NoError(t, err)
		assert.False(t, res)
	})
}

func TestNew(t *testing.T) {
	t.Run("should return new repository", func(t *testing.T) {
		s := service.New(nil)
		assert.NotNil(t, s)
		assert.Implements(t, (*service.Servicer)(nil), s)
	})
}
