# Go Nullable
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fhassieswift621%2Fgo-nullable.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fhassieswift621%2Fgo-nullable?ref=badge_shield)

Nullable types for Go with JSON support where a JSON value may be optional and nullable.

## Nullable Types
- Bool
- Float32
- Float64
- Int
- Int32
- Int64
- String

## Tutorial
**Creating nullable value types with values**
```go
b := BoolFrom(true)
f32 := Float32From(0.445)
f64 := Float64From(453.392392)
i := IntFrom(4343521)
i32 := Int32From(1324322)
i64 := Int64From(7685345432)
s := StringFrom("str")
```

**Creating nullable value types with nil values**
```go
b := NilBool()
f32 := NilFloat32()
f64 := NilFloat64()
i := NilInt()
i32 := NilInt32()
i64 := NilInt64()
s := NilString()
```

**Using nullable types**
```go
// Checking whether the value is nil.
f32 := Float32From(345.232)
if f32.IsValid() {
	// Value is not nil.
}
```

```go
// Getting the value from the nullable type or the zero value
// for the type if it is nil.
i := NilInt()
v := i.Value() // This returns 0 because the int is nil.
```

```go
// Specifying a default value if the value is nil.
s := NilString()
// This returns "default str" because the value is nil.
v := s.ValueOrDefault("default str")
```

## JSON Tutorial
JSON support has been added which allows these types to be marshaled and unmarshaled.

**Ex 1: Nullable string - Username**
```go
type ExampleJsonStruct struct {
	UserID      string              `json:"user_id"`
	Username    nullable.String     `json:"username"`
}

var parsed ExampleJsonStruct
json.Unmarshal([]byte, &ExampleJsonStruct)

// Check if Username has a non nil value.
if parsed.Username.IsValid() {
	// Value is non nil.
	value := parsed.Username.Value()
} else {
	// Value is nil.
}
```

**Ex 2: Optional and nullable int - MemberCount**
```go
type ExampleJsonStruct struct {
	GuildID         string              `json:"guild_id"`
	MemberCount     *nullable.Int       `json:"member_count"`
}

var parsed ExampleJsonStruct
json.Unmarshal([]byte, &ExampleJsonStruct)

// Check if the member count was provided in the JSON response.
if parsed.MemberCount != nil {
	// Value was provided in the JSON.
	if parsed.MemberCount.IsValid() {
		// Value was not null in JSON response.
	} else {
		// Value was null in JSON response.
	}
}
```

## Todo
- Tests
- Add support for converting between types for JSON unmarshalling, i.e. string to nullable int.

## License
Copyright &copy;2019 Hassie.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

<http://www.apache.org/licenses/LICENSE-2.0>

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.


[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fhassieswift621%2Fgo-nullable.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fhassieswift621%2Fgo-nullable?ref=badge_large)