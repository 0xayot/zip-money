package money

import (
	"log"
	"strings"

	"github.com/shopspring/decimal"
)

var (
	codeToCurrency map[string]*Currency = make(map[string]*Currency)
)

type Currency struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	MinorFactor uint   `json:"minorFactor"`
	Fiat        bool   `json:"fiat"`
}

// RequireMinorFactorFromCurrency ...
func (currency *Currency) RequireMinorFactorFromCurrency() decimal.Decimal {
	failWhenCurrenciesAreNotInitialised()
	return decimal.NewFromInt(int64(currency.MinorFactor))
}

// ToCurrency ...
// This panics if the currency cannot be found
func ToCurrency(currencyCode string) Currency {
	failWhenCurrenciesAreNotInitialised()
	initStringToCurrencyMap()
	currencyVal := codeToCurrency[strings.ToUpper(currencyCode)]
	if currencyVal == nil {
		log.Panicf("Unable to find currency for %v", currencyCode)
	}
	return *currencyVal
}

func (currency Currency) String() string {
	failWhenCurrenciesAreNotInitialised()
	return currency.Code
}

func AllCurrencies() []Currency {
	failWhenCurrenciesAreNotInitialised()
	return initialisedCurrencies
}

func initStringToCurrencyMap() {
	if len(codeToCurrency) == len(initialisedCurrencies) {
		return
	}
	for i, currency := range initialisedCurrencies {
		codeToCurrency[currency.Code] = &initialisedCurrencies[i]
	}
}

func failWhenCurrenciesAreNotInitialised() {
	if !initialised {
		log.Fatal("Currencies have not been initialised")
	}
}
