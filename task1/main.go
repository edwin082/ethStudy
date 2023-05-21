package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Fatal("Error loading .env file")
			return
		}

		GOERLI_HTTPS := os.Getenv("GOERLI_HTTPS")
    client, err := ethclient.Dial(GOERLI_HTTPS)
    if err != nil {
        log.Fatal(err)
				return
    }

		START_BLOCK_NUMBER := os.Getenv("START_BLOCK_NUMBER")
		START_BLOCK_NUMBER_INT64, err := strconv.ParseInt(START_BLOCK_NUMBER, 10, 64)
		if err != nil {
			log.Fatal(err)
			return
		}

		END_BLOCK_NUMBER_INT64 := big.NewInt(START_BLOCK_NUMBER_INT64 + 1000)
		EVENT := os.Getenv("EVENT")
		eventSignature := []byte(EVENT)
		eventSignatureHash := crypto.Keccak256Hash(eventSignature)
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(START_BLOCK_NUMBER_INT64),
			ToBlock: END_BLOCK_NUMBER_INT64,
			Topics: [][]common.Hash{
				{eventSignatureHash},
			},
		}

		logs, err := client.FilterLogs(context.Background(), query)
		if err != nil{
			log.Fatal(err)
			return
		}

		for _, log := range logs{
			fmt.Println("Address:",log.Address)
			fmt.Println("Topics:", log.Topics)
			fmt.Println()

			fmt.Println()
			fmt.Println("Emitted Contract Event: From, To")
			fmt.Println("Owner - From:", log.Topics[1])
			fmt.Println("Owner - To:", log.Topics[2])
			fmt.Println()
			
			fmt.Println("Data:", log.Data)
			fmt.Println("BlockNumber:", log.BlockNumber)
			fmt.Println("TxHash:", log.TxHash)
			fmt.Println("TxIndex:", log.TxIndex)
			fmt.Println("BlockHash:", log.BlockHash)
			fmt.Println("Index:", log.Index)
			fmt.Println("Removed:", log.Removed)
		}
}