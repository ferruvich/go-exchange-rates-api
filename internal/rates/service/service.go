package service

import (
	"github.com/pkg/errors"

	"github.com/ferruvich/go-exchange-rates-api/internal/rates"
	"github.com/ferruvich/go-exchange-rates-api/internal/rates/repository"
)

const (
	gbp = "GBP"
	usd = "USD"
)

var (
	// ErrRepo is used when there's an error on Repository
	ErrRepo = errors.New("repo_error")
)

// Servicer is the rates service interface
type Servicer interface {
	CurrentGBPUSDRates() ([]*rates.BasedRates, error)
	CurrentEURRate(currency string) (*rates.BasedRates, error)
	RecommendEURExchange(currency string) error
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
	// TODO
	return nil, nil
}

// RecommendEURExchange makes a naive recommendation as to whether
// this is a good time to exchange amounts of the 'currency' with euros
func (s *Service) RecommendEURExchange(currency string) error {
	// TODO
	return nil
}

// New initializes a new rates service
func New(repo repository.Repositorer) *Service {
	return &Service{
		repo: repo,
	}
}
