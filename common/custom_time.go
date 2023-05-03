package common

import (
	"encoding/json"
	"strings"
	"time"
)

type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) error {
	dd, err := time.Parse("2006-01-02", strings.Trim(string(b), "\""))
	if err != nil {
		return NewServerError(err)
	}

	*d = Date(dd)

	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d))
}
