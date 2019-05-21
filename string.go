package nullable

import (
	"encoding/json"
)

// String is a nullable string.
type String struct {
	valid bool
	value string
}

// IsValid returns whether the string is not null, i.e. has a non nil value.
func (s String) IsValid() bool {
	return s.valid
}

// Value attempts to get the value of the string.
// If the string is nil, this will return "".
func (s String) Value() string {
	// Check if the string is valid.
	if !s.IsValid() {
		return ""
	}

	return s.value
}

// ValueOrDefault attempts to get the value of the string.
// If the value is nil, the provided default value is returned.
func (s String) ValueOrDefault(def string) string {
	// Check if the string has a non nil value.
	if !s.IsValid() {
		return def
	}

	return s.value
}

// MarshalJSON encodes the value to JSON.
func (s String) MarshalJSON() ([]byte, error) {
	// Check if the string is valid.
	if !s.IsValid() {
		return []byte("null"), nil
	}

	return json.Marshal(s.value)
}

// UnmarshalJSON decodes the JSON data to the value.
func (s *String) UnmarshalJSON(data []byte) error {
	var value *string

	// Attempt to unmarshal JSON.
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	// Update value.
	if value == nil {
		s.valid = false
	} else {
		s.valid = true
		s.value = *value
	}

	return nil
}
