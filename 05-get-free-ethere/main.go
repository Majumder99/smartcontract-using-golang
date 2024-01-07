package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	url = "https://sepolia.infura.io/v3/2MJEo6qegjyxHarloYlEKgpUREW"
)

func main() {
	// password := "123456"
	// key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	// account, err := key.NewAccount(password)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(account.Address.Hex())

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	account1 := common.HexToAddress("0xcBf76bfC3931fde50B5c7c76B2062493Dbbe1dC6")
	fmt.Println("account1 address is : ", account1)

	balance, err := client.BalanceAt(context.Background(), account1, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("account balance is : ", balance)
}
