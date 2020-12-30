package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

var iface = flag.String("i", "wlp0s20f3", "Interface to use")
var snaplen = flag.Int("s", 65535, "Snap length (max number of bytes to read from each packets")
var timeout = flag.Int("t", -1, "Packet timeout (Per second)")
var promisc = flag.Bool("p", false, "Set promiscuous")
var bpf = flag.String("bpf", "tcp and port 80", "Set packet filtering (default='tcp and port 80')")
var fileName = flag.String("f", "capture.pcap", "PCAP Filename to save the packets (Default: capture.pcap")
var nPacketsSave = flag.Int("n", 50, "Number of packets to save (default: 50)")

func main() {
	flag.Parse()
	handler, err := pcap.OpenLive(*iface, int32(*snaplen), *promisc, time.Second*time.Duration(*timeout))
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create(*fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	writer := pcapgo.NewWriter(file)
	writer.WriteFileHeader(uint32(*snaplen), handler.LinkType())

	nPackets := 0

	if err := handler.SetBPFFilter(*bpf); err != nil {
		log.Fatalln(err)
	}

	for packet := range gopacket.NewPacketSource(handler, handler.LinkType()).Packets() {
		// fmt.Println(packet)
		// fmt.Println(packet.String())
		fmt.Println(packet.Dump())
		// fmt.Println(string(packet.Data()))
		fmt.Println("-------------")

		nPackets++
		if nPackets <= *nPacketsSave {
			writer.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		}
	}
}
