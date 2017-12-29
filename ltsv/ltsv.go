package ltsv

import (
	"encoding/json"
	//"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Marshal(v interface{}) (string, error) {
	rv := reflect.ValueOf(v)
	rt := rv.Type()

	var pairs []string
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		key := field.Tag.Get("ltsv")
		value := rv.Field(i)
		converted := ""

		if value.IsValid() {
			switch field.Type.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				converted = strconv.FormatInt(value.Int(), 10)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				converted = strconv.FormatUint(value.Uint(), 10)
			case reflect.Float32, reflect.Float64:
				converted = strconv.FormatFloat(value.Float(), 'f', -1, 64)
			case reflect.Bool:
				converted = strconv.FormatBool(value.Bool())
			case reflect.String:
				converted = value.String()
			default:
				// NOTE:
				// not support type is json serialize.
				bytes, err := json.Marshal(value.Interface())
				if err == nil {
					converted = string(bytes)
				}
			}
		}

		pair := key + ":" + converted
		pairs = append(pairs, pair)
	}

	return strings.Join(pairs, "\t"), nil
}

func Unmarshal(log string, v interface{}) error {
	// TODO:
	return nil
}

type marshaler struct{}

type unmarshaler struct{}
