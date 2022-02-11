package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	topics         [][]common.Hash
	logSwapSig     []byte      = []byte("Swap(address,uint256,uint256,uint256,uint256,address)")
	logSwapSigHash common.Hash = crypto.Keccak256Hash(logSwapSig)
)

func main() {
	client, err := ethclient.Dial("wss://mainnet.aurora.dev")
	if err != nil {
		log.Fatal(err)
	}

	// topics = append(topics, []common.Hash{logSwapSigHash})

	query := ethereum.FilterQuery{
		// Topics: topics,
	}

	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}
