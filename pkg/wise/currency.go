package wise

type Currency struct {
	Code             string   `json:"code"`
	Symbol           string   `json:"symbol"`
	Name             string   `json:"name"`
	CountryKeywords  []string `json:"countryKeywords"`
	SupportsDecimals bool     `json:"supportsDecimals"`
}
