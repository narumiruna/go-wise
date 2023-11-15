package wise

type Rate struct {
	Source string    `json:"source"`
	Target string    `json:"target"`
	Value  float64   `json:"value"`
	Time   Timestamp `json:"time"`
}
