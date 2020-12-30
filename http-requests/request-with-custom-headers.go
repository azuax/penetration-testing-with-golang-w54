package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "http://localhost:9090/test.php", nil)
	if err != nil {
		log.Fatalln(err)
	}
	var client http.Client = http.Client{}
	req.Header.Set("User-Agent", "Go get it!")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", body)
}
