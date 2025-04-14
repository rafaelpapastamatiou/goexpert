package main

import (
	"encoding/json"

	"github.com/valyala/fastjson"
)

type User struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	City    string   `json:"city"`
	Friends []string `json:"friends"`
}

func main() {
	var p fastjson.Parser

	jsonData := `{ "user": { "name": "John Doe", "age": 30, "city": "New York", "friends": ["Alice", "Bob"] } }`

	v, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	userJson := v.GetObject("user")

	// ! Access the userJson object fields directly
	// name := userJson.Get("name").GetStringBytes()
	// age := userJson.Get("age").GetInt()
	// city := userJson.Get("city").GetStringBytes()
	// friends := userJson.Get("friends").GetArray()

	// println("Name:", string(name))
	// println("Age:", age)
	// println("City:", string(city))

	// println("Friends: ")
	// for i := 0; i < len(friends); i++ {
	// 	friend := friends[i].GetStringBytes()
	// 	println(" -", string(friend))
	// }

	// ! Convert the userJson object to a User struct
	// ! This is the correct way to unmarshal JSON into a struct
	var user User

	if err := json.Unmarshal([]byte(userJson.String()), &user); err != nil {
		panic(err)
	}

	println("Name:", user.Name)
	println("Age:", user.Age)
	println("City:", user.City)

	println("Friends: ")
	for _, friend := range user.Friends {
		println(" -", friend)
	}
}
