package utils

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func CaptureOutput(f func()) string {
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

func CalculateBalance(funded, spent int64) int64 {
	return funded - spent
}

func SatsToBtc(sats int64) float64 {
	return float64(sats) / 100_000_000
}

func PrettyPrintSats(sats int64, locale string) string {
	num := FormatNumber(sats, locale)
	return fmt.Sprintf("%s sats", num)
}

func FormatNumber(number int64, locale string) string {
	langTag := language.Make(locale)
	if langTag == language.Und {
		return fmt.Sprintf("%d", number)
	}
	printer := message.NewPrinter(langTag)
	return printer.Sprintf("%d", number)
}
