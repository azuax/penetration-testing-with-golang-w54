package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/google/gopacket"

	"github.com/google/gopacket/pcap"
)

var fileName = flag.String("f", "capture.pcap", "Filename to open (Default: capture.pcap)")

func main() {
	handler, err := pcap.OpenOffline(*fileName)
	if err != nil {
		log.Fatalln(err)
	}

	for packet := range gopacket.NewPacketSource(handler, handler.LinkType()).Packets() {
		fmt.Println(string(packet.Dump()))
	}

}
