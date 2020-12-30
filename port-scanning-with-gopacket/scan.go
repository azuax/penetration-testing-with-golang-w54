package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	iface    = "any"
	ip       = "127.0.0.1"
	ports    = strings.Split("21,22,25,80,135,443,631,1723", ",")
	bpf      = "src host %s and (tcp[13] == 0x11 or tcp[13] == 0x10 or tcp[13] == 0x18)"
	snaplen  = int32(1024)
	timeout  = pcap.BlockForever
	promisc  = false
	response = make(map[string]int)
)

func main() {
	bpf = fmt.Sprintf(bpf, ip)

	go func() {
		handler, err := pcap.OpenLive(iface, snaplen, promisc, timeout)
		if err != nil {
			log.Fatalln(err)
		}
		defer handler.Close()
		if err := handler.SetBPFFilter(bpf); err != nil {
			log.Fatalln(err)
		}
		for packet := range gopacket.NewPacketSource(handler, handler.LinkType()).Packets() {
			srcPort := packet.TransportLayer().TransportFlow().Src().String()
			response[srcPort]++
			// log.Println(string(packet.Dump()))
		}

	}()
	time.Sleep(time.Second * 1)
	WG := new(sync.WaitGroup)
	for _, port := range ports {
		WG.Add(1)
		go func(i string, wg *sync.WaitGroup) {
			defer wg.Done()
			req, err := net.DialTimeout("tcp", ip+":"+i, time.Second*2)
			if err != nil {
				return
			}
			req.Close()
		}(port, WG)
	}
	WG.Wait()
	// time.Sleep(time.Second * 7)
	for port, confidence := range response {
		fmt.Printf("Port %s is open, confidence: %d\n", port, confidence)
	}
}
