package nulltype

import (
	"encoding/json"
	"fmt"
)

// Bool is null friendly type for bool.
type Bool bool

func PtrBoolOf(value bool) *Bool {
	var b Bool
	b.Set(value)
	return &b
}

func (b *Bool) ValueWithDefault(Default bool) bool {
	if b == nil{
		return Default
	}
	return bool(*b)
}

func (b *Bool) Value() bool {
	return bool(*b)
}

// Set set the value.
func (b *Bool) Set(value bool) *Bool {
	*b = Bool(value)
	return b
}

// String return string indicated the value.
func (b *Bool) String() string {
	fmt.Println(b == nil,12)
	if b == nil {
		return ""
	}
	if *b {
		return "true"
	}
	return "false"
}

// MarshalJSON encode the value to JSON.
func (b *Bool) MarshalJSON() ([]byte, error) {
	if b == nil {
		return []byte("null"), nil
	}
	return json.Marshal(*b)
}

// UnmarshalJSON decode data to the value.
func (b *Bool) UnmarshalJSON(data []byte) error {
	var value *bool
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	if b != nil{
		*b = true
	}
	return nil
}
