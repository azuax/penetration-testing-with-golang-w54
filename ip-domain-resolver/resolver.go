package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s [IPv4|IPv6|HOSTNAME]\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	parse := net.ParseIP(name)
	if parse == nil {
		// if is not an IP
		ips, err := net.LookupIP(name)
		if err != nil {
			fmt.Println("Can't resolve IP:", err)
			os.Exit(1)
		}
		fmt.Println("Resolved:")
		for _, ip := range ips {
			fmt.Printf("\t%s\n", ip.String())
		}
	} else {
		// the parameter requests is an IP
		addrs, err := net.LookupAddr(name)
		if err != nil {
			fmt.Println("Can't resolve address:", err)
			os.Exit(1)
		}
		fmt.Println("Resolved:")
		for _, addr := range addrs {
			fmt.Printf("\t%s\n", addr)
		}
	}

	// check for mx records
	mxRecords, err := net.LookupMX(name)
	if err != nil {
		fmt.Println("Error on MX lookup", err.Error())
		os.Exit(1)
	}
	for _, mx := range mxRecords {
		fmt.Printf("Host: %s\tPrecedence: %d\n", mx.Host, mx.Pref)
	}
}
