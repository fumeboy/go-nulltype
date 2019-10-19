package nulltype

import (
	"encoding/json"
	"fmt"
)

// Int64 is null friendly type for int64.
type Int64 int64

func PtrInt64Of(value int64) *Int64 {
	var s Int64
	s.Set(value)
	return &s
}

func (i *Int64) ValueWithDefault(Default int64) int64 {
	if i == nil{
		return Default
	}
	return int64(*i)
}

func (i *Int64) Value() int64 {
	if i == nil{
		panic("null string has no value")
	}
	return int64(*i)
}
// Set set the value.
func (i *Int64) Set(value int64) *Int64 {
	*i = Int64(value)
	return i
}

// String return string indicated the value.
func (i *Int64) String() string {
	if i == nil {
		return ""
	}
	return fmt.Sprint(*i)
}

// MarshalJSON encode the value to JSON.
func (i *Int64) MarshalJSON() ([]byte, error) {
	if i == nil {
		return []byte("null"), nil
	}
	return json.Marshal(*i)
}

// UnmarshalJSON decode data to the value.
func (i *Int64) UnmarshalJSON(data []byte) error {
	var value *int64
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value != nil {
		*i = Int64(*value)
	}
	return nil
}
