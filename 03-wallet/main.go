package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

var ganacheURL = "http://localhost:8545"

func main() {
	pvk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	pData := crypto.FromECDSA(pvk)
	fmt.Printf("Private key: %s\n", hexutil.Encode(pData))

	pubKey := crypto.FromECDSAPub(&pvk.PublicKey)
	fmt.Printf("Public key: %s\n", hexutil.Encode(pubKey))

	// Hash the public key using Keccak-256
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pubKey[1:]) // Exclude the 0x04 prefix
	address := hash.Sum(nil)

	// Take the last 20 bytes of the hashed public key to form the address
	fmt.Printf("Address: 0x%x\n", address[len(address)-20:])
}
