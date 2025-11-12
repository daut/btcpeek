package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/daut/btcpeek/commands"
)

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestRun(t *testing.T) {
	t.Run("address command", func(t *testing.T) {
		t.Run("valid address", func(t *testing.T) {
			mockApi := createMockApi()
			defer mockApi.Close()

			output := captureOutput(func() {
				os.Setenv("API_BASE_URL", mockApi.URL+"/")
				args := []string{"btcpeek", "address", "1wiz18xYmhRX6xStj2b9t1rwWX4GKUgpv"}
				run(args)
			})

			if !strings.Contains(output, "Address: 1wiz18xYmhRX6xStj2b9t1rwWX4GKUgpv") {
				t.Errorf("expected output to contain address info, got: %s", output)
			}

			if !strings.Contains(output, "Balance: 87909 sats") {
				t.Errorf("expected output to contain balance info, got: %s", output)
			}
		})
	})

	t.Run("tx command", func(t *testing.T) {
		t.Run("valid txid", func(t *testing.T) {
			mockApi := createMockApi()
			defer mockApi.Close()

			output := captureOutput(func() {
				os.Setenv("API_BASE_URL", mockApi.URL+"/")
				args := []string{"btcpeek", "tx", "sampletxid"}
				run(args)
			})

			if !strings.Contains(output, "Transaction ID: sampletxid") {
				t.Errorf("expected output to contain txid info, got: %s", output)
			}

			if !strings.Contains(output, "Fee: 5000 sats") {
				t.Errorf("expected output to contain fee info, got: %s", output)
			}
		})
	})

	t.Run("block command", func(t *testing.T) {
		t.Run("valid blockhash", func(t *testing.T) {
			mockApi := createMockApi()
			defer mockApi.Close()

			output := captureOutput(func() {
				os.Setenv("API_BASE_URL", mockApi.URL+"/")
				args := []string{"btcpeek", "block", "000000000000000015dc777b3ff2611091336355d3f0ee9766a2cf3be8e4b1ce"}
				run(args)
			})

			if !strings.Contains(output, "Height: 363366") {
				t.Errorf("expected output to contain block height info, got: %s", output)
			}

			if !strings.Contains(output, "Transaction Count: 494") {
				t.Errorf("expected output to contain transaction count info, got: %s", output)
			}
		})
	})

	t.Run("invalid command", func(t *testing.T) {
		output := captureOutput(func() {
			args := []string{"btcpeek", "invalidcmd"}
			run(args)
		})

		if !strings.Contains(output, "Usage: btcpeek <command> [arguments]") {
			t.Errorf("expected output to show usage info for invalid command, got: %s", output)
		}
	})

	t.Run("help command", func(t *testing.T) {
		output := captureOutput(func() {
			args := []string{"btcpeek", "help"}
			run(args)
		})

		if !strings.Contains(output, "Usage: btcpeek <command> [arguments]") {
			t.Errorf("expected output to show usage info for help command, got: %s", output)
		}
	})
}

func createMockApi() *httptest.Server {
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
