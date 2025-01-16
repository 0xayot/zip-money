package money

import (
	"encoding/json"
	"fmt"

	"io"
	"log"
	"os"

	"github.com/shopspring/decimal"
)

var (
	initialisedCurrencies []Currency
	initialised           bool
)

// Money ...
type Money struct {
	Currency Currency
	Amount   decimal.Decimal
}

func (money *Money) String() string {
	return fmt.Sprintf("%v %v", money.Currency, money.Amount)
}

// From ...
func From(currency Currency, amount decimal.Decimal) Money {
	return Money{Currency: currency, Amount: amount}
}

// InitFromCurrencies ...
// This should be called during startup since it panics if the currencies cannot be initialised
func InitFromCurrencies(currencies []Currency) {
	validateCurrencies(currencies)
	initialisedCurrencies = currencies
	initialised = true
}

// InitFromJsonString ...
// This should be called during startup since it panics if the currencies cannot be initialised
func InitFromJsonString(jsonString string) {
	var currencies []Currency
	if err := json.Unmarshal([]byte(jsonString), &currencies); err != nil {
		log.Fatal(err)
	}

	InitFromCurrencies(currencies)
}

// InitFromFile ...
// This should be called during startup since it panics if the currencies cannot be initialised
func InitFromFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	contents, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	InitFromJsonString(string(contents))
}

func GetCurrencies() []Currency {
	if len(initialisedCurrencies) == 0 {
		log.Fatal("Currencies not yet initialised")
	}
	return initialisedCurrencies
}

func validateCurrencies(currencies []Currency) {
	genValidationError := func(field, expectation, reality string) string {
		return fmt.Sprintf("Expected %v to be %v but got %v instead", field, expectation, reality)
	}

	if len(currencies) == 0 {
		log.Fatal(genValidationError("Currencies", "a populated list", "an empty list"))
	}

	for i, config := range currencies {
		if config.Code == "" {
			log.Fatal(genValidationError(fmt.Sprintf("Code[%v]", i), "a valid currency code", "an empty/unset string"))
		}

		if config.MinorFactor == 0 {
			log.Fatal(genValidationError(fmt.Sprintf("MinorFactor[%v]", i), "a valid minor factor", "a zero/unset value"))
		}

		if config.Name == "" {
			log.Fatal(genValidationError(fmt.Sprintf("Name[%v]", i), "a valid currency name", "an empty/unset string"))
		}
	}
}

func DefaultCurrencies() []Currency {
	return []Currency{
		{
			Name:        "Bitcoin",
			Code:        "BTC",
			MinorFactor: 100000000,
		},
		{
			Name:        "Ethereum",
			Code:        "ETH",
			MinorFactor: 1000000000,
		},
		{
			Name:        "Naira",
			Code:        "NGN",
			MinorFactor: 100,
			Fiat:        true,
		},
		{
			Name:        "Dollar",
			Code:        "USD",
			MinorFactor: 100,
			Fiat:        true,
		},
		{
			Name:        "USD Coin",
			Code:        "USDC",
			MinorFactor: 100000000,
		},
		{
			Name:        "USD Tether",
			Code:        "USDT",
			MinorFactor: 100,
		},
	}
}
