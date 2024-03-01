package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"jairxortega.tech/go/cryptoterminal/api"
)

const totalLineWidth = 35

func main() {
	clearConsole()

	for _, ticker := range Tickers {
		getCurrencyData(ticker)
	}
}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err == nil {

		// randchars := []string{
		// 	"*",
		// 	"%",
		// 	"$",
		// 	"&",
		// 	"@",
		// 	"!",
		// 	"^",
		// 	"~",
		// 	"+",
		// 	"?",
		// 	"/",
		// 	"|",
		// 	"<",
		// 	">",
		// }

		// rand.Seed(time.Now().UnixNano())
		// randomChar := randchars[rand.Intn(len(randchars))]
		// priceStr := fmt.Sprintf("%.2f%s", rate.Price, randomChar)
		priceStr := fmt.Sprintf("%.2f", rate.Price)

		red := "\033[31m"
		green := "\033[32m"
		reset := "\033[0m"
		yellow := "\033[33m"

		var percentChangeStr string
		if rate.PercentChange == "" {
			percentChangeStr = "0.00%"
		} else if strings.HasPrefix(rate.PercentChange, "-") {
			percentChangeStr = fmt.Sprintf("%s%s%%%s", red, rate.PercentChange, reset)
		} else {
			percentChangeStr = fmt.Sprintf("%s%s%%%s", green, rate.PercentChange, reset)
		}

		// Calculate the spacing and dashes
		dashesBefore := 7 - len(rate.Currency)
		if dashesBefore < 1 {
			dashesBefore = 1
		}
		dashes := strings.Repeat("-", dashesBefore)

		dashesAfter := 18 - len(rate.Currency) - len(dashes) - len(priceStr)
		if strings.HasPrefix(rate.PercentChange, "-") {
			dashesAfter -= 1 // Subtract one more dash if percent change is negative
		}
		if dashesAfter < 1 {
			dashesAfter = 1 // Ensure at least one dash
		}
		dashesInbetween := strings.Repeat("-", dashesAfter)

		dashesEnd := totalLineWidth - len(rate.Currency) - len(dashes) - len(priceStr) - len(dashesInbetween) - len(percentChangeStr)
		if dashesEnd < 0 {
			dashesEnd = 0
		}
		dashesFinish := strings.Repeat("-", dashesEnd)

		fmt.Printf("%s%s%s%s%s%s%s%s%s%s%s%s\n", yellow, rate.Currency, dashes, reset, priceStr, yellow, dashesInbetween, percentChangeStr, reset, yellow, dashesFinish, reset)
	}
}

func clearConsole() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux", "darwin":
		cmd = exec.Command("clear") // Unix-like OS
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls") // Windows
	default:
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
