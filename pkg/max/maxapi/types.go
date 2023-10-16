package maxapi

type Ticker struct {
	At       int    `json:"at"`
	Buy      string `json:"buy"`
	Sell     string `json:"sell"`
	Open     string `json:"open"`
	Low      string `json:"low"`
	High     string `json:"high"`
	Last     string `json:"last"`
	Vol      string `json:"vol"`
	VolInBtc string `json:"vol_in_btc"`
}
