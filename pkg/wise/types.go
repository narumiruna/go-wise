package wise

import (
	"encoding/json"
	"time"
)

type Resolution string

const (
	ResolutionHourly Resolution = "hourly"
	ResolutionDaily  Resolution = "daily"
)

type Unit string

const (
	UnitDay   Unit = "day"
	UnitMonth Unit = "month"
	UnitYear  Unit = "year"
)

type Timestamp time.Time

func (t Timestamp) Unix() float64 {
	return float64(time.Time(t).Unix())
}

func (t Timestamp) String() string {
	return time.Time(t).String()
}

func (t *Timestamp) UnmarshalJSON(o []byte) error {
	var timestamp int64
	if err := json.Unmarshal(o, &timestamp); err != nil {
		return err
	}

	*t = Timestamp(time.UnixMilli(timestamp))
	return nil
}
