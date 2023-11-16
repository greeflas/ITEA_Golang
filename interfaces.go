package main

import "fmt"

type AmountConvertionRateProvider interface {
	GetCurrentRate() float64
}

type AmountConverter interface {
	Convert(amount float64) float64
}

type Converter interface {
	AmountConvertionRateProvider
	AmountConverter
}

type MonobankConverter struct {
	currencyRage float64
}

func NewMonobankConverter() *MonobankConverter {
	// TODO: fetch real rate from API
	const fixedCurrencyRate = 37.30

	return &MonobankConverter{currencyRage: fixedCurrencyRate}
}

func (c MonobankConverter) GetCurrentRate() float64 {
	return c.currencyRage
}

func (c MonobankConverter) Convert(amount float64) float64 {
	return c.GetCurrentRate() * amount
}

type PrivatbankConverter struct {
	currentCurrencyRage float64
}

func NewPrivatbankConverter() *PrivatbankConverter {
	// TODO: fetch real rate from API
	const fixedCurrencyRate = 37.55

	return &PrivatbankConverter{
		currentCurrencyRage: fixedCurrencyRate,
	}
}

func (c PrivatbankConverter) GetCurrentRate() float64 {
	return c.currentCurrencyRage
}

func (c PrivatbankConverter) Convert(amount float64) float64 {
	return c.GetCurrentRate() * amount
}

func main() {
	amountUSD := 1000.00

	monobankConverter := NewMonobankConverter()
	privatConverter := NewPrivatbankConverter()

	converters := []Converter{
		monobankConverter,
		privatConverter,
	}

	showLowestRate(amountUSD, converters...)
	fmt.Println()

	rateProviders := []AmountConvertionRateProvider{
		monobankConverter,
		privatConverter,
	}

	showConvertionRates(rateProviders)
}

func showLowestRate(amount float64, converters ...Converter) {
	var (
		lowestRate      float64
		lowestConverter Converter
	)

	for i, converter := range converters {
		currentRate := converter.GetCurrentRate()
		if i == 0 || currentRate < lowestRate {
			lowestRate = currentRate
			lowestConverter = converter
		}
	}

	fmt.Printf("Lowest rate is %.2f, so you will get %.2f UAH\n", lowestRate, lowestConverter.Convert(amount))
}

func showConvertionRates(rateProviders []AmountConvertionRateProvider) {
	for _, provider := range rateProviders {
		fmt.Printf("Rate: %.2f\n", provider.GetCurrentRate())
	}
}
