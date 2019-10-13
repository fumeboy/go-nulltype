package nulltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// Bool is null friendly type for bool.
type Bool struct {
	b sql.NullBool
}

// BoolOf return Bool that he value is set.
func BoolOf(value bool) Bool {
	var b Bool
	b.Set(value)
	return b
}

func PtrBoolOf(value bool) *Bool {
	var b Bool
	b.Set(value)
	return &b
}

// Valid return the value is valid. If true, it is not null value.
func (b *Bool) Valid() bool {
	return b.b.Valid
}

// BoolValue return the value.
func (b *Bool) BoolValue() bool {
	return b.b.Bool
}

// Reset set nil to the value.
func (b *Bool) Reset() {
	b.b.Bool = false
	b.b.Valid = false
}

// Set set the value.
func (b *Bool) Set(value bool) *Bool {
	b.b.Valid = true
	b.b.Bool = value
	return b
}

// Scan is a method for database/sql.
func (b *Bool) Scan(value interface{}) error {
	return b.b.Scan(value)
}

// String return string indicated the value.
func (b Bool) String() string {
	if !b.b.Valid {
		return ""
	}
	if b.b.Bool {
		return "true"
	}
	return "false"
}

// MarshalJSON encode the value to JSON.
func (b Bool) MarshalJSON() ([]byte, error) {
	if !b.b.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(b.b.Bool)
}

// UnmarshalJSON decode data to the value.
func (b *Bool) UnmarshalJSON(data []byte) error {
	var value *bool
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	b.b.Valid = value != nil
	if value == nil {
		b.b.Bool = false
	} else {
		b.b.Bool = true
	}
	return nil
}

// Value implement driver.Valuer.
func (b Bool) Value() (driver.Value, error) {
	if !b.Valid() {
		return nil, nil
	}
	return b.b.Bool, nil
}
