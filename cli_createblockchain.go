package main

import "fmt"

func (cli *CLI) createBlockchain(address string) {
	bc := CreateBlockchain(address)
	defer bc.db.Close()

	UTXOSet := UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("DONE!")
}
