package wise

import (
	"fmt"
	"net/url"
	"strconv"
)

type PriceRequest struct {
	SourceAmount   float64 `json:"sourceAmount"`
	TargetAmount   float64 `json:"targetAmount"`
	SourceCurrency string  `json:"sourceCurrency"`
	TargetCurrency string  `json:"targetCurrency"`
	ProfileID      int     `json:"profileId"`
	ProfileCountry string  `json:"profileCountry"`
	ProfileType    string  `json:"profileType"`
	Markers        string  `json:"markers"`
}

func (r *PriceRequest) Values() url.Values {
	values := url.Values{}

	if r.SourceAmount != 0 {
		values.Add("sourceAmount", strconv.FormatFloat(r.SourceAmount, 'f', 2, 64))
	}
	if r.TargetAmount != 0 {
		values.Add("targetAmount", strconv.FormatFloat(r.TargetAmount, 'f', 2, 64))
	}
	values.Add("sourceCurrency", r.SourceCurrency)
	values.Add("targetCurrency", r.TargetCurrency)

	if r.ProfileID != 0 {
		values.Add("profileId", strconv.Itoa(r.ProfileID))
	}

	if r.ProfileCountry != "" {
		values.Add("profileCountry", r.ProfileCountry)
	}

	if r.ProfileType == "" {
		r.ProfileType = "PERSONAL"
	}
	values.Add("profileType", r.ProfileType)

	if r.Markers == "" {
		r.Markers = "FCF_PRICING"
	}
	values.Add("markers", r.Markers)
	return values
}

type PriceResponse []Price

type Price struct {
	PriceSetID           int              `json:"priceSetId"`
	SourceAmount         float64          `json:"sourceAmount"`
	TargetAmount         float64          `json:"targetAmount"`
	PayInMethod          string           `json:"payInMethod"`
	PayOutMethod         string           `json:"payOutMethod"`
	SourceCurrency       string           `json:"sourceCcy"`
	TargetCurrency       string           `json:"targetCcy"`
	Total                float64          `json:"total"`
	VariableFee          float64          `json:"variableFee"`
	VariableFeePercent   float64          `json:"variableFeePercent"`
	SwiftPayoutFlatFee   float64          `json:"swiftPayoutFlatFee"`
	FlatFee              float64          `json:"flatFee"`
	MidRate              float64          `json:"midRate"`
	EcbRate              float64          `json:"ecbRate"`
	EcbRateTimestamp     int              `json:"ecbRateTimestamp"`
	EcbMarkupPercent     any              `json:"ecbMarkupPercent"`
	AdditionalFeeDetails TaxDetailsByType `json:"additionalFeeDetails"`
}

type TaxDetailsByType struct {
	ReceiveInr       any `json:"receiveInr"`
	TaxDetailsByType any `json:"taxDetailsByType"`
}

func (p PriceResponse) GooglePayInBalanceOut() (*Price, error) {
	for _, price := range p {
		if price.PayInMethod == "GOOGLE_PAY" && price.PayOutMethod == "BALANCE" {
			return &price, nil
		}
	}

	return nil, fmt.Errorf("method not found")
}

func (p PriceResponse) DirectDebitInBalanceOut() (*Price, error) {
	for _, price := range p {
		if price.PayInMethod == "DIRECT_DEBIT" && price.PayOutMethod == "BALANCE" {
			return &price, nil
		}
	}

	return nil, fmt.Errorf("method not found")
}

func (p PriceResponse) BankTransferInBalanceOut() (*Price, error) {
	for _, price := range p {
		if price.PayInMethod == "BANK_TRANSFER" && price.PayOutMethod == "BALANCE" {
			return &price, nil
		}
	}

	return nil, fmt.Errorf("method not found")
}
