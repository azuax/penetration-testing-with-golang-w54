package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
)

var (
	// PORT sever listening port
	PORT = 2020
	// ADDRESS ip/host of server
	ADDRESS = "localhost"
	// PROTOCOL which protocol uses for communication
	PROTOCOL = "tcp"
)

func echo(conn net.Conn) {
	defer conn.Close()
	var code string = `HTTP/1.1 200 OK
	Accept-Ranges: bytes
	Content-Type: text/html


	<html>
		<form method="POST">
			<input type="text" name="user">
			<input type="password" name="pass">
			<input type="submit" name="sub">
		</form>
	</html>`
	buf := make([]byte, 4096)
	reader := bufio.NewReader(conn)
	content, err := reader.Read(buf)
	if err != nil {
		log.Println("Unable to read data!")
		return
	}
	log.Printf("Received %d bytes from the client: %s\n", content, string(buf[:content]))

	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(code); err != nil {
		log.Println("Unable to write data!")
		return
	}
	writer.Flush()
}

func main() {
	server, err := net.Listen(PROTOCOL, ADDRESS+":"+strconv.Itoa(PORT))
	if err != nil {
		log.Fatalln(fmt.Sprintf("Cannot bind port %d\n", PORT))
	}
	log.Printf("Server binded on: %s:%d\n", ADDRESS, PORT)
	defer server.Close()
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Println("Cannot accept the request")
			continue
		}
		log.Println("Request received!")
		go echo(conn)
	}
}
