package nullable

import (
	"encoding/json"
)

// Int is a nullable int.
type Int struct {
	valid bool
	value int
}

// IsValid returns whether the int is not null, i.e. has a non nil value.
func (i Int) IsValid() bool {
	return i.valid
}

// Value attempts to get the value of the int.
// If the int is nil, this will return 0.
func (i Int) ValueOrZero() int {
	// Check if the int is valid.
	if !i.IsValid() {
		return 0
	}

	return i.value
}

// ValueOrDefault attempts to get the value of the int.
// If the value is nil, the provided default value is returned.
func (i Int) ValueOrDefault(def int) int {
	// Check if the int has a non nil value.
	if !i.IsValid() {
		return def
	}

	return i.value
}

// MarshalJSON encodes the value to JSON.
func (i Int) MarshalJSON() ([]byte, error) {
	// Check if the int is valid.
	if !i.IsValid() {
		return []byte("null"), nil
	}

	return json.Marshal(i.value)
}

// UnmarshalJSON decodes the JSON data to the value.
func (i *Int) UnmarshalJSON(data []byte) error {
	var value *int

	// Attempt to unmarshal JSON.
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	// Update value.
	if value == nil {
		i.valid = false
	} else {
		i.valid = true
		i.value = *value
	}

	return nil
}
