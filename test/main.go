package main

import (
	"fmt"

	"github.com/arthasyou/utility-go/decimal"
)

func main() {
	fmt.Println(decimal.Add("1.34", "2.35"))
	fmt.Println(decimal.Sub("1.34", "2.35"))
	fmt.Println(decimal.Mul("1.34", "2.35"))
	fmt.Println(decimal.Div("1.34", "2.35"))
	fmt.Println(decimal.Abs("-1.34"))
}
