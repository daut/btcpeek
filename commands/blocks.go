package commands

import (
	"fmt"

	"github.com/daut/btcpeek/utils"
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

func (s *Command) LookupBlock(blockHash string) {
	fmt.Println("Looking up block:", blockHash)
	var blockInfo *BlockInfo
	err := s.FetchData("block/"+blockHash, &blockInfo)
	if err != nil {
		println("Error fetching block data:", err.Error())
		return
	}

	fmt.Println("=================================")
	fmt.Printf("Block Hash: %s\n", blockInfo.Id)
	fmt.Println("=================================")
	fmt.Printf("Height: %d\n", blockInfo.Height)
	fmt.Printf("Timestamp: %s (%s)\n", utils.FormatIso8601(blockInfo.Timestamp), utils.TimeAgo(blockInfo.Timestamp))
	fmt.Printf("Transaction Count: %d\n", blockInfo.TxCount)
	fmt.Printf("Size: %d bytes\n", blockInfo.Size)
	fmt.Printf("Weight: %d WU\n", blockInfo.Weight)
	fmt.Printf("Merkle Root: %s\n", blockInfo.MerkleRoot)
	fmt.Printf("Previous Block Hash: %s\n", blockInfo.PreviousBlockHash)
	fmt.Printf("Nonce: %d\n", blockInfo.Nonce)
	fmt.Printf("Difficulty: %.2f (%.2e times harder than easiest)\n", blockInfo.Difficulty, blockInfo.Difficulty)
}

func (s *Command) LookupLatestBlocks() {
	fmt.Println("Looking up latest blocks.")
	var latestBlocks []BlockInfo
	err := s.FetchData("blocks/", &latestBlocks)
	if err != nil {
		println("Error fetching latest blocks:", err.Error())
		return
	}

	fmt.Println("=================================")
	for _, blockInfo := range latestBlocks {
		fmt.Printf("Block Height: %d, Hash: %s, Timestamp: %s (%s), Tx Count: %d\n",
			blockInfo.Height, blockInfo.Id, utils.FormatIso8601(blockInfo.Timestamp), utils.TimeAgo(blockInfo.Timestamp), blockInfo.TxCount)
	}
}
