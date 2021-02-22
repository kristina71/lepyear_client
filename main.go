package main

import (
	"bytes"
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

	resp, err := http.Post("http://localhost:3000/leapYear", "application/json", bytes.NewBufferString(fmt.Sprintf("{\"year\":%d}",year)))
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
