package main

import (
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

var (
	SERVER = "scanme.nmap.org"
	RANGE  = 1024
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i < RANGE; i++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			req, err := net.DialTimeout("tcp", SERVER+":"+strconv.Itoa(p), time.Second*2)
			if err != nil {
				return
			}
			log.Printf("Port %d is open\n", p)
			buf := make([]byte, 4096)
			req.SetReadDeadline(time.Now().Add(time.Second * 4))
			nRead, err := req.Read(buf)
			if err != nil {
				return
			}
			log.Printf("Banner: %s\n", buf[:nRead])
		}(i)
	}
	wg.Wait()
}
