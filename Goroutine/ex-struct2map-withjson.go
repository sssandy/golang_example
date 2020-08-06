package main

import (
	"fmt"
	"reflect"
)

func Struct2Map(obj interface{}) map[string]interface{} {
	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)

	var data = make(map[string]interface{})

	for i := 0; i < objType.NumField(); i++ {
		data[objType.Field(i).Tag.Get("json")] = objValue.Field(i).Interface()
	}

	return data
}

type Ssd struct {
	Hello string `json:"hello_body"`
}

func main() {
	s := Ssd{
		"hh",
	}

	sm := Struct2Map(s)

	fmt.Println(sm)
}
