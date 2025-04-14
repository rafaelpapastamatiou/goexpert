package main

import (
	"github.com/valyala/fastjson"
)

func main() {
	var p fastjson.Parser

	jsonData := `{ "name": "John Doe", "age": 30, "city": "New York", "friends": ["Alice", "Bob"] }`

	v, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	name := v.GetStringBytes("name")
	age := v.GetInt("age")
	city := v.GetStringBytes("city")
	friends := v.GetArray("friends")

	println("Name:", string(name))
	println("Age:", age)
	println("City:", string(city))

	println("Friends:")
	for _, friend := range friends {
		friendName := friend.GetStringBytes()
		println(" -", string(friendName))
	}
}
