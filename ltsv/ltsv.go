package ltsv

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const (
	ltsvTag = "ltsv"

	tabKey   = "\t"
	colonKey = ":"
)

// Marshal interface convert to LTSV format string.
func Marshal(v interface{}) string {
	handler := marshaler{}
	rv := reflect.ValueOf(v)
	rt := rv.Type()

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		key := field.Tag.Get(ltsvTag)
		value := rv.Field(i)
		converted := handler.convertString(field.Type, value)

		handler.pairs.append(key, converted)
	}

	return handler.pairs.join()
}

// Unmarshal todo...
func Unmarshal(log string, v interface{}) error {
	handler := unmarshaler{}
	err := handler.pairs.parse(log)
	if err != nil {
		return err
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("target interface is not pointer or nil")
	}

	rv = reflect.ValueOf(v).Elem()
	rt := rv.Type()

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		key := field.Tag.Get(ltsvTag)
		value := handler.pairs.get(key)
		if value == "" {
			continue
		}

		handler.setAndConvertValue(field.Type, &rv, i, value)
	}

	return nil
}

// marshaler handle `Marhsal`.
type marshaler struct {
	pairs pairArray
}

// convertValue target struct value convert to `string`.
func (m *marshaler) convertString(t reflect.Type, v reflect.Value) string {
	converted := ""
	if isEmptyValue(v) {
		return converted
	}

	if v.IsValid() {
		switch t.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			converted = strconv.FormatInt(v.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			converted = strconv.FormatUint(v.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			converted = strconv.FormatFloat(v.Float(), 'f', -1, 64)
		case reflect.Bool:
			converted = strconv.FormatBool(v.Bool())
		case reflect.String:
			converted = v.String()
		default:
			// not support type is json serialize.
			bytes, err := json.Marshal(v.Interface())
			if err == nil {
				converted = string(bytes)
			}
		}
	}
	return converted
}

// unmarshaler handle `Unmarshal`.
type unmarshaler struct {
	pairs pairArray
}

// setAndConvertValue target value convert to each type and set target struct.
func (u *unmarshaler) setAndConvertValue(t reflect.Type, v *reflect.Value, i int, value string) {
	if !v.IsValid() || !v.Field(i).CanSet() {
		return
	}

	switch t.Kind() {
	case reflect.Int:
		converted, err := strconv.ParseInt(value, 0, 0)
		if err != nil {
			return
		}
		v.Field(i).SetInt(converted)
	case reflect.Int8:
		converted, err := strconv.ParseInt(value, 0, 8)
		if err != nil {
			return
		}
		v.Field(i).SetInt(converted)
	case reflect.Int16:
		converted, err := strconv.ParseInt(value, 0, 16)
		if err != nil {
			return
		}
		v.Field(i).SetInt(converted)
	case reflect.Int32:
		converted, err := strconv.ParseInt(value, 0, 32)
		if err != nil {
			return
		}
		v.Field(i).SetInt(converted)
	case reflect.Int64:
		converted, err := strconv.ParseInt(value, 0, 64)
		if err != nil {
			return
		}
		v.Field(i).SetInt(converted)
	case reflect.Uint:
		converted, err := strconv.ParseUint(value, 0, 0)
		if err != nil {
			return
		}
		v.Field(i).SetUint(converted)
	case reflect.Uint8:
		converted, err := strconv.ParseUint(value, 0, 8)
		if err != nil {
			return
		}
		v.Field(i).SetUint(converted)
	case reflect.Uint16:
		converted, err := strconv.ParseUint(value, 0, 16)
		if err != nil {
			return
		}
		v.Field(i).SetUint(converted)
	case reflect.Uint32:
		converted, err := strconv.ParseUint(value, 0, 32)
		if err != nil {
			return
		}
		v.Field(i).SetUint(converted)
	case reflect.Uint64:
		converted, err := strconv.ParseUint(value, 0, 64)
		if err != nil {
			return
		}
		v.Field(i).SetUint(converted)
	case reflect.Float32:
		converted, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return
		}
		v.Field(i).SetFloat(converted)
	case reflect.Float64:
		converted, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		v.Field(i).SetFloat(converted)
	case reflect.Bool:
		converted, err := strconv.ParseBool(value)
		if err != nil {
			return
		}
		v.Field(i).SetBool(converted)
	case reflect.String:
		v.Field(i).SetString(value)
	case reflect.Struct, reflect.Slice, reflect.Array, reflect.Map, reflect.Interface:
		s := reflect.New(t)

		bytes := []byte(value)
		err := json.Unmarshal(bytes, s.Interface())
		if err != nil {
			fmt.Println("error: ", err)
			return
		}

		v.Field(i).Set(s.Elem())
	default:
		return
	}
}

// pairArray control `pair` slice.
type pairArray struct {
	pairs []pair
}

// len return `pairs` length.
func (pa *pairArray) len() int {
	return len(pa.pairs)
}

// append add `pair` struct to `pair` field from key and value.
func (pa *pairArray) append(key string, value string) {
	p := pair{
		key:   key,
		value: value,
	}
	pa.pairs = append(pa.pairs, p)
}

// parse this method parse LTSV format to `pairArray` struct.
func (pa *pairArray) parse(ltsv string) error {
	if len(ltsv) < 1 {
		return errors.New("ltsv log is empty")
	}

	keyvalues := strings.Split(ltsv, "\t")
	if len(keyvalues) < 1 {
		return errors.New("ltsv log is empty")
	}

	for _, kv := range keyvalues {
		kandv := strings.SplitN(kv, colonKey, 2)
		if len(kandv) < 2 {
			// TODO: log
			continue
		}
		k, v := kandv[0], kandv[1]
		pa.append(k, v)
	}

	return nil
}

// get return value from `pair`.key match.
func (pa *pairArray) get(key string) string {
	for _, p := range pa.pairs {
		if p.key == key {
			return p.value
		}
	}
	return ""
}

// join keyvalue slice joined by tab.
// output "key:value\tkeyvalue\t..."
func (pa *pairArray) join() string {
	var ret []string
	for _, p := range pa.pairs {
		ret = append(ret, p.join())
	}
	return strings.Join(ret, tabKey)
}

// pair this struct have key and value.
type pair struct {
	key   string
	value string
}

// join key and value joined by colon
// output "key:value"
func (p *pair) join() string {
	return p.key + colonKey + p.value
}

// util

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	//case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	//return v.Int() == 0
	//case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
	//return v.Uint() == 0
	//case reflect.Float32, reflect.Float64:
	//return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
