package to

import (
	"errors"
	"fmt"
	"html/template"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var errNegativeNotAllowed = errors.New("unable to convert negative value")

// BoolE convert an interface to a bool type.
func BoolE(i interface{}) (bool, error) {
	i = indirect(i)

	switch b := i.(type) {
	case bool:
		return b, nil
	case nil:
		return false, nil
	case int:
		if i.(int) != 0 {
			return true, nil
		}
		return false, nil
	case string:
		return strconv.ParseBool(i.(string))
	default:
		return false, fmt.Errorf("unable to cast %#v of type %T to bool", i, i)
	}
}

// BoolSliceE convert an interface to a []bool type.
func BoolSliceE(i interface{}) ([]bool, error) {
	if i == nil {
		return []bool{}, fmt.Errorf("unable to cast %#v of type %T to []bool", i, i)
	}

	switch v := i.(type) {
	case []bool:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]bool, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := BoolE(s.Index(j).Interface())
			if err != nil {
				return []bool{}, fmt.Errorf("unable to cast %#v of type %T to []bool", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []bool{}, fmt.Errorf("unable to cast %#v of type %T to []bool", i, i)
	}
}

// DurationE convert an interface to a time.Duration type.
func DurationE(i interface{}) (d time.Duration, err error) {
	i = indirect(i)

	switch s := i.(type) {
	case time.Duration:
		return s, nil
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8:
		d = time.Duration(Int64(s))
		return
	case float32, float64:
		d = time.Duration(Float64(s))
		return
	case string:
		if strings.ContainsAny(s, "nsuµmh") {
			d, err = time.ParseDuration(s)
		} else {
			d, err = time.ParseDuration(s + "ns")
		}
		return
	default:
		err = fmt.Errorf("unable to cast %#v of type %T to Duration", i, i)
		return
	}
}

// DurationSliceE convert an interface to a []time.Duration type.
func DurationSliceE(i interface{}) ([]time.Duration, error) {
	if i == nil {
		return []time.Duration{}, fmt.Errorf("unable to cast %#v of type %T to []time.Duration", i, i)
	}

	switch v := i.(type) {
	case []time.Duration:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]time.Duration, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := DurationE(s.Index(j).Interface())
			if err != nil {
				return []time.Duration{}, fmt.Errorf("unable to cast %#v of type %T to []time.Duration", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []time.Duration{}, fmt.Errorf("unable to cast %#v of type %T to []time.Duration", i, i)
	}
}

// Float32E convert an interface to a float32 type.
func Float32E(i interface{}) (float32, error) {
	i = indirect(i)

	switch s := i.(type) {
	case float64:
		return float32(s), nil
	case float32:
		return s, nil
	case int:
		return float32(s), nil
	case int64:
		return float32(s), nil
	case int32:
		return float32(s), nil
	case int16:
		return float32(s), nil
	case int8:
		return float32(s), nil
	case uint:
		return float32(s), nil
	case uint64:
		return float32(s), nil
	case uint32:
		return float32(s), nil
	case uint16:
		return float32(s), nil
	case uint8:
		return float32(s), nil
	case string:
		v, err := strconv.ParseFloat(s, 32)
		if err == nil {
			return float32(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to float32", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to float32", i, i)
	}
}

// Float64E convert an interface to a float64 type.
func Float64E(i interface{}) (float64, error) {
	i = indirect(i)

	switch s := i.(type) {
	case float64:
		return s, nil
	case float32:
		return float64(s), nil
	case int:
		return float64(s), nil
	case int64:
		return float64(s), nil
	case int32:
		return float64(s), nil
	case int16:
		return float64(s), nil
	case int8:
		return float64(s), nil
	case uint:
		return float64(s), nil
	case uint64:
		return float64(s), nil
	case uint32:
		return float64(s), nil
	case uint16:
		return float64(s), nil
	case uint8:
		return float64(s), nil
	case string:
		v, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return v, nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to float64", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to float64", i, i)
	}
}

// IntE convert an interface to an int type.
func IntE(i interface{}) (int, error) {
	i = indirect(i)

	switch s := i.(type) {
	case int:
		return s, nil
	case int64:
		return int(s), nil
	case int32:
		return int(s), nil
	case int16:
		return int(s), nil
	case int8:
		return int(s), nil
	case uint:
		return int(s), nil
	case uint64:
		return int(s), nil
	case uint32:
		return int(s), nil
	case uint16:
		return int(s), nil
	case uint8:
		return int(s), nil
	case float64:
		return int(s), nil
	case float32:
		return int(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
	}
}

// Int8E convert an interface to an int8 type.
func Int8E(i interface{}) (int8, error) {
	i = indirect(i)

	switch s := i.(type) {
	case int:
		return int8(s), nil
	case int64:
		return int8(s), nil
	case int32:
		return int8(s), nil
	case int16:
		return int8(s), nil
	case int8:
		return s, nil
	case uint:
		return int8(s), nil
	case uint64:
		return int8(s), nil
	case uint32:
		return int8(s), nil
	case uint16:
		return int8(s), nil
	case uint8:
		return int8(s), nil
	case float64:
		return int8(s), nil
	case float32:
		return int8(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int8(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
	}
}

// Int16E convert an interface to an int16 type.
func Int16E(i interface{}) (int16, error) {
	i = indirect(i)

	switch s := i.(type) {
	case int:
		return int16(s), nil
	case int64:
		return int16(s), nil
	case int32:
		return int16(s), nil
	case int16:
		return s, nil
	case int8:
		return int16(s), nil
	case uint:
		return int16(s), nil
	case uint64:
		return int16(s), nil
	case uint32:
		return int16(s), nil
	case uint16:
		return int16(s), nil
	case uint8:
		return int16(s), nil
	case float64:
		return int16(s), nil
	case float32:
		return int16(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int16(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
	}
}

// Int32E convert an interface to an int32 type.
func Int32E(i interface{}) (int32, error) {
	i = indirect(i)

	switch s := i.(type) {
	case int:
		return int32(s), nil
	case int64:
		return int32(s), nil
	case int32:
		return s, nil
	case int16:
		return int32(s), nil
	case int8:
		return int32(s), nil
	case uint:
		return int32(s), nil
	case uint64:
		return int32(s), nil
	case uint32:
		return int32(s), nil
	case uint16:
		return int32(s), nil
	case uint8:
		return int32(s), nil
	case float64:
		return int32(s), nil
	case float32:
		return int32(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int32(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int32", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int32", i, i)
	}
}

// Int64E convert an interface to an int64 type.
func Int64E(i interface{}) (int64, error) {
	i = indirect(i)

	switch s := i.(type) {
	case int:
		return int64(s), nil
	case int64:
		return s, nil
	case int32:
		return int64(s), nil
	case int16:
		return int64(s), nil
	case int8:
		return int64(s), nil
	case uint:
		return int64(s), nil
	case uint64:
		return int64(s), nil
	case uint32:
		return int64(s), nil
	case uint16:
		return int64(s), nil
	case uint8:
		return int64(s), nil
	case float64:
		return int64(s), nil
	case float32:
		return int64(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return v, nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
	}
}

// IntSliceE convert an interface to a []int type.
func IntSliceE(i interface{}) ([]int, error) {
	if i == nil {
		return []int{}, fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
	}

	switch v := i.(type) {
	case []int:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]int, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := IntE(s.Index(j).Interface())
			if err != nil {
				return []int{}, fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []int{}, fmt.Errorf("unable to cast %#v of type %T to []int", i, i)
	}
}

// SliceE convert an interface to a []interface{} type.
func SliceE(i interface{}) ([]interface{}, error) {
	var s []interface{}

	switch v := i.(type) {
	case []interface{}:
		return append(s, v...), nil
	case []map[string]interface{}:
		for _, u := range v {
			s = append(s, u)
		}
		return s, nil
	default:
		return s, fmt.Errorf("unable to cast %#v of type %T to []interface{}", i, i)
	}
}

// StringE convert an interface to a string type.
func StringE(i interface{}) (string, error) {
	i = indirectStringerOrError(i)

	switch s := i.(type) {
	case string:
		return s, nil
	case bool:
		return strconv.FormatBool(s), nil
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32), nil
	case int:
		return strconv.Itoa(s), nil
	case int64:
		return strconv.FormatInt(s, 10), nil
	case int32:
		return strconv.Itoa(int(s)), nil
	case int16:
		return strconv.FormatInt(int64(s), 10), nil
	case int8:
		return strconv.FormatInt(int64(s), 10), nil
	case uint:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint64:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(s), 10), nil
	case []byte:
		return string(s), nil
	case template.HTML:
		return string(s), nil
	case template.URL:
		return string(s), nil
	case template.JS:
		return string(s), nil
	case template.CSS:
		return string(s), nil
	case template.HTMLAttr:
		return string(s), nil
	case nil:
		return "", nil
	case fmt.Stringer:
		return s.String(), nil
	case error:
		return s.Error(), nil
	default:
		return "", fmt.Errorf("unable to cast %#v of type %T to string", i, i)
	}
}

// StringMapE convert an interface to a map[string]interface{} type.
func StringMapE(i interface{}) (map[string]interface{}, error) {
	var m = map[string]interface{}{}

	switch v := i.(type) {
	case map[interface{}]interface{}:
		for k, val := range v {
			m[String(k)] = val
		}
		return m, nil
	case map[string]interface{}:
		return v, nil
	case string:
		err := jsonStringToObject(v, &m)
		return m, err
	default:
		return m, fmt.Errorf("unable to cast %#v of type %T to map[string]interface{}", i, i)
	}
}

// StringMapBoolE convert an interface to a map[string]bool type.
func StringMapBoolE(i interface{}) (map[string]bool, error) {
	var m = map[string]bool{}

	switch v := i.(type) {
	case map[interface{}]interface{}:
		for k, val := range v {
			m[String(k)] = Bool(val)
		}
		return m, nil
	case map[string]interface{}:
		for k, val := range v {
			m[String(k)] = Bool(val)
		}
		return m, nil
	case map[string]bool:
		return v, nil
	case string:
		err := jsonStringToObject(v, &m)
		return m, err
	default:
		return m, fmt.Errorf("unable to cast %#v of type %T to map[string]bool", i, i)
	}
}

// StringMapIntE convert an interface to a map[string]int{} type.
func StringMapIntE(i interface{}) (map[string]int, error) {
	var m = map[string]int{}
	if i == nil {
		return m, fmt.Errorf("unable to cast %#v of type %T to map[string]int", i, i)
	}

	switch v := i.(type) {
	case map[interface{}]interface{}:
		for k, val := range v {
			m[String(k)] = Int(val)
		}
		return m, nil
	case map[string]interface{}:
		for k, val := range v {
			m[k] = Int(val)
		}
		return m, nil
	case map[string]int:
		return v, nil
	case string:
		err := jsonStringToObject(v, &m)
		return m, err
	}

	if reflect.TypeOf(i).Kind() != reflect.Map {
		return m, fmt.Errorf("unable to cast %#v of type %T to map[string]int", i, i)
	}

	mVal := reflect.ValueOf(m)
	v := reflect.ValueOf(i)
	for _, keyVal := range v.MapKeys() {
		val, err := IntE(v.MapIndex(keyVal).Interface())
		if err != nil {
			return m, fmt.Errorf("unable to cast %#v of type %T to map[string]int", i, i)
		}
		mVal.SetMapIndex(keyVal, reflect.ValueOf(val))
	}
	return m, nil
}

// StringMapInt64E convert an interface to a map[string]int64{} type.
func StringMapInt64E(i interface{}) (map[string]int64, error) {
	var m = map[string]int64{}
	if i == nil {
		return m, fmt.Errorf("unable to cast %#v of type %T to map[string]int64", i, i)
	}

	switch v := i.(type) {
	case map[interface{}]interface{}:
		for k, val := range v {
			m[String(k)] = Int64(val)
		}
		return m, nil
	case map[string]interface{}:
		for k, val := range v {
			m[k] = Int64(val)
		}
		return m, nil
	case map[string]int64:
		return v, nil
	case string:
		err := jsonStringToObject(v, &m)
		return m, err
	}

	if reflect.TypeOf(i).Kind() != reflect.Map {
		return m, fmt.Errorf("unable to cast %#v of type %T to map[string]int64", i, i)
	}
	mVal := reflect.ValueOf(m)
	v := reflect.ValueOf(i)
	for _, keyVal := range v.MapKeys() {
		val, err := Int64E(v.MapIndex(keyVal).Interface())
		if err != nil {
			return m, fmt.Errorf("unable to cast %#v of type %T to map[string]int64", i, i)
		}
		mVal.SetMapIndex(keyVal, reflect.ValueOf(val))
	}
	return m, nil
}

// StringMapStringE convert an interface to a map[string]string type.
func StringMapStringE(i interface{}) (map[string]string, error) {
	var m = map[string]string{}

	switch v := i.(type) {
	case map[string]string:
		return v, nil
	case map[string]interface{}:
		for k, val := range v {
			m[String(k)] = String(val)
		}
		return m, nil
	case map[interface{}]string:
		for k, val := range v {
			m[String(k)] = String(val)
		}
		return m, nil
	case map[interface{}]interface{}:
		for k, val := range v {
			m[String(k)] = String(val)
		}
		return m, nil
	case string:
		err := jsonStringToObject(v, &m)
		return m, err
	default:
		return m, fmt.Errorf("unable to cast %#v of type %T to map[string]string", i, i)
	}
}

// StringMapStringSliceE convert an interface to a map[string][]string type.
func StringMapStringSliceE(i interface{}) (map[string][]string, error) {
	var m = map[string][]string{}

	switch v := i.(type) {
	case map[string][]string:
		return v, nil
	case map[string][]interface{}:
		for k, val := range v {
			m[String(k)] = StringSlice(val)
		}
		return m, nil
	case map[string]string:
		for k, val := range v {
			m[String(k)] = []string{val}
		}
	case map[string]interface{}:
		for k, val := range v {
			switch vt := val.(type) {
			case []interface{}:
				m[String(k)] = StringSlice(vt)
			case []string:
				m[String(k)] = vt
			default:
				m[String(k)] = []string{String(val)}
			}
		}
		return m, nil
	case map[interface{}][]string:
		for k, val := range v {
			m[String(k)] = StringSlice(val)
		}
		return m, nil
	case map[interface{}]string:
		for k, val := range v {
			m[String(k)] = StringSlice(val)
		}
		return m, nil
	case map[interface{}][]interface{}:
		for k, val := range v {
			m[String(k)] = StringSlice(val)
		}
		return m, nil
	case map[interface{}]interface{}:
		for k, val := range v {
			key, err := StringE(k)
			if err != nil {
				return m, fmt.Errorf("unable to cast %#v of type %T to map[string][]string", i, i)
			}
			value, err := StringSliceE(val)
			if err != nil {
				return m, fmt.Errorf("unable to cast %#v of type %T to map[string][]string", i, i)
			}
			m[key] = value
		}
	case string:
		err := jsonStringToObject(v, &m)
		return m, err
	default:
		return m, fmt.Errorf("unable to cast %#v of type %T to map[string][]string", i, i)
	}
	return m, nil
}

// StringSliceE convert an interface to a []string type.
func StringSliceE(i interface{}) ([]string, error) {
	var a []string

	switch v := i.(type) {
	case []interface{}:
		for _, u := range v {
			a = append(a, String(u))
		}
		return a, nil
	case []string:
		return v, nil
	case []int8:
		for _, u := range v {
			a = append(a, String(u))
		}
		return a, nil
	case []int:
		for _, u := range v {
			a = append(a, String(u))
		}
		return a, nil
	case []int32:
		for _, u := range v {
			a = append(a, String(u))
		}
		return a, nil
	case []int64:
		for _, u := range v {
			a = append(a, String(u))
		}
		return a, nil
	case []float32:
		for _, u := range v {
			a = append(a, String(u))
		}
		return a, nil
	case []float64:
		for _, u := range v {
			a = append(a, String(u))
		}
		return a, nil
	case string:
		return strings.Fields(v), nil
	case []error:
		for _, err := range i.([]error) {
			a = append(a, err.Error())
		}
		return a, nil
	case interface{}:
		str, err := StringE(v)
		if err != nil {
			return a, fmt.Errorf("unable to cast %#v of type %T to []string", i, i)
		}
		return []string{str}, nil
	default:
		return a, fmt.Errorf("unable to cast %#v of type %T to []string", i, i)
	}
}

// TimeE convert an interface to a time.Time type.
func TimeE(i interface{}) (tim time.Time, err error) {
	i = indirect(i)

	switch v := i.(type) {
	case time.Time:
		return v, nil
	case string:
		return StringToDate(v)
	case int:
		return time.Unix(int64(v), 0), nil
	case int64:
		return time.Unix(v, 0), nil
	case int32:
		return time.Unix(int64(v), 0), nil
	case uint:
		return time.Unix(int64(v), 0), nil
	case uint64:
		return time.Unix(int64(v), 0), nil
	case uint32:
		return time.Unix(int64(v), 0), nil
	default:
		return time.Time{}, fmt.Errorf("unable to cast %#v of type %T to Time", i, i)
	}
}

// UintE convert an interface to a uint type.
func UintE(i interface{}) (uint, error) {
	i = indirect(i)

	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 0)
		if err == nil {
			return uint(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v to uint: %s", i, err)
	case int:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(s), nil
	case int64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(s), nil
	case int32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(s), nil
	case int16:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(s), nil
	case int8:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(s), nil
	case uint:
		return s, nil
	case uint64:
		return uint(s), nil
	case uint32:
		return uint(s), nil
	case uint16:
		return uint(s), nil
	case uint8:
		return uint(s), nil
	case float64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(s), nil
	case float32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint", i, i)
	}
}

// Uint8E convert an interface to a uint type.
func Uint8E(i interface{}) (uint8, error) {
	i = indirect(i)

	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 8)
		if err == nil {
			return uint8(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v to uint8: %s", i, err)
	case int:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(s), nil
	case int64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(s), nil
	case int32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(s), nil
	case int16:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(s), nil
	case int8:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(s), nil
	case uint:
		return uint8(s), nil
	case uint64:
		return uint8(s), nil
	case uint32:
		return uint8(s), nil
	case uint16:
		return uint8(s), nil
	case uint8:
		return s, nil
	case float64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(s), nil
	case float32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint8(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint8", i, i)
	}
}

// Uint16E convert an interface to a uint16 type.
func Uint16E(i interface{}) (uint16, error) {
	i = indirect(i)

	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 16)
		if err == nil {
			return uint16(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v to uint16: %s", i, err)
	case int:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(s), nil
	case int64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(s), nil
	case int32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(s), nil
	case int16:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(s), nil
	case int8:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(s), nil
	case uint:
		return uint16(s), nil
	case uint64:
		return uint16(s), nil
	case uint32:
		return uint16(s), nil
	case uint16:
		return s, nil
	case uint8:
		return uint16(s), nil
	case float64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(s), nil
	case float32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint16(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint16", i, i)
	}
}

// Uint32E convert an interface to a uint32 type.
func Uint32E(i interface{}) (uint32, error) {
	i = indirect(i)

	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 32)
		if err == nil {
			return uint32(v), nil
		}
		return 0, fmt.Errorf("unable to cast %#v to uint32: %s", i, err)
	case int:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(s), nil
	case int64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(s), nil
	case int32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(s), nil
	case int16:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(s), nil
	case int8:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(s), nil
	case uint:
		return uint32(s), nil
	case uint64:
		return uint32(s), nil
	case uint32:
		return s, nil
	case uint16:
		return uint32(s), nil
	case uint8:
		return uint32(s), nil
	case float64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(s), nil
	case float32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint32(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
	}
}

// Uint64E convert an interface to a uint64 type.
func Uint64E(i interface{}) (uint64, error) {
	i = indirect(i)

	switch s := i.(type) {
	case string:
		v, err := strconv.ParseUint(s, 0, 64)
		if err == nil {
			return v, nil
		}
		return 0, fmt.Errorf("unable to cast %#v to uint64: %s", i, err)
	case int:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case int64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case int32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case int16:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case int8:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case uint:
		return uint64(s), nil
	case uint64:
		return s, nil
	case uint32:
		return uint64(s), nil
	case uint16:
		return uint64(s), nil
	case uint8:
		return uint64(s), nil
	case float32:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case float64:
		if s < 0 {
			return 0, errNegativeNotAllowed
		}
		return uint64(s), nil
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
	}
}
