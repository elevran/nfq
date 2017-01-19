package main

import (
	"fmt"
	"log"

	"github.com/subgraph/go-nfnetlink/nfqueue"
)

//
// instructions:
//
// Set up IPTables
// # sudo iptables -A OUTPUT -p tcp -j NFQUEUE --queue-num 1 --queue-bypass
//
// Run this program
// # sudo ./nfq
//
// In another terminal, try initiating some outbound connection (e.g., curl www.google.com)
//
// Delete the rule using
// # sudo iptables -L OUTPUT --line-numbers
// followed by
// # sudo iptables -D OUTPUT <line number>
// or
// sudo iptables -D OUTPUT -p tcp -j NFQUEUE --queue-num 1 --queue-bypass


func main() {
	q := nfqueue.NewNFQueue(1)

	ps, err := q.Open()
	if err != nil {
		log.Fatal("Error opening NFQueue:", err)
	}
	defer q.Close()

	for p := range ps {
		fmt.Printf("Packet: %v\n", p.Packet)
		p.Accept()
	}
}

