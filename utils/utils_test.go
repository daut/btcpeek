package utils_test

import (
	"testing"

	"github.com/daut/btcpeek/utils"
)

func TestFormatNumber(t *testing.T) {
	tests := []struct {
		name   string
		sats   int64
		locale string
		want   string
	}{
		{
			name:   "US locale formatting",
			sats:   123456789,
			locale: "en-US",
			want:   "123,456,789",
		},
		{
			name:   "Serbian locale formatting",
			sats:   123456789,
			locale: "sr-RS",
			want:   "123.456.789",
		},
		{
			name:   "French locale formatting",
			sats:   123456789,
			locale: "fr-FR",
			want:   "123 456 789",
		},
		{
			name:   "Indian locale formatting",
			sats:   123456789,
			locale: "en-IN",
			want:   "12,34,56,789",
		},
		{
			name:   "Invalid locale defaults to no formatting",
			sats:   123456789,
			locale: "invalid-LOCALE",
			want:   "123456789",
		},
		{
			name:   "Zero sats",
			sats:   0,
			locale: "en-US",
			want:   "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.FormatNumber(tt.sats, tt.locale)
			if tt.want != got {
				t.Errorf("PrettyPrintSats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSatsToBtc(t *testing.T) {
	tests := []struct {
		name string
		sats int64
		want float64
	}{
		{
			name: "Convert 100000000 sats to 1 BTC",
			sats: 100000000,
			want: 1.0,
		},
		{
			name: "Convert 50000000 sats to 0.5 BTC",
			sats: 50000000,
			want: 0.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.SatsToBtc(tt.sats)
			if tt.want != got {
				t.Errorf("SatsToBtc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateBalance(t *testing.T) {
	tests := []struct {
		name   string
		funded int64
		spent  int64
		want   int64
	}{
		{
			name:   "Calculate balance with funded greater than spent",
			funded: 150000,
			spent:  50000,
			want:   100000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.CalculateBalance(tt.funded, tt.spent)
			if tt.want != got {
				t.Errorf("CalculateBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrettyPrintSats(t *testing.T) {
	tests := []struct {
		name   string
		sats   int64
		locale string
		want   string
	}{
		{
			name:   "Pretty print sats in US locale",
			sats:   123456789,
			locale: "en-US",
			want:   "123,456,789 sats",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.PrettyPrintSats(tt.sats, tt.locale)
			if tt.want != got {
				t.Errorf("PrettyPrintSats() = %v, want %v", got, tt.want)
			}
		})
	}
}
