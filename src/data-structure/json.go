package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Action struct {
	Run		bool	`json:"run"`
	Swim	bool	`json:"swim"`
	Fly		bool 	`json:"fly"`
}

type Animal struct {
	Sex		uint8	`json:"sex"`
	Name	string	`json:"name"`
	Age 	int		`json:"age"`
	Action	Action	`json:"action"`
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
	// 序列化
	data, err := json.MarshalIndent(bird, "", "	")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	// 反序列化
	var decodeAnimal Animal
	err = json.Unmarshal(data, &decodeAnimal)
	if err != nil {
		log.Fatal("反序列化失败：", err)
	}
	fmt.Printf("%#v\n", decodeAnimal)
}