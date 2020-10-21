package helpers

type (
	// String is an optional types of string
	String interface{}
	// Int32 is an optional types of int32
	Int32 interface{}
	// Int64 is an optional types of int64
	Int64 interface{}
	// Float64 is an optional types of float64
	Float64 interface{}
	// Bool is an optional types of bool
	Bool interface{}
)

// ToBool parse the optional Bool to the bool type
func ToBool(value interface{}) (bool, bool) {
	v, ok := value.(bool)
	return v, ok
}

// ToFloat64 parse the optional Float64 to the float64 type
func ToFloat64(value interface{}) (float64, bool) {
	v, ok := value.(float64)
	return v, ok
}

// ToInt32 parse the optional Int32 to the int32 type
func ToInt(value interface{}) (int, bool) {
	v, ok := value.(int)
	return v, ok
}

// ToInt64 parse the optional Int64 to the int64 type
func ToInt64(value interface{}) (int64, bool) {
	v, ok := value.(int64)
	return v, ok
}

// ToString parse the optional String to the string type
func ToString(value interface{}) (string, bool) {
	v, ok := value.(string)
	return v, ok
}
