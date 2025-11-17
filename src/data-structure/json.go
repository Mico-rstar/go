package main

import (
	"encoding/json"
	"fmt"
)

type Action struct {
	Run		bool
	Swim	bool
	Fly		bool 
}

type Animal struct {
	Sex		uint8
	Name	string
	Age 	int
	Action	Action
}

func test() {
	bird := Animal{
		Sex: 0,
		Name: "bird",
		Age: 2,
		Action: Action{
			Run: false,
			Swim: false,
			Fly: true,
		},
	}
	// object := []Animal{bird}
	// data, err := json.Marshal(bird)
	data, err := json.MarshalIndent(bird, "", "	")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}