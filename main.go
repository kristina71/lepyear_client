package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Print("Enter an year:")
	var year uint64
	_, err := fmt.Scan(&year)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Get(fmt.Sprintf("http://localhost:3000/leapYear?year=%d", year))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	println(string(body))
}
