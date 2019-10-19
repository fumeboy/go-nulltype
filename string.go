package nulltype

import (
	"encoding/json"
)

// String is null friendly type for string.
type String string

func PtrStringOf(value string) *String {
	var s String
	s.Set(value)
	return &s
}

func (s *String) ValueWithDefault(Default string) string {
	if s == nil{
		return Default
	}
	return string(*s)
}

func (s *String) Value() string {
	return string(*s)
}
// Set set the value.
func (s *String) Set(value string) {
	*s = String(value)
}
// String return string indicated the value.
func (s *String) String() string {
	if s == nil {
		return ""
	}
	return string(*s)
}

// MarshalJSON encode the value to JSON.
func (s *String) MarshalJSON() ([]byte, error) {
	if s == nil {
		return []byte("null"), nil
	}
	return json.Marshal(*s)
}

// UnmarshalJSON decode data to the value.
func (s *String) UnmarshalJSON(data []byte) error {
	var value *string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if value == nil {
		*s = String(*value)
	}
	return nil
}
