package main

import (
	"fmt"
	"strings"

	"jairxortega.tech/go/cryptoterminal/api"
)

const totalLineWidth = 20

func main() {

	for _, ticker := range Tickers {
		getCurrencyData(ticker)
	}
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err == nil {
		priceStr := fmt.Sprintf("%.2f", rate.Price)
		// Calculate the number of dashes before and after the price
		dashesBefore := 7 - len(rate.Currency)
		dashesAfter := totalLineWidth - len(rate.Currency) - len(priceStr) - dashesBefore
		dashesAgain := strings.Repeat("-", dashesAfter)

		// Ensure we have at least 1 dash between the ticker and the price
		if dashesBefore < 1 {
			dashesBefore = 1
		}

		dashes := strings.Repeat("-", dashesBefore)
		// Print the formatted line with the price starting 7 characters from the right
		fmt.Printf("%s%s%s%s\n", rate.Currency, dashes, priceStr, dashesAgain)
	}
}
