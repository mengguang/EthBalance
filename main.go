package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"fmt"
	"flag"
	"time"
)

func main() {

	rpcUrl := flag.String("rpcUrl","https://rpc2.newchain.cloud.diynova.com/","Geth json rpc or ipc url")
	hexAddress := flag.String("address","0x688F271A0A7a0BDfA41E8b8911d8389fCca1f52D","Ethereum address")
	flag.Parse()

	client,err := ethclient.Dial(*rpcUrl)
	address := common.HexToAddress(*hexAddress)

	if err != nil {
		fmt.Println("Dial error:",err)
		os.Exit(1)
	}
	d := time.Now().Add(5000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(),d)
	defer cancel()
	balance, err := client.BalanceAt(ctx, address,nil)
	if err != nil {
		fmt.Println("error:",err)
		os.Exit(2)
	}
	fmt.Println("Balance:",balance.String())
	os.Exit(0)
}

