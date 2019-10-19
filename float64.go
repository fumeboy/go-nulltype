package nulltype

import (
	"encoding/json"
	"fmt"
)

// Float64 is null friendly type for float64.
type Float64 float64

func PtrFloat64Of(value float64) *Float64 {
	var s Float64
	s.Set(value)
	return &s
}

// Set set the value.
func (f *Float64) Set(value float64) *Float64 {
	*f = Float64(value)
	return f
}

// String return string indicated the value.
func (f *Float64) String() string {
	if f == nil {
		return ""
	}
	return fmt.Sprint(*f)
}

// MarshalJSON encode the value to JSON.
func (f *Float64) MarshalJSON() ([]byte, error) {
	if f == nil {
		return []byte("null"), nil
	}
	return json.Marshal(*f)
}

// UnmarshalJSON decode data to the value.
func (f *Float64) UnmarshalJSON(data []byte) error {
	var value *float64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value != nil {
		*f = Float64(*value)
	}
	return nil
}