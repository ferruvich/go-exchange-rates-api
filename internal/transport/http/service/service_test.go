package service_test

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/ferruvich/go-exchange-rates-api/internal/transport/http/service"
	service_mock "github.com/ferruvich/go-exchange-rates-api/internal/transport/http/service/mock"
)

func TestService_NewRequest(t *testing.T) {
	t.Run("should return error due to problem on NewRequest", func(t *testing.T) {
		s := service.New(nil)

		res, err := s.NewRequest("bad method", "", nil, nil)
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("should return the request", func(t *testing.T) {
		s := service.New(nil)

		res, err := s.NewRequest("POST", "http://www.foo.com", nil, map[string]string{
			"some": "value",
		})
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
}

func TestService_Do(t *testing.T) {
	t.Run("should return error due to problem on Doer", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockDoer := service_mock.NewMockDoer(controller)
		mockDoer.EXPECT().Do(gomock.Any()).Return(nil, errors.New("error"))

		s := service.New(mockDoer)
		res, err := s.Do(new(http.Request))
		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("should return response", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		mockDoer := service_mock.NewMockDoer(controller)
		mockDoer.EXPECT().Do(gomock.Any()).Return(new(http.Response), nil)

		s := service.New(mockDoer)
		res, err := s.Do(new(http.Request))
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
}

func TestNew(t *testing.T) {
	t.Run("should return new service", func(t *testing.T) {
		s := service.New(nil)
		assert.NotNil(t, s)
		assert.Implements(t, (*service.Servicer)(nil), s)
	})
}
