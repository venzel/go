package src

import (
	"fmt"
	"reflect"
)

func GetAssertion() {
	k := []*interface{}{}
	var queijo interface{} = "Queijo"
	k = append(k, &queijo)

	m := map[string]interface{}{
		"name":  "Tiago",
		"price": 30.1,
	}

	v, exists := m["price"]

	if exists {
		tp := reflect.TypeOf(v).String()

		if tp == "float64" {
			println("é float")
		}

		fmt.Println(v.(float64))
	}

	for _, t := range k {
		switch (*t).(type) {
		case string:
			println("string")
		case int:
			println("int")
		}
	}

	/* Map other */

	type fnx func(name string, points int)

	mx := map[string]*interface{}{
		"name": nil,
		"age":  nil,
		"fn":   nil,
	}

	name := interface{}("Tiago")
	age := interface{}(20)
	fn := interface{}(func(name string, points int) {
		fmt.Println("fn ->", name, points)
	})

	mx["name"] = &name
	mx["age"] = &age
	mx["fn"] = &fn

	name = "Rizzo x"

	for _, v := range mx {
		switch (*v).(type) {
		case string:
			fmt.Println("string")
			fmt.Println(*v)
		case int:
			fmt.Println("int")
		case fnx:
			(*v).(fnx)("Tiago", 10)
		}
	}
}
