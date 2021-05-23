package main

import (
	"fmt"
	"log"
)

func (cli *CLI) startNode(nodeID, minerAddress string) {
	fmt.Printf("Starting node %s\n", nodeID)

	if len(minerAddress) > 0 {
		if ValidateAddress(minerAddress) {
			fmt.Println("Minig is on. Address to Receive rewards: ", minerAddress)
		} else {
			log.Panic("Wrong miner address!")
		}
	}

	StartServer(nodeID, minerAddress)
}
