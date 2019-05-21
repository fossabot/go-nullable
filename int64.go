package nullable

import (
	"encoding/json"
)

// Int64 is a nullable int64.
type Int64 struct {
	valid bool
	value int64
}

// IsValid returns whether the int is not null, i.e. has a non nil value.
func (i Int64) IsValid() bool {
	return i.valid
}

// Value attempts to get the value of the int.
// If the int is nil, this will return 0.
func (i Int64) Value() int64 {
	// Check if the int is valid.
	if !i.IsValid() {
		return 0
	}

	return i.value
}

// ValueOrDefault attempts to get the value of the int.
// If the value is nil, the provided default value is returned.
func (i Int64) ValueOrDefault(def int64) int64 {
	// Check if the int has a non nil value.
	if !i.IsValid() {
		return def
	}

	return i.value
}

// MarshalJSON encodes the value to JSON.
func (i Int64) MarshalJSON() ([]byte, error) {
	// Check if the int is valid.
	if !i.IsValid() {
		return []byte("null"), nil
	}

	return json.Marshal(i.value)
}

// UnmarshalJSON decodes the JSON data to the value.
func (i *Int64) UnmarshalJSON(data []byte) error {
	var value *int64

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
