package forecast

import (
	"encoding/json"
)

func (f *Forecast) String() string {
	b, err := json.Marshal(f)
	if err != nil {
		return err.Error()
	}

	return string(b)
}
