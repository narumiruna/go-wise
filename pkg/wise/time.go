package wise

import (
	"encoding/json"
	"time"
)

type Time time.Time

func (t Time) String() string {
	return time.Time(t).String()
}

func (t *Time) UnmarshalJSON(o []byte) error {
	var timestamp int64
	if err := json.Unmarshal(o, &timestamp); err != nil {
		return err
	}

	*t = Time(time.UnixMilli(timestamp))
	return nil
}
