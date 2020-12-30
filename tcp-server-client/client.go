package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	req, err := net.Dial("tcp", "localhost:2020")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(req, "Hi there from the client")
	buf := make([]byte, 4096)
	read, err := req.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Received data: \n%s\n", buf[:read])

}
