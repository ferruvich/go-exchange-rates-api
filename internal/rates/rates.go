package rates

// Rates contains the rate object
type Rates struct {
	Rate Rate   `json:"rates"`
	Base string `json:"base"`
	Date string `json:"date"`
}

// Rate represent a single rate object
// it contains USD-EUR and GBP-EUR exchange rates
type Rate struct {
	GBP float32 `json:"GBP"`
	UDS float32 `json:"USD"`
}
