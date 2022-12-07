package main

import (
	"encoding/json"
	"fmt"
	"time"
)
type Customer struct {
	ID int `json:id`
	Name string `json:"name"`
	DOB time.Time `json:"time"`
}

type Transaction struct {
	TID int `json:tid`
	Time time.Time `json:"time"`
	Amount int `json:AMT`
	Credit bool `json: credit`
}

type Account struct {
	AccNo int `json:"id"`
	Holder *Customer `json:"Customer"`
	Transaction []Transaction `json:"Transaction"`
}
func main() {
	var A Account
	A = Account {
		AccNo : 1,
		Holder: &Customer {
			ID: 1,
			Name: "Veera",
		},
		Transaction: []Transaction {
			{1,time.Now(),100,true},
		},
	}
	data,err := json.Marshal(A)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s",data)
}
