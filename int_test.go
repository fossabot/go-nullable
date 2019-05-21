package nullable

import (
	"testing"
)

func TestInt_IsValid(t *testing.T) {
	v := 65665
	i := IntFrom(v)
	if !i.IsValid() {
		t.Errorf("Int was nil for %d", v)
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

func TestInt_ValueOrZero(t *testing.T) {
	v := 4385
	i := IntFrom(v)
	actual := i.ValueOrZero()
	if actual != v {
		t.Errorf("Int was not valid, expected %d, got %d", v, actual)
	}
}
