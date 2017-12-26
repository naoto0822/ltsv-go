package ltsv

import (
	"fmt"
	"reflect"
)

func Marshal(v interface{}) (string, error) {
	rv := reflect.ValueOf(v)
	rt := rv.Type()

	//var pairs []string
	for i := 0; i < rt.NumField(); i++ {
		key := rt.Field(i).Tag.Get("ltsv")
		value := rv.Field(i).Interface()

		fmt.Println(key)
		fmt.Println(value)
		fmt.Println("---")
	}

	return "", nil
}

func Unmarshal(log string, v interface{}) error {
	// TODO:
	return nil
}
