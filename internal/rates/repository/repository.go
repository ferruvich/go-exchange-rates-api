package repository

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/ferruvich/go-exchange-rates-api/internal/rates"
	internal_http "github.com/ferruvich/go-exchange-rates-api/internal/transport/http/service"
	"github.com/pkg/errors"
)

const (
	dateFormat = "^[0-9]{4}-([0][1-9]|[1][0-2])-([0][1-9]|[1-2][0-9]|[3][0-1])$"
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

//go:generate mockgen -source=repository.go -destination=mock/repository_mock.go -package=repository_mock -self_package=. Repositorer

// Repositorer is the repo interface
type Repositorer interface {
	CurrentRates(base string) (*rates.BasedRates, error)
	HistoricalRates(base, start, end string) (*rates.HistoricalRates, error)
	SpecificRates(base string, currencies []string) (*rates.BasedRates, error)
}

// Repository is the Repositorer implementation
type Repository struct {
	httpSvc internal_http.Servicer
	baseURL string
}

// checkDate checks for a date, ignoring error thrown by regexp
func checkDate(date string) bool {
	ok, _ := regexp.Match(dateFormat, []byte(date))
	return ok
}

// CurrentRates returns the current exchange rates
// for the given 'base' currency
func (r *Repository) CurrentRates(base string) (*rates.BasedRates, error) {
	if base == "" {
		return nil, errors.Wrap(ErrInvalidParam, "base")
	}

	req, err := r.httpSvc.NewRequest(
		"GET",
		strings.Join([]string{
			r.baseURL, "latest",
		}, "/"),
		nil,
		map[string]interface{}{
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

// HistoricalRates returns the historical exchange rates
// for the given 'base' currency, starting from 'start' date and ending to 'end' date
func (r *Repository) HistoricalRates(base, start, end string) (*rates.HistoricalRates, error) {
	if base == "" {
		return nil, errors.Wrap(ErrInvalidParam, "base")
	}
	if start == "" && !checkDate(start) {
		return nil, errors.Wrap(ErrInvalidParam, "start")
	}
	if end == "" && !checkDate(end) {
		return nil, errors.Wrap(ErrInvalidParam, "end")
	}

	req, err := r.httpSvc.NewRequest(
		"GET",
		strings.Join([]string{
			r.baseURL, "history",
		}, "/"),
		nil,
		map[string]interface{}{
			"base":     base,
			"start_at": start,
			"end_at":   end,
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

	res := new(rates.HistoricalRates)
	err = json.Unmarshal(respb, res)
	if err != nil {
		return nil, errors.Wrap(ErrResponse, err.Error())
	}

	return res, nil
}

// SpecificRates returns the current exchange rate
// for the given 'base' currency and the specific 'currency'
func (r *Repository) SpecificRates(base string, currencies []string) (*rates.BasedRates, error) {
	if base == "" {
		return nil, errors.Wrap(ErrInvalidParam, "base")
	}
	if len(currencies) == 0 {
		return nil, errors.Wrap(ErrInvalidParam, "currencies")
	}

	req, err := r.httpSvc.NewRequest(
		"GET",
		strings.Join([]string{
			r.baseURL, "latest",
		}, "/"),
		nil,
		map[string]interface{}{
			"base":    base,
			"symbols": currencies,
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

// New initializes a new rates repo
func New(c internal_http.Servicer) *Repository {
	return &Repository{
		httpSvc: c,
		baseURL: "https://api.exchangeratesapi.io",
	}
}
