package testutils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/daut/btcpeek/commands"
)

func CreateMockApi() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/address") {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(commands.AddressInfo{
				Address: "1wiz18xYmhRX6xStj2b9t1rwWX4GKUgpv",
				ChainStats: commands.AddressStats{
					FundedTxoCount: 5,
					FundedTxoSum:   15007686949,
					SpentTxoCount:  5,
					SpentTxoSum:    15007599040,
					TxCount:        7,
				},
				MempoolStats: commands.AddressStats{
					FundedTxoCount: 0,
					FundedTxoSum:   0,
					SpentTxoCount:  0,
					SpentTxoSum:    0,
					TxCount:        0,
				},
			})
		} else if strings.Contains(r.URL.Path, "/blocks") {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode([]commands.BlockInfo{
				{
					Id:        "0000000000000000000sampleblockhash1",
					Height:    700002,
					Timestamp: 1617182920,
					TxCount:   2000,
				},
				{
					Id:        "0000000000000000000sampleblockhash2",
					Height:    700001,
					Timestamp: 1617182420,
					TxCount:   1800,
				},
			})
		} else if strings.Contains(r.URL.Path, "/block") {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(commands.BlockInfo{
				Id:                "000000000000000015dc777b3ff2611091336355d3f0ee9766a2cf3be8e4b1ce",
				Height:            363366,
				Version:           2,
				Timestamp:         1435766771,
				TxCount:           494,
				Size:              286494,
				Weight:            1145976,
				MerkleRoot:        "9d3cb87bf05ebae366b4262ed5f768ce8c62fc385c3886c9cb097647b04b686c",
				PreviousBlockHash: "000000000000000010c545b6fa3ef1f7cf45a2a8760b1ee9f2e89673218207ce",
				MedianTime:        1435763435,
				Nonce:             404111758,
				Bits:              404111758,
				Difficulty:        49402014931.22746,
			})
		} else if strings.Contains(r.URL.Path, "/tx") {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(commands.TransactionInfo{
				TxId:     "sampletxid",
				Version:  2,
				LockTime: 0,
				Size:     225,
				Weight:   900,
				SigOps:   2,
				Fee:      5000,
				Status: commands.TransactionStatus{
					Confirmed:   true,
					BlockHeight: 700000,
					BlockHash:   "0000000000000000000sampleblockhash",
					BlockTime:   1617181920,
				},
				Vins:  []commands.TransactionVin{},
				Vouts: []commands.TransactionVout{},
			})
		} else {
			w.WriteHeader(404)
		}
	}))
}
