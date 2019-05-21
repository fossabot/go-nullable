package nullable

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
	"testing"
)

func TestInt_IsValid(t *testing.T) {
	v := 65665
	i := IntFrom(v)
	if !i.IsValid() {
		t.Errorf("Int was nil for %d", v)
	}
}

func TestInt_MarshalJSON(t *testing.T) {
	v := 83794

	// Encode into JSON.
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(IntFrom(v)); err != nil {
		t.Errorf(err.Error())
	}

	expected := strconv.FormatInt(int64(v), 10)
	actual := strings.TrimSpace(buf.String())
	if actual != expected {
		t.Errorf("Int was invalid, expected %s, got %s", expected, actual)
	}
}

func TestInt_Value(t *testing.T) {
	v := 429934
	i := IntFrom(v)
	actual := i.Value()
	if actual != v {
		t.Errorf("Int was not valid, expected %d, got %d", v, actual)
	}
}

func TestInt_ValueOrDefault(t *testing.T) {
	v := 39409394
	i := IntFrom(v)
	actual := i.ValueOrDefault(10)
	if actual != v {
		t.Errorf("Int was not valid, expected %d, got %d", v, actual)
	}
}

func TestInt_UnmarshalJSON(t *testing.T) {
	data := []byte(`{"val": 7386}`)

	type St struct {
		Val Int `json:"val"`
	}
	var decoded St
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Errorf(err.Error())
	}

	actual := decoded.Val.Value()
	if actual != 7386 {
		t.Errorf("Int was not valid, expected %d, got %d", 7386, actual)
	}
}
