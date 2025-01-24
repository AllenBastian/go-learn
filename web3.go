package main

import (
	"context"
	"fmt"

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

}
