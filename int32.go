package nullable

import (
	"encoding/json"
)

// Int32 is a nullable int32.
type Int32 struct {
	valid bool
	value int32
}

// IsValid returns whether the int is not null, i.e. has a non nil value.
func (i Int32) IsValid() bool {
	return i.valid
}

// Value attempts to get the value of the int.
// If the int is nil, this will return 0.
func (i Int32) Value() int32 {
	// Check if the int is valid.
	if !i.IsValid() {
		return 0
	}

	return i.value
}

// ValueOrDefault attempts to get the value of the int.
// If the value is nil, the provided default value is returned.
func (i Int32) ValueOrDefault(def int32) int32 {
	// Check if the int has a non nil value.
	if !i.IsValid() {
		return def
	}

	return i.value
}

// MarshalJSON encodes the value to JSON.
func (i Int32) MarshalJSON() ([]byte, error) {
	// Check if the int is valid.
	if !i.IsValid() {
		return []byte("null"), nil
	}

	return json.Marshal(i.value)
}

// UnmarshalJSON decodes the JSON data to the value.
func (i *Int32) UnmarshalJSON(data []byte) error {
	var value *int32

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
