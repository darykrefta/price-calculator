package main

import (
	// "calculator-project/cmdmanager.go"
	"fmt"
	"calculator-project/file-manager"
	"calculator-project/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates { // passa por todos os valores de taxRates
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}

		// TaxIncludedPrices := make([]float64, len(prices)) // declara a array TaxIncludedPrices, com o lenght de prices
		// for priceIndex, price := range prices { // passa por todos os valores de prices
		// 	TaxIncludedPrices[priceIndex] = price * (1 + taxRate) // armazena na array TaxIncludedPrices os valores de price + taxRate
		// }
		// result[taxRate] = TaxIncludedPrices // a array result armazena os valores do TaxIncludedPrices, no index do taxRate respectivo
	}
}
