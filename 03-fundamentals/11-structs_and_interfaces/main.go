package main

import "fmt"

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

type Person interface {
	Disable()
	ToString() string
}

type Walk interface {
	Walk()
}

type Client struct {
	Name    string
	Age     int
	Active  bool
	Address Address
}

func (c *Client) ToString() string {
	return fmt.Sprintf("Name: %s, Age: %d, Active: %v, Address: %s, %d - %s/%s\n", c.Name, c.Age, c.Active, c.Address.Street, c.Address.Number, c.Address.City, c.Address.State)
}

func (c *Client) Disable() {
	c.Active = false
	fmt.Printf("Client %s was disabled\n", c.Name)
}

func main() {
	john := Client{
		Name:   "John",
		Age:    28,
		Active: true,
		Address: Address{
			Street: "Av. Amazonas",
			Number: 2002,
			City:   "Belo Horizonte",
			State:  "MG",
		},
	}

	disablePerson(&john) // can do because Client implements the Person interface

	// makePersonWalk(&john) // can´t do because Client doesn't implement the Walk interface
}

func disablePerson(p Person) {
	println(p.ToString())

	p.Disable()

	println(p.ToString())
}

func makePersonWalk(w Walk) {
	w.Walk()
}
