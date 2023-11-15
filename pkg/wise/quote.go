package wise

import "time"

type Quote struct {
	ID                            string           `json:"id"`
	SourceCurrency                string           `json:"sourceCurrency"`
	TargetCurrency                string           `json:"targetCurrency"`
	SourceAmount                  float64          `json:"sourceAmount"`
	PayOut                        string           `json:"payOut"`
	PreferredPayIn                string           `json:"preferredPayIn"`
	Rate                          float64          `json:"rate"`
	CreatedTime                   time.Time        `json:"createdTime"`
	User                          int              `json:"user"`
	Profile                       int              `json:"profile"`
	RateType                      string           `json:"rateType"`
	RateExpirationTime            time.Time        `json:"rateExpirationTime"`
	GuaranteedTargetAmountAllowed bool             `json:"guaranteedTargetAmountAllowed"`
	TargetAmountAllowed           bool             `json:"targetAmountAllowed"`
	GuaranteedTargetAmount        bool             `json:"guaranteedTargetAmount"`
	ProvidedAmountType            string           `json:"providedAmountType"`
	PaymentOptions                []PaymentOptions `json:"paymentOptions"`
	Status                        string           `json:"status"`
	ExpirationTime                time.Time        `json:"expirationTime"`
	Notices                       []Notices        `json:"notices"`
}

type EstimatedDeliveryDelays struct {
	Reason string `json:"reason"`
}

type Fee struct {
	Transferwise float64 `json:"transferwise"`
	PayIn        float64 `json:"payIn"`
	Discount     float64 `json:"discount"`
	Partner      float64 `json:"partner"`
	Total        float64 `json:"total"`
}

type Value struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Label    string  `json:"label:"`
}

type Total struct {
	Type  string `json:"type"`
	Label string `json:"label"`
	Value Value  `json:"value"`
}

type Explanation struct {
	PlainText string `json:"plainText"`
}

type Items struct {
	Type        string      `json:"type"`
	Label       string      `json:"label"`
	Value       Value       `json:"value"`
	ID          int         `json:"id,omitempty"`
	Explanation Explanation `json:"explanation,omitempty"`
}

type PaymentPrice struct {
	PriceSetID int     `json:"priceSetId"`
	Total      Total   `json:"total"`
	Items      []Items `json:"items"`
}

type DisabledReason struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type PaymentOptions struct {
	Disabled                   bool                      `json:"disabled"`
	EstimatedDelivery          time.Time                 `json:"estimatedDelivery"`
	FormattedEstimatedDelivery string                    `json:"formattedEstimatedDelivery"`
	EstimatedDeliveryDelays    []EstimatedDeliveryDelays `json:"estimatedDeliveryDelays"`
	Fee                        Fee                       `json:"fee"`
	Price                      PaymentPrice              `json:"price"`
	SourceAmount               float64                   `json:"sourceAmount"`
	TargetAmount               float64                   `json:"targetAmount"`
	SourceCurrency             string                    `json:"sourceCurrency"`
	TargetCurrency             string                    `json:"targetCurrency"`
	PayIn                      string                    `json:"payIn"`
	PayOut                     string                    `json:"payOut"`
	AllowedProfileTypes        []string                  `json:"allowedProfileTypes"`
	PayInProduct               string                    `json:"payInProduct"`
	FeePercentage              float64                   `json:"feePercentage"`
	DisabledReason             DisabledReason            `json:"disabledReason,omitempty"`
}

type Notices struct {
	Text string      `json:"text"`
	Link interface{} `json:"link"`
	Type string      `json:"type"`
}
