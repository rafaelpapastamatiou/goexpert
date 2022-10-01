package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Account struct {
	Number  int
	Balance int
}

type AccountWithTags struct {
	Number      int    `json:"number"`
	Balance     int    `json:"balance"`
	HiddenField string `json:"-"`
}

func main() {
	// prints using json.Marshal (stores the json value to a variable)
	account := Account{
		Number:  1,
		Balance: 1500,
	}

	accountToJson, err := json.Marshal(account)

	if err != nil {
		panic(err)
	}

	println(string(accountToJson))

	// prints to stdout using Enconder (converts to json then it prints, without using a var)
	enconder := json.NewEncoder(os.Stdout)

	err = enconder.Encode(account)

	if err != nil {
		panic(err)
	}

	// converts json to struct then stores the value using a pointer
	var account2 Account

	account2Json := []byte(`{"Number": 2, "Balance": 22000}`)

	err = json.Unmarshal(account2Json, &account2)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", account2)

	// using custom json tags
	accountWithTags1 := AccountWithTags{
		Number:      3,
		Balance:     3000,
		HiddenField: "should not be visible",
	}

	err = enconder.Encode(accountWithTags1)

	if err != nil {
		panic(err)
	}

	var accountWithTags2 AccountWithTags

	accountWithTags2Json := []byte(`{"number": 4, "balance": 1500}`)

	err = json.Unmarshal(accountWithTags2Json, &accountWithTags2)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", accountWithTags2)
}
