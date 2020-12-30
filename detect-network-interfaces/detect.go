package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

func main() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalln(err)
	}
	for _, device := range devices {
		fmt.Printf("Device name: %s\n", device.Description)
		for _, address := range device.Addresses {
			fmt.Printf("\t\tIP: %s", address.IP)
			fmt.Printf("\t\tNetmask: %s", address.Netmask)
			fmt.Printf("\t\tBroad Address: %s", address.Broadaddr)
			fmt.Printf("\t\tP2P: %s", address.P2P)
			fmt.Println("-----------")
		}
	}
}
