package nullable

import (
	"encoding/json"
)

// Bool is a nullable bool.
type Bool struct {
	valid bool
	value bool
}

// IsValid returns whether the bool is not null, i.e. has a non nil value.
func (b Bool) IsValid() bool {
	return b.valid
}

// Value is an alias for ValueOrZero.
// See ValueOrZero for more info.
func (b Bool) Value() bool {
	return b.ValueOrZero()
}

// Value attempts to get the value of the bool.
// If the bool is nil, this will return false.
func (b Bool) ValueOrZero() bool {
	// Check if the bool is valid.
	if !b.IsValid() {
		return false
	}

	return b.value
}

// ValueOrDefault attempts to get the value of the bool.
// If the value is nil, the provided default value is returned.
func (b Bool) ValueOrDefault(def bool) bool {
	// Check if the bool has a non nil value.
	if !b.IsValid() {
		return def
	}

	return b.value
}

// MarshalJSON encodes the value to JSON.
func (b Bool) MarshalJSON() ([]byte, error) {
	// Check if the bool is valid.
	if !b.IsValid() {
		return []byte("null"), nil
	}

	return json.Marshal(b.value)
}

// UnmarshalJSON decodes the JSON data to the value.
func (b *Bool) UnmarshalJSON(data []byte) error {
	var value *bool

	// Attempt to unmarshal JSON.
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	// Update value.
	if value == nil {
		b.valid = false
	} else {
		b.valid = true
		b.value = *value
	}

	return nil
}
