package main

import "fmt"

type Client struct {
	Name string
}

func (c Client) ChangeNameWithoutPointer(name string) {
	c.Name = name
}

func (c *Client) ChangeNameUsingPointer(name string) {
	c.Name = name
}

func NewClient(name string) *Client {
	return &Client{Name: name}
}

func main() {
	john := Client{
		Name: "John",
	}

	fmt.Println(john.Name) // prints John

	john.ChangeNameWithoutPointer("David") // can't change John's name because it creates a new copy

	fmt.Println(john.Name) // prints John

	john.ChangeNameUsingPointer("David") // changes John's name because it uses a pointer to John's data in memory

	fmt.Println(john.Name) // prints David

	arthur := NewClient("Arthur") // returns a pointer to a Client

	fmt.Println(arthur)
}
