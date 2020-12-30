package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strconv"
)

var (
	port     = 2020
	address  = "127.0.0.1"
	protocol = "tcp"
)

func echo(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 131072)
	for {
		read, err := conn.Read(buf)
		if err != nil {
			log.Println("Unable to read data")
			continue
		}
		log.Println("\n", string(buf[:read]))
		log.Print("> \b")
		reader := bufio.NewReader(os.Stdin)
		command, _ := reader.ReadString('\n')
		writer := bufio.NewWriter(conn)
		if _, err := writer.WriteString(command[:len(command)-1]); err != nil {
			log.Println("Unable to write data")
			continue
		}
		writer.Flush()
	}
}

func main() {
	server, err := net.Listen(protocol, address+":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalln("Can't bind port")
	}
	defer server.Close()
	log.Println("Server binded")
	counter := 0
	for {
		conn, err := server.Accept()
		counter++
		if err != nil {
			log.Println("Can't accept request")
			continue
		}
		log.Printf("\nClient-%d connected.\n", counter)
		go echo(conn)
	}
}
