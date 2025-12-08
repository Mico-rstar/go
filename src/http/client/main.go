package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Post("http://localhost:8080/ping", "", nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Printf("%s\n",body)
}
