package nullable

import (
	"encoding/json"
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
	var parsed St
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Errorf(err.Error())
	}

	actual := parsed.Val.Value()
	if actual != 7386 {
		t.Errorf("Int was not valid, expected %d, got %d", 7386, actual)
	}
}
