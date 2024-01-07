package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)


func main() {
	password := "123456"
	// key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	// account, err := key.NewAccount(password)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(account.Address.Hex())

	// 0xB7b3A53ec7fA9f07BE877c718e5A51f6Cbaae241

	b, err := ioutil.ReadFile("./wallet/UTC--2024-01-06T21-01-53.754274600Z--bd917e7b7e644c95a965224eda57535eca570f2a")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal(err)
	}

	// Private key
	pData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("Private key:", hexutil.Encode(pData))

	// Public key
	pubData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("Public key:", hexutil.Encode(pubData))

	// Address
	address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()
	fmt.Println("Address:", address)
}
