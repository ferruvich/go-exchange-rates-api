package service

import (
	"time"

	"github.com/pkg/errors"

	"github.com/DzananGanic/numericalgo/fit/linear"
	"github.com/ferruvich/go-exchange-rates-api/internal/rates"
	"github.com/ferruvich/go-exchange-rates-api/internal/rates/repository"
)

const (
	gbp = "GBP"
	eur = "EUR"
	usd = "USD"

	dateFormat = "2006-01-02"
)

var (
	// ErrRepo is used when there's an error on Repository
	ErrRepo = errors.New("repo_error")
	// ErrInvalidParam is used when an invalid parameter is passed
	ErrInvalidParam = errors.New("invalid_parameter")
	// ErrNotEnoughData is used on RecommendEURExchange if there's not
	// enough data to make a recommendation
	ErrNotEnoughData = errors.New("not_enough_data")
)

//go:generate mockgen -source=service.go -destination=mock/service_mock.go -package=service_mock -self_package=. Servicer

// Servicer is the rates service interface
type Servicer interface {
	CurrentGBPUSDRates() ([]*rates.BasedRates, error)
	CurrentEURRate(currency string) (*rates.BasedRates, error)
	RecommendEURExchange(currency string) (bool, error)
}

// Service is the Servicer implementation
type Service struct {
	repo repository.Repositorer
}

// CurrentGBPUSDRates returns the latest exchange rates for the base
// currencies of GBP and USD
func (s *Service) CurrentGBPUSDRates() ([]*rates.BasedRates, error) {
	gbpRates, err := s.repo.CurrentRates(gbp)
	if err != nil {
		return nil, errors.Wrap(ErrRepo, err.Error())
	}

	usdRates, err := s.repo.CurrentRates(usd)
	if err != nil {
		return nil, errors.Wrap(ErrRepo, err.Error())
	}

	return []*rates.BasedRates{gbpRates, usdRates}, nil
}

// CurrentEURRate returns the 'currency' value in euros
func (s *Service) CurrentEURRate(currency string) (*rates.BasedRates, error) {
	rates, err := s.repo.CurrentSpecificRates(currency, []string{eur})
	if err != nil {
		if repository.ErrInvalidParam == errors.Cause(err) {
			return nil, errors.Wrap(ErrInvalidParam, err.Error())
		}
		return nil, errors.Wrap(ErrRepo, err.Error())
	}

	return rates, nil
}

// RecommendEURExchange makes a naive recommendation as to whether
// this is a good time to exchange amounts of the 'currency' with euros.
// The recommendation value is true if we say that is a good time to exchange money,
// false otherwise (or when an error is returned)
func (s *Service) RecommendEURExchange(currency string) (bool, error) {
	now := time.Now()
	weekAgo := time.Now().AddDate(0, 0, -7)
	rates, err := s.repo.HistoricalSpecificRates(
		currency, weekAgo.Format(dateFormat),
		now.Format(dateFormat), []string{eur},
	)
	if err != nil {
		if repository.ErrInvalidParam == errors.Cause(err) {
			return false, errors.Wrap(ErrInvalidParam, err.Error())
		}
		return false, errors.Wrap(ErrRepo, err.Error())
	}

	if len(rates.Rates) <= 1 {
		return false, ErrNotEnoughData
	}

	// We use single value prediction in order to make a naive
	// recommendation
	x := []float64{}
	y := []float64{}
	i := 1.0
	for d := weekAgo; d.Day() != now.Day(); d = d.AddDate(0, 0, 1) {
		res := rates.Rates[d.Format("2006-01-02")][eur]
		if res > 0 {
			x = append(x, i)
			y = append(y, res)
			i += 1.0
		}
	}

	li := linear.New()
	li.Fit(x, y)
	estimate := li.Predict(i)
	// We recommend to exchange since 'base' currency
	// money value may decrease
	if estimate <= y[len(y)-1] {
		return true, nil
	}
	// We recommend to not exchange since 'base' currency
	// money value may increase
	return false, nil
}

// New initializes a new rates service
func New(repo repository.Repositorer) *Service {
	return &Service{
		repo: repo,
	}
}
