package to

import "time"

// Bool convert an interface to a bool type.
func Bool(i interface{}) bool {
	v, _ := BoolE(i)
	return v
}

// Time convert an interface to a time.Time type.
func Time(i interface{}) time.Time {
	v, _ := TimeE(i)
	return v
}

// Duration convert an interface to a time.Duration type.
func Duration(i interface{}) time.Duration {
	v, _ := DurationE(i)
	return v
}

// Float64 convert an interface to a float64 type.
func Float64(i interface{}) float64 {
	v, _ := Float64E(i)
	return v
}

// Float32 convert an interface to a float32 type.
func Float32(i interface{}) float32 {
	v, _ := Float32E(i)
	return v
}

// Int64 convert an interface to an int64 type.
func Int64(i interface{}) int64 {
	v, _ := Int64E(i)
	return v
}

// Int32 convert an interface to an int32 type.
func Int32(i interface{}) int32 {
	v, _ := Int32E(i)
	return v
}

// Int16 convert an interface to an int16 type.
func Int16(i interface{}) int16 {
	v, _ := Int16E(i)
	return v
}

// Int8 convert an interface to an int8 type.
func Int8(i interface{}) int8 {
	v, _ := Int8E(i)
	return v
}

// Int convert an interface to an int type.
func Int(i interface{}) int {
	v, _ := IntE(i)
	return v
}

// Uint convert an interface to a uint type.
func Uint(i interface{}) uint {
	v, _ := UintE(i)
	return v
}

// Uint64 convert an interface to a uint64 type.
func Uint64(i interface{}) uint64 {
	v, _ := Uint64E(i)
	return v
}

// Uint32 convert an interface to a uint32 type.
func Uint32(i interface{}) uint32 {
	v, _ := Uint32E(i)
	return v
}

// Uint16 convert an interface to a uint16 type.
func Uint16(i interface{}) uint16 {
	v, _ := Uint16E(i)
	return v
}

// Uint8 convert an interface to a uint8 type.
func Uint8(i interface{}) uint8 {
	v, _ := Uint8E(i)
	return v
}

// String convert an interface to a string type.
func String(i interface{}) string {
	v, _ := StringE(i)
	return v
}

// StringMapString convert an interface to a map[string]string type.
func StringMapString(i interface{}) map[string]string {
	v, _ := StringMapStringE(i)
	return v
}

// StringMapStringSlice convert an interface to a map[string][]string type.
func StringMapStringSlice(i interface{}) map[string][]string {
	v, _ := StringMapStringSliceE(i)
	return v
}

// StringMapBool convert an interface to a map[string]bool type.
func StringMapBool(i interface{}) map[string]bool {
	v, _ := StringMapBoolE(i)
	return v
}

// StringMapInt convert an interface to a map[string]int type.
func StringMapInt(i interface{}) map[string]int {
	v, _ := StringMapIntE(i)
	return v
}

// StringMapInt64 convert an interface to a map[string]int64 type.
func StringMapInt64(i interface{}) map[string]int64 {
	v, _ := StringMapInt64E(i)
	return v
}

// StringMap convert an interface to a map[string]interface{} type.
func StringMap(i interface{}) map[string]interface{} {
	v, _ := StringMapE(i)
	return v
}

// Slice convert an interface to a []interface{} type.
func Slice(i interface{}) []interface{} {
	v, _ := SliceE(i)
	return v
}

// BoolSlice convert an interface to a []bool type.
func BoolSlice(i interface{}) []bool {
	v, _ := BoolSliceE(i)
	return v
}

// StringSlice convert an interface to a []string type.
func StringSlice(i interface{}) []string {
	v, _ := StringSliceE(i)
	return v
}

// IntSlice convert an interface to a []int type.
func IntSlice(i interface{}) []int {
	v, _ := IntSliceE(i)
	return v
}

// DurationSlice convert an interface to a []time.Duration type.
func DurationSlice(i interface{}) []time.Duration {
	v, _ := DurationSliceE(i)
	return v
}
