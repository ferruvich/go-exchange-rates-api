package service

import (
	"io"
	"net/http"

	"github.com/pkg/errors"
)

var (
	// ErrInvalidParam is used when an invalid param is passed as input
	ErrInvalidParam = errors.New("invalid_parameter")
)

//go:generate mockgen -source=service.go -destination=mock/service_mock.go -package=service_mock -self_package=. Doer,Servicer

// Doer is an interface used as a simple http.Client interface
// in order handle only interfaces as input
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Servicer is the HTTP service interface
type Servicer interface {
	Doer
	NewRequest(method, url string, body io.ReadCloser, qp map[string]interface{}) (*http.Request, error)
}

// Service is the Servicer implementation
type Service struct {
	client Doer
}

// NewRequest craft a new http.Request, given all its parameters
func (s *Service) NewRequest(
	method, url string, body io.ReadCloser, qp map[string]interface{},
) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for k, v := range qp {
		switch t := v.(type) {
		case string:
			q.Add(k, t)
		case []string:
			for _, el := range t {
				q.Add(k, el)
			}
		default:
			return nil, errors.Wrapf(ErrInvalidParam, "qp: %s", k)
		}
	}
	req.URL.RawQuery = q.Encode()

	return req, nil
}

// Do send the request and returns its response
func (s *Service) Do(req *http.Request) (*http.Response, error) {
	return s.client.Do(req)
}

// New initializes a new HTTP service
func New(c Doer) *Service {
	return &Service{
		client: c,
	}
}
