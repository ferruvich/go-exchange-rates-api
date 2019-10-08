package rates

// BasedRates contains the rates of a specific base currency
type BasedRates struct {
	Rates map[string]float32 `json:"rates"`
	Base  string             `json:"base"`
	Date  string             `json:"date"`
}

// HistoricalRates contains the historical rates of
// a specific base currency
type HistoricalRates struct {
	Rates   map[string]map[string]interface{} `json:"rates"`
	Base    string                            `json:"base"`
	StartAt string                            `json:"start_at"`
	EndAt   string                            `json:"end_at"`
}
