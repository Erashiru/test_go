package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type CryptoCurrency struct {
	ID           string  `json:"id"`
	Symbol       string  `json:"symbol"`
	Name         string  `json:"name"`
	CurrentPrice float64 `json:"current_price"`
}

func getMarketData() ([]CryptoCurrency, error) {
	url := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	code := resp.StatusCode
	if code != 200 {
		return nil, fmt.Errorf("Bad status code: %d", code)
	}

	var data []CryptoCurrency
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func getCurrencyPrice(id string, marketData []CryptoCurrency) (float64, error) {

	for _, currency := range marketData {
		if currency.ID == id {
			return currency.CurrentPrice, nil
		}
	}

	return 0, fmt.Errorf("Currency not found")
}

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("Bad request!")
		return
	}
	id := os.Args[1]
	marketData, err := getMarketData()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	price, err := getCurrencyPrice(id, marketData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Current price of %s: $%.2f\n", id, price)

	ticker := time.NewTicker(10 * time.Minute)
	for range ticker.C {
		marketData, err := getMarketData()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		price, err := getCurrencyPrice(id, marketData)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Printf("Updated price of %s: $%.2f\n", id, price)
	}
}
