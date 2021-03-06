package main

import (
	"fmt"
	"log"
	"reflect"
)

type user struct {
	name, email, phoneNumber string
}

type order struct {
	date         string
	sum          float64
	productsList []string
}

func main() {

	u := user{"Ivan", "ivan@example.com", "+790005552211"}
	fmt.Println(buildQuery(u))

}

func buildQuery(s interface{}) string {

	typeOf := reflect.TypeOf(s)
	if typeOf.Kind() != reflect.Struct {
		return ""
	}

	builder := SqlBuilder{}
	obj := reflect.ValueOf(s)
	for i := 0; i < obj.NumField(); i++ {
		switch obj.Field(i).Kind() {
		// case reflect.Int:
		// 	builder.Fields(typeOf.Field(i).Name).Values(obj.Field(i).Int())
		// case reflect.Float64:
		// 	builder.Fields(typeOf.Field(i).Name).Values(obj.Field(i).Float())
		case reflect.String:
			builder.Fields(typeOf.Field(i).Name).Values(obj.Field(i).String())
		// case reflect.Slice:
		default:
			fmt.Printf("Unknown type %v of value %v\n", obj.Field(i).Kind(), obj.Field(i))
		}
	}

	builder.Into("users")
	q, err := builder.Build()
	if err != nil {
		log.Println(err)
		return ""
	}

	return q

}
