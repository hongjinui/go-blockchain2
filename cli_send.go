package main

import "fmt"

func (cli *CLI) send(from, to string, amount int) {
	bc := NewBlockchain()
	UTXOSet := UTXOSet{bc}
	defer bc.db.Close()

	tx := NewUTXOTransaction(from, to, amount, &UTXOSet)
	cbTx := NewCoinbaseTX(from, "")
	txs := []*Transaction{cbTx, tx}

	newBlock := bc.MineBlock(txs)
	UTXOSet.Update(newBlock)

	bc.MineBlock(txs)
	fmt.Println("SUCCESS!")
}
