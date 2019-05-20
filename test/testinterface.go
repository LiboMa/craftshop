package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type User struct {
	ID     int     `json: "id"`
	Name   string  `json: "name"`
	Age    int     `json: "age"`
	Weight float64 `json: "age"`
}

func check(a ...interface{}) {

	for obj := range a {

		fmt.Println(reflect.TypeOf(a[obj]))
	}

	a[2] = 1.3
	fmt.Println(a[2])

	//dict := reflect.ValueOf(a[3])
	//fmt.Println(a[3]["key"])

	//fmt.Println(dict["name"])
	fmt.Println(a[5])

}

func handleObj(arg interface{}) {

	typ := reflect.TypeOf(arg)
	val := reflect.ValueOf(arg)
	fmt.Println(val, typ, arg)

	v := reflect.ValueOf(arg)
	fmt.Println("Number of fields", v.NumField())
	//fmt.Println("Number of fields", v.NumField()[0])
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field:%d type:%T value:%v, VTYPE: %s\n", i, v.Field(i), v.Field(i), v.Type())
	}
}

func CreateQuery(q interface{}, tname string) {

	v := reflect.ValueOf(q).Elem()
	fmt.Println("------>", v)
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		//tname := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("INSERT INTO `%s` VALUES(", tname)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.Float64:
				if i == 0 {
					query = fmt.Sprintf("%s%.2f", query, v.Field(i).Float())
				} else {
					query = fmt.Sprintf("%s, %.2f", query, v.Field(i).Float())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return

	}
	fmt.Println("unsupported type")

}

func main() {

	a := 1
	b := "1"
	test := "hello string"
	c := 1.1

	dict := make(map[string]string)

	dict["key"] = "a"
	dict["name"] = "b"
	fmt.Println(dict["key"])

	u := User{ID: 1, Name: "peter", Age: 19, Weight: 50.50}
	json_data, _ := json.Marshal(&u)
	fmt.Println("json here", string(json_data))

	//var uobj User
	uobj := User{}
	//var uobj User{}
	err := json.Unmarshal(json_data, &uobj)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json_data), uobj)
	fmt.Println(u)
	check(a, b, c, test, dict, uobj)
	handleObj(uobj)
	//CreateQuery(uobj, "products")

}
