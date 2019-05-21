package nullable

import "testing"

func TestBoolFrom(t *testing.T) {
	b := BoolFrom(true)
	if !b.valid || !b.value {
		t.Errorf("Bool was not valid")
	}
}

func TestFloat32From(t *testing.T) {
	f := Float32From(34.546)
	if !f.valid || f.value != 34.546 {
		t.Errorf("Float32 was not valid")
	}
}

func TestFloat64From(t *testing.T) {
	f := Float64From(6542.9859849584053)
	if !f.valid || f.value != 6542.9859849584053 {
		t.Errorf("Float64 was not valid")
	}
}

func TestIntFrom(t *testing.T) {
	i := IntFrom(8823828)
	if !i.IsValid() || i.value != 8823828 {
		t.Errorf("Int was not valid")
	}
}

func TestInt32From(t *testing.T) {
	i := Int32From(984983)
	if !i.IsValid() || i.value != 984983 {
		t.Errorf("Int32 was not valid")
	}
}

func TestInt64From(t *testing.T) {
	i := Int64From(843010193443)
	if !i.IsValid() || i.value != 843010193443 {
		t.Errorf("Int64 was not valid")
	}
}

func TestNilBool(t *testing.T) {
	b := NilBool()
	if b.valid {
		t.Errorf("Bool was not nil")
	}
}

func TestNilFloat32(t *testing.T) {
	f := NilFloat32()
	if f.IsValid() {
		t.Errorf("Float32 was not nil")
	}
}

func TestNilFloat64(t *testing.T) {
	f := NilFloat64()
	if f.IsValid() {
		t.Errorf("Float64 was not nil")
	}
}

func TestNilInt(t *testing.T) {
	i := NilInt()
	if i.IsValid() {
		t.Errorf("Int was not nil")
	}
}

func TestNilInt32(t *testing.T) {
	i := NilInt32()
	if i.IsValid() {
		t.Errorf("Int32 was not nil")
	}
}

func TestNilInt64(t *testing.T) {
	i := NilInt64()
	if i.IsValid() {
		t.Errorf("Int64 was not nil")
	}
}

func TestNilString(t *testing.T) {
	s := NilString()
	if s.IsValid() {
		t.Errorf("String was not valid")
	}
}

func TestStringFrom(t *testing.T) {
	s := StringFrom("Test str")
	if !s.IsValid() || s.value != "Test str" {
		t.Errorf("String was not valid")
	}
}
