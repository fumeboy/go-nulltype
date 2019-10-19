package nulltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"time"
)

// Time is null friendly type for string.
type Time time.Time

func PtrTimeOf(value time.Time) *Time {
	var t Time
	t.Set(value)
	return &t
}

// Set set the value.
func (t *Time) Set(value time.Time) {
	*t = Time(value)
}

var timestampFormats = []string{
	"2006-01-02 15:04:05.999999999-07:00",
	"2006-01-02T15:04:05.999999999-07:00",
	"2006-01-02 15:04:05.999999999",
	"2006-01-02T15:04:05.999999999",
	"2006-01-02 15:04:05",
	"2006-01-02T15:04:05",
	"2006-01-02 15:04",
	"2006-01-02T15:04",
	"2006-01-02",
	"2006/01/02 15:04:05",
}
// Time return string indicated the value.
func (t *Time) String() string {
	if t == nil {
		return ""
	}
	tt := Time(*t)
	return tt.Format("2006/01/02 15:04:05")
}

// MarshalJSON encode the value to JSON.
func (t *Time) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte("null"), nil
	}
	return json.Marshal(t.t.Format(time.RFC3339))
}

// UnmarshalJSON decode data to the value.
func (t *Time) UnmarshalJSON(data []byte) error {
	var value *string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		tt, err := time.Parse(time.RFC3339, *value)
		if err != nil {
			return err
		}
		t.t = tt
	}
	return nil
}