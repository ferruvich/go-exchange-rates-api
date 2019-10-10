package repository_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/ferruvich/go-exchange-rates-api/internal/rates"
	"github.com/ferruvich/go-exchange-rates-api/internal/rates/repository"
	service_mock "github.com/ferruvich/go-exchange-rates-api/internal/transport/http/service/mock"
)

const (
	reqTimeoutSecs = 5
)

type errorReader struct{}

func (r *errorReader) Read(p []byte) (int, error) {
	return 0, errors.New("error")
}

func TestRepository_DailyRates(t *testing.T) {
	const (
		method = "GET"
		url    = "someURL"
		base   = "EUR"
	)

	t.Run("should return error due to invalid base param", func(t *testing.T) {
		repo := repository.New(nil)
		res, err := repo.DailyRates("")
		assert.Error(t, err)
		assert.Nil(t, res)
		assert.Equal(t, repository.ErrInvalidParam, errors.Cause(err))
	})

	t.Run("should return error due to problem on NewRequest", func(t *testing.T) {
		controller := gomock.NewController(t)
		controller.Finish()

		mockHttpService := service_mock.NewMockServicer(controller)
		mockHttpService.EXPECT().NewRequest(
			method, gomock.Any(), nil, map[string]string{"base": base},
		).Return(nil, errors.New("error"))

		repo := repository.New(mockHttpService)
		res, err := repo.DailyRates(base)
		assert.Error(t, err)
		assert.Nil(t, res)
		assert.Equal(t, repository.ErrRequest, errors.Cause(err))
	})

	t.Run("should return error due to problem on Do", func(t *testing.T) {
		controller := gomock.NewController(t)
		controller.Finish()

		mockHttpService := service_mock.NewMockServicer(controller)
		mockHttpService.EXPECT().NewRequest(
			method, gomock.Any(), nil, map[string]string{"base": base},
		).Return(new(http.Request), nil)
		mockHttpService.EXPECT().Do(new(http.Request)).Return(nil, errors.New("error"))

		repo := repository.New(mockHttpService)
		res, err := repo.DailyRates(base)
		assert.Error(t, err)
		assert.Nil(t, res)
		assert.Equal(t, repository.ErrRequest, errors.Cause(err))
	})

	t.Run("should return error due to problem with ioutil.ReadAll", func(t *testing.T) {
		controller := gomock.NewController(t)
		controller.Finish()

		mockHttpService := service_mock.NewMockServicer(controller)
		mockHttpService.EXPECT().NewRequest(
			method, gomock.Any(), nil, map[string]string{"base": base},
		).Return(new(http.Request), nil)
		mockHttpService.EXPECT().Do(new(http.Request)).Return(&http.Response{
			Body: ioutil.NopCloser(new(errorReader)),
		}, nil)

		repo := repository.New(mockHttpService)
		res, err := repo.DailyRates(base)
		assert.Error(t, err)
		assert.Nil(t, res)
		assert.Equal(t, repository.ErrResponse, errors.Cause(err))
	})

	t.Run("should return error unmarshalling an empty body", func(t *testing.T) {
		controller := gomock.NewController(t)
		controller.Finish()

		mockHttpService := service_mock.NewMockServicer(controller)
		mockHttpService.EXPECT().NewRequest(
			method, gomock.Any(), nil, map[string]string{"base": base},
		).Return(new(http.Request), nil)
		mockHttpService.EXPECT().Do(new(http.Request)).Return(&http.Response{
			Body: ioutil.NopCloser(bytes.NewReader([]byte(""))),
		}, nil)

		repo := repository.New(mockHttpService)
		res, err := repo.DailyRates(base)
		assert.Error(t, err)
		assert.Nil(t, res)
		assert.Equal(t, repository.ErrResponse, errors.Cause(err))
	})

	t.Run("should return response successfully", func(t *testing.T) {
		controller := gomock.NewController(t)
		controller.Finish()

		r := new(rates.BasedRates)
		rb, _ := json.Marshal(r)

		mockHttpService := service_mock.NewMockServicer(controller)
		mockHttpService.EXPECT().NewRequest(
			method, gomock.Any(), nil, map[string]string{"base": base},
		).Return(new(http.Request), nil)
		mockHttpService.EXPECT().Do(new(http.Request)).Return(&http.Response{
			Body: ioutil.NopCloser(bytes.NewReader(rb)),
		}, nil)

		repo := repository.New(mockHttpService)
		res, err := repo.DailyRates(base)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, *r, *res)
	})
}

func TestNew(t *testing.T) {
	t.Run("should return new repository", func(t *testing.T) {
		repo := repository.New(nil)
		assert.NotNil(t, repo)
	})
}
