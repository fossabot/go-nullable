package nullable

import (
	"encoding/json"
)

// Float32 is a nullable float32.
type Float32 struct {
	valid bool
	value float32
}

// IsValid returns whether the float is not null, i.e. has a non nil value.
func (f Float32) IsValid() bool {
	return f.valid
}

// Value is an alias for ValueOrZero.
// See ValueOrZero for more info.
func (f Float32) Value() float32 {
	return f.ValueOrZero()
}

// Value attempts to get the value of the float.
// If the float is nil, this will return 0.
func (f Float32) ValueOrZero() float32 {
	// Check if the float is valid.
	if !f.IsValid() {
		return 0
	}

	return f.value
}

// ValueOrDefault attempts to get the value of the float.
// If the value is nil, the provided default value is returned.
func (f Float32) ValueOrDefault(def float32) float32 {
	// Check if the float has a non nil value.
	if !f.IsValid() {
		return def
	}

	return f.value
}

// MarshalJSON encodes the value to JSON.
func (f Float32) MarshalJSON() ([]byte, error) {
	// Check if the float is valid.
	if !f.IsValid() {
		return []byte("null"), nil
	}

	return json.Marshal(f.value)
}

// UnmarshalJSON decodes the JSON data to the value.
func (f *Float32) UnmarshalJSON(data []byte) error {
	var value *float32

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
