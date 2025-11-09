package commands

import (
	"fmt"

	"github.com/daut/btcpeek/clients"
)

type BlockInfo struct {
	Id                string  `json:"id"`
	Height            int     `json:"height"`
	Version           int     `json:"version"`
	Timestamp         int64   `json:"timestamp"`
	TxCount           int     `json:"tx_count"`
	Size              int     `json:"size"`
	Weight            int     `json:"weight"`
	MerkleRoot        string  `json:"merkle_root"`
	PreviousBlockHash string  `json:"previousblockhash"`
	MedianTime        int64   `json:"mediantime"`
	Nonce             int     `json:"nonce"`
	Bits              int     `json:"bits"`
	Difficulty        float64 `json:"difficulty"`
}

func LookupBlock(blockHash string) {
	fmt.Println("Looking up block:", blockHash)
	blockInfo, err := clients.FetchData[BlockInfo]("block/" + blockHash)
	if err != nil {
		println("Error fetching block data:", err.Error())
		return
	}

	fmt.Println("=================================")
	fmt.Printf("Block Hash: %s\n", blockInfo.Id)
	fmt.Println("=================================")
	fmt.Printf("Height: %d\n", blockInfo.Height)
	fmt.Printf("Timestamp: %d\n", blockInfo.Timestamp)
	fmt.Printf("Transaction Count: %d\n", blockInfo.TxCount)
	fmt.Printf("Size: %d bytes\n", blockInfo.Size)
	fmt.Printf("Weight: %d WU\n", blockInfo.Weight)
	fmt.Printf("Merkle Root: %s\n", blockInfo.MerkleRoot)
	fmt.Printf("Previous Block Hash: %s\n", blockInfo.PreviousBlockHash)
	fmt.Printf("Nonce: %d\n", blockInfo.Nonce)
	fmt.Printf("Difficulty: %f\n", blockInfo.Difficulty)
}
