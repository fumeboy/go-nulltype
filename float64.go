package main

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

func (f *Float64) ValueWithDefault(Default float64) float64 {
	if f == nil{
		return Default
	}
	return float64(*f)
}

func (f *Float64) Value() float64 {
	if f == nil{
		panic("null string has no value")
	}
	return float64(*f)
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