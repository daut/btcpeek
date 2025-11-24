package commands

import (
	"fmt"

	"github.com/daut/btcpeek/utils"
)

type AddressInfo struct {
	Address      string       `json:"address"`
	ChainStats   AddressStats `json:"chain_stats"`
	MempoolStats AddressStats `json:"mempool_stats"`
}

type AddressStats struct {
	FundedTxoCount int64 `json:"funded_txo_count"`
	FundedTxoSum   int64 `json:"funded_txo_sum"`
	SpentTxoCount  int64 `json:"spent_txo_count"`
	SpentTxoSum    int64 `json:"spent_txo_sum"`
	TxCount        int64 `json:"tx_count"`
}

func (s *Command) LookupAddress(address string) {
	fmt.Println("Looking up address:", address)
	var addressInfo *AddressInfo
	err := s.client.FetchData("address/"+address, &addressInfo)
	if err != nil {
		println("Error fetching address data:", err.Error())
		return
	}

	balance := utils.PrettyPrintSats(utils.CalculateBalance(addressInfo.ChainStats.FundedTxoSum, addressInfo.ChainStats.SpentTxoSum), "en-US")
	totalReceived := utils.PrettyPrintSats(addressInfo.ChainStats.FundedTxoSum, "en-US")
	totalSpent := utils.PrettyPrintSats(addressInfo.ChainStats.SpentTxoSum, "en-US")
	transactionCount := utils.FormatNumber(addressInfo.ChainStats.TxCount, "en-US")

	fmt.Println("=================================")
	fmt.Printf("Address: %s\n", addressInfo.Address)
	fmt.Println("=================================")
	fmt.Printf("Balance: %s\n", balance)
	fmt.Printf("Total Received: %s\n", totalReceived)
	fmt.Printf("Total Spent: %s\n", totalSpent)
	fmt.Printf("Transaction Count: %s\n", transactionCount)
	fmt.Println("=================================")
}
