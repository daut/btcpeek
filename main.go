package main

import "os"

func main() {
	command := os.Args[1]

	switch command {
	case "address":
		if len(os.Args) < 3 {
			println("Usage: btcpeek address <address>")
			return
		}
		lookupAddress(os.Args[2])
	}
}
