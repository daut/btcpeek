package main

import (
	"os"
	"strings"
	"testing"

	"github.com/daut/btcpeek/testutils"
	"github.com/daut/btcpeek/utils"
)

func TestRun(t *testing.T) {
	t.Run("address command", func(t *testing.T) {
		t.Run("valid address", func(t *testing.T) {
			mockApi := testutils.CreateMockApi()
			defer mockApi.Close()

			output := utils.CaptureOutput(func() {
				os.Setenv("BTCPEEK_API_BASE_URL", mockApi.URL+"/")
				args := []string{"btcpeek", "address", "1wiz18xYmhRX6xStj2b9t1rwWX4GKUgpv"}
				run(args)
			})

			if !strings.Contains(output, "Address: 1wiz18xYmhRX6xStj2b9t1rwWX4GKUgpv") {
				t.Errorf("expected output to contain address info, got: %s", output)
			}

			if !strings.Contains(output, "Balance: 87,909 sats") {
				t.Errorf("expected output to contain balance info, got: %s", output)
			}
		})
	})

	t.Run("tx command", func(t *testing.T) {
		t.Run("valid txid", func(t *testing.T) {
			mockApi := testutils.CreateMockApi()
			defer mockApi.Close()

			output := utils.CaptureOutput(func() {
				os.Setenv("BTCPEEK_API_BASE_URL", mockApi.URL+"/")
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
			mockApi := testutils.CreateMockApi()
			defer mockApi.Close()

			output := utils.CaptureOutput(func() {
				os.Setenv("BTCPEEK_API_BASE_URL", mockApi.URL+"/")
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

	t.Run("latest command", func(t *testing.T) {
		t.Run("fetch latest blocks", func(t *testing.T) {
			mockApi := testutils.CreateMockApi()
			defer mockApi.Close()

			output := utils.CaptureOutput(func() {
				os.Setenv("BTCPEEK_API_BASE_URL", mockApi.URL+"/")
				args := []string{"btcpeek", "latest"}
				run(args)
			})

			if !strings.Contains(output, "Block Height: 700002") {
				t.Errorf("expected output to contain latest block height info, got: %s", output)
			}
		})
	})

	t.Run("help/invalid command", func(t *testing.T) {
		commands := []string{"help", "invalidcmd"}
		for _, cmd := range commands {
			output := utils.CaptureOutput(func() {
				args := []string{"btcpeek", cmd}
				run(args)
			})

			if !strings.Contains(output, "Usage: btcpeek <command> [arguments]") {
				t.Errorf("expected output to show usage info for %s command, got: %s", cmd, output)
			}
		}
	})
}
