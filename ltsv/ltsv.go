package ltsv

import (
	"encoding/json"
	// "fmt"
	"reflect"
	"strconv"
	"strings"
)

const (
	ltsvTag = "ltsv"
)

func Marshal(v interface{}) string {
	handler := marshaler{}
	rv := reflect.ValueOf(v)
	rt := rv.Type()

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		key := field.Tag.Get(ltsvTag)
		value := rv.Field(i)
		converted := handler.convertValue(field.Type, value)

		handler.pairs.append(key, converted)
	}

	return handler.pairs.join()
}

func Unmarshal(log string, v interface{}) error {
	// TODO:
	return nil
}

type marshaler struct {
	pairs pairArray
}

func (m *marshaler) convertValue(t reflect.Type, v reflect.Value) string {
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
			// NOTE:
			// not support type is json serialize.
			bytes, err := json.Marshal(v.Interface())
			if err == nil {
				converted = string(bytes)
			}
		}
	}
	return converted
}

type unmarshaler struct {
	// TODO: implement
}

type pairArray struct {
	pairs []pair
}

func (pa *pairArray) len() int {
	return len(pa.pairs)
}

func (pa *pairArray) append(key string, value string) {
	p := pair{
		key:   key,
		value: value,
	}
	pa.pairs = append(pa.pairs, p)
}

func (pa *pairArray) join() string {
	var ret []string
	for _, p := range pa.pairs {
		ret = append(ret, p.join())
	}
	return strings.Join(ret, "\t")
}

type pair struct {
	key   string
	value string
}

func (p *pair) join() string {
	return p.key + ":" + p.value
}

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
