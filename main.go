package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Println("Start")
	response, error := http.Get("https://kids.yahoo.co.jp/today/")
	defer response.Body.Close()
	if error != nil {
		log.Println("Access failed.")
		return
	}

	if response.StatusCode != 200 {
		log.Fatalf("Access failed. StatusCode: %d\n", response.StatusCode)
		return
	}

	log.Println("Success access.")

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	fmt.Println(string(bytes))
}
