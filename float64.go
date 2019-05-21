package nullable

import (
	"encoding/json"
)

// Float64 is a nullable float64.
type Float64 struct {
	valid bool
	value float64
}

// IsValid returns whether the float is not null, i.e. has a non nil value.
func (f Float64) IsValid() bool {
	return f.valid
}

// Value attempts to get the value of the float.
// If the float is nil, this will return 0.
func (f Float64) ValueOrZero() float64 {
	// Check if the float is valid.
	if !f.IsValid() {
		return 0
	}

	return f.value
}

// ValueOrDefault attempts to get the value of the float.
// If the value is nil, the provided default value is returned.
func (f Float64) ValueOrDefault(def float64) float64 {
	// Check if the float has a non nil value.
	if !f.IsValid() {
		return def
	}

	return f.value
}

// MarshalJSON encodes the value to JSON.
func (f Float64) MarshalJSON() ([]byte, error) {
	// Check if the float is valid.
	if !f.IsValid() {
		return []byte("null"), nil
	}

	return json.Marshal(f.value)
}

// UnmarshalJSON decodes the JSON data to the value.
func (f *Float64) UnmarshalJSON(data []byte) error {
	var value *float64

	// Attempt to unmarshal JSON.
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	// Update value.
	if value == nil {
		f.valid = false
	} else {
		f.valid = true
		f.value = *value
	}

	return nil
}
