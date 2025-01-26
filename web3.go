package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	client, error := ethclient.Dial("http://localhost:8001")
	if error != nil {
		panic(error)
	} else {
		_ = client
		fmt.Println("Client created")
	}

	chainid, error := client.ChainID(context.Background())
	if error != nil {
		fmt.Println("error connecting!")
	} else {
		fmt.Println(chainid.String() + "success connection")
	}

	hashStore, err := MyHashStore(common.HexToAddress("0x8CdaF0CD259887258Bc13a92C0a6dA92698644C0"), client)
	if err != nil {
		log.Fatalf("Failed to instantiate the smart contract")
	}

	opts := &bind.CallOpts{Context: context.Background()}
	a, b, c, err := hashStore.GetContractDetails(opts)
	if err != nil {
		log.Fatalf("Error getting contract details: %v", err)
	}

	fmt.Println("Contract details:", a, b, c)
}

func MyHashStore(address common.Address, backend bind.ContractBackend) (*HashStore, error) {
	contract, err := bindHashStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HashStore{HashStoreCaller: HashStoreCaller{contract: contract}, HashStoreTransactor: HashStoreTransactor{contract: contract}, HashStoreFilterer: HashStoreFilterer{contract: contract}}, nil

}
