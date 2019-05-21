package nullable

// BoolFrom creates a non nil bool using the provided value.
func BoolFrom(value bool) Bool {
	return Bool{true, value}
}

// Float32From creates a non nil float32 using the provided value.
func Float32From(value float32) Float32 {
	return Float32{true, value}
}

// Float64From creates a non nil float64 using the provided value.
func Float64From(value float64) Float64 {
	return Float64{true, value}
}

// IntFrom creates a non nil int using the provided value.
func IntFrom(value int) Int {
	return Int{true, value}
}

// Int32From creates a non nil int32 using the provided value.
func Int32From(value int32) Int32 {
	return Int32{true, value}
}

// Int64From creates a non nil int64 using the provided value.
func Int64From(value int64) Int64 {
	return Int64{true, value}
}

// StringFrom creates a non nil string using the provided value.
func StringFrom(value string) String {
	return String{true, value}
}

// NilBool creates a zero value nil bool.
func NilBool() Bool {
	return Bool{valid: false}
}

// NilFloat32 creates a zero value nil float32.
func NilFloat32() Float32 {
	return Float32{valid: false}
}

// NilFloat64 creates a zero value nil float64.
func NilFloat64() Float64 {
	return Float64{valid: false}
}

// NilInt creates a zero value nil int.
func NilInt() Int {
	return Int{valid: false}
}

// NilInt32 creates a zero value nil int32.
func NilInt32() Int32 {
	return Int32{valid: false}
}

// NilInt64 creates a zero value nil int64.
func NilInt64() Int64 {
	return Int64{valid: false}
}

// NilString creates a zero value nil string.
func NilString() String {
	return String{valid: false}
}
