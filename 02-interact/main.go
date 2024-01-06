package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ganacheURL = "http://localhost:8545"

func main() {
	client, err := ethclient.DialContext(context.Background(), ganacheURL)
	if err != nil {
		log.Fatal("Error to create a ethere client:%v", err)
	}
	defer client.Close()
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal("Error to get the block:%v", err)
	}
	fmt.Println("block number: ", block.Number())

	addr := "0x8eFA5E4dA6C154Db76E94837ea5036b45532AF06"

	fmt.Println("The address: ",addr)

	address := common.HexToAddress(addr) //because we need the hex of the address

	fmt.Println("The address: ",address)

	balance, err := client.BalanceAt(context.Background(), address, nil) //get the balance of the address
	if err != nil {
		log.Fatal("Error to get the balance:%v", err)
	}
	fmt.Println("The balance: ",balance)
}