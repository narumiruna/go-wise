package v1

type Rate struct {
	Rate   float64 `json:"rate"`
	Target string  `json:"target"`
	Source string  `json:"source"`
	Time   Time    `json:"time"`
}
