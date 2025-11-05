package main

import (
	"fmt"
)

type AddressInfo struct {
	Address      string `json:"address"`
	ChainStats   Stats  `json:"chain_stats"`
	MempoolStats Stats  `json:"mempool_stats"`
}

type Stats struct {
	FundedTxoCount int64 `json:"funded_txo_count"`
	FundedTxoSum   int64 `json:"funded_txo_sum"`
	SpentTxoCount  int64 `json:"spent_txo_count"`
	SpentTxoSum    int64 `json:"spent_txo_sum"`
	TxCount        int64 `json:"tx_count"`
}

func lookupAddress(address string) {
	println("Looking up address:", address)
	addressInfo, err := fetchData[AddressInfo]("address/" + address)
	if err != nil {
		println("Error fetching address data:", err.Error())
		return
	}

	balance := addressInfo.ChainStats.FundedTxoSum - addressInfo.ChainStats.SpentTxoSum

	fmt.Println("=================================")
	fmt.Printf("Address: %s\n", addressInfo.Address)
	fmt.Println("=================================")
	fmt.Printf("Balance: %d sats\n", balance)
	fmt.Printf("Total Received: %d sats\n", addressInfo.ChainStats.FundedTxoSum)
	fmt.Printf("Total Spent: %d sats\n", addressInfo.ChainStats.SpentTxoSum)
	fmt.Printf("Transaction Count: %d\n", addressInfo.ChainStats.TxCount)
	fmt.Println("=================================")
}
