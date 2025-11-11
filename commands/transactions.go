package commands

import (
	"fmt"
)

type TransactionInfo struct {
	TxId     string            `json:"txid"`
	Version  int               `json:"version"`
	LockTime int               `json:"locktime"`
	Size     int               `json:"size"`
	Weight   int               `json:"weight"`
	SigOps   int               `json:"sigops"`
	Fee      int64             `json:"fee"`
	Status   TransactionStatus `json:"status"`
	Vins     []TransactionVin  `json:"vin"`
	Vouts    []TransactionVout `json:"vout"`
}

type TransactionStatus struct {
	Confirmed   bool   `json:"confirmed"`
	BlockHeight int    `json:"block_height"`
	BlockHash   string `json:"block_hash"`
	BlockTime   int64  `json:"block_time"`
}

type TransactionVin struct {
	TxId         string  `json:"txid"`
	Vout         int     `json:"vout"`
	Prevout      Prevout `json:"prevout"`
	Scriptsig    string  `json:"scriptsig"`
	ScriptsigAsm string  `json:"scriptsig_asm"`
	IsCoinbase   bool    `json:"is_coinbase"`
	Sequence     int64   `json:"sequence"`
}

type TransactionVout struct {
	Scriptpubkey        string `json:"scriptpubkey"`
	ScriptpubkeyAsm     string `json:"scriptpubkey_asm"`
	ScriptpubkeyType    string `json:"scriptpubkey_type"`
	ScriptpubkeyAddress string `json:"scriptpubkey_address"`
	Value               int64  `json:"value"`
}

type Prevout struct {
	Scriptpubkey        string `json:"scriptpubkey"`
	ScriptpubkeyAsm     string `json:"scriptpubkey_asm"`
	ScriptpubkeyType    string `json:"scriptpubkey_type"`
	ScriptpubkeyAddress string `json:"scriptpubkey_address"`
	Value               int64  `json:"value"`
}

func (s *Command) LookupTransaction(txId string) {
	fmt.Println("Looking up transaction:", txId)
	var txInfo *TransactionInfo
	err := s.FetchData("tx/"+txId, &txInfo)
	if err != nil {
		println("Error fetching transaction data:", err.Error())
		return
	}

	fmt.Println("=================================")
	fmt.Printf("Transaction ID: %s\n", txInfo.TxId)
	fmt.Println("=================================")
	fmt.Printf("Status: %s\n", func() string {
		if txInfo.Status.Confirmed {
			return fmt.Sprintf("Confirmed in block %d", txInfo.Status.BlockHeight)
		}
		return "Unconfirmed"
	}())
	fmt.Printf("Fee: %d sats\n", txInfo.Fee)
	fmt.Printf("Size: %d bytes\n", txInfo.Size)
	fmt.Printf("Size vBytes: %d vBytes\n", txInfo.Weight/4)
}
