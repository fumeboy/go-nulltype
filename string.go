package nulltype

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// String is null friendly type for string.
type String struct {
	s sql.NullString
}

// StringOf return String that he value is set.
func StringOf(value string) String {
	var s String
	s.Set(value)
	return s
}

func PtrStringOf(value string) *String {
	var s String
	s.Set(value)
	return &s
}

// Valid return the value is valid. If true, it is not null value.
func (s *String) Valid() bool {
	return s.s.Valid
}

// StringValue return the value.
func (s *String) StringValue() string {
	return s.s.String
}

// Reset set nil to the value.
func (s *String) Reset() {
	s.s.String = ""
	s.s.Valid = false
}

// Set set the value.
func (s *String) Set(value string) {
	s.s.Valid = true
	s.s.String = value
}

// Scan is a method for database/sql.
func (s *String) Scan(value interface{}) error {
	return s.s.Scan(value)
}

// String return string indicated the value.
func (s String) String() string {
	if !s.s.Valid {
		return ""
	}
	return s.s.String
}

// MarshalJSON encode the value to JSON.
func (s String) MarshalJSON() ([]byte, error) {
	if !s.s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.s.String)
}

// UnmarshalJSON decode data to the value.
func (s *String) UnmarshalJSON(data []byte) error {
	var value *string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	s.s.Valid = value != nil
	if value == nil {
		s.s.String = ""
	} else {
		s.s.String = *value
	}
	return nil
}

// Value implement driver.Valuer.
func (s String) Value() (driver.Value, error) {
	if !s.Valid() {
		return nil, nil
	}
	return s.s.String, nil
}
