package repository

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/ferruvich/go-exchange-rates-api/internal/rates"
	internal_http "github.com/ferruvich/go-exchange-rates-api/internal/transport/http/service"
	"github.com/pkg/errors"
)

var (
	// ErrInvalidParam is used when there's an invalid parameter
	// on functions input
	ErrInvalidParam = errors.New("invalid_parameter")
	// ErrRequest is used when there's an error
	// creating/sending HTTP request
	ErrRequest = errors.New("request_error")
	// ErrResponse is used when there's an error
	// reading/unmarshalling response
	ErrResponse = errors.New("response_error")
)

// Repositorer is the repo interface
type Repositorer interface {
	DailyRates(base string) (*rates.BasedRates, error)
	HistoricalRates(base string) (*rates.HistoricalRates, error)
	SpecificRates(base, conversion string) (*rates.BasedRates, error)
}

// Repository is the Repositorer implementation
type Repository struct {
	httpSvc internal_http.Servicer
	baseURL string
}

// DailyRates returns the daily exchange rates
// for the given 'base' currency
func (r *Repository) DailyRates(base string) (*rates.BasedRates, error) {
	if base == "" {
		return nil, errors.Wrap(ErrInvalidParam, "base")
	}

	req, err := r.httpSvc.NewRequest(
		"GET",
		strings.Join([]string{
			r.baseURL, "latest",
		}, "/"),
		nil,
		map[string]string{
			"base": base,
		},
	)
	if err != nil {
		return nil, errors.Wrap(ErrRequest, err.Error())
	}

	resp, err := r.httpSvc.Do(req)
	if err != nil {
		return nil, errors.Wrap(ErrRequest, err.Error())
	}
	defer resp.Body.Close()

	respb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(ErrResponse, err.Error())
	}

	res := new(rates.BasedRates)
	err = json.Unmarshal(respb, res)
	if err != nil {
		return nil, errors.Wrap(ErrResponse, err.Error())
	}

	return res, nil
}

// New initializes a new rates service
func New(c internal_http.Servicer) *Repository {
	return &Repository{
		httpSvc: c,
		baseURL: "https://api.exchangeratesapi.io",
	}
}
