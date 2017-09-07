package main

import (
	"fmt"
)

func main() {
	var (
		money float64 = 10000
		rate  float64 = 0.12
	)
	for i := 1; i <= 12*38; i++ {
		ic := calc(money, rate, i)
		c := calcNoInterest(money, i)
		fmt.Printf("saving [%d] int : [%f] , normal [%f] , diff [%f]\n", i, ic, c, ic-c)
	}
}

func calc(money float64, interestRate float64, month int) (total float64) {
	for i := 0; i < month; i++ {
		total = total * float64(1+interestRate/12)
		total += money
	}
	return
}

func calcNoInterest(money float64, month int) float64 {
	return money * float64(month)
}
