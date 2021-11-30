package src

import (
	"fmt"
	"strconv"
)

var usdrate float32 = 120.17

func CheckBalance(custinf ...map[string]string) {
	fmt.Printf(" TOTAL BALANCE %s \n", custinf[0]["totalbalance"])
}
func TakeLoan() string {
	return "hello:"
}
func Topup() string {
	return "hello:"
}

func PrintBankStatement(custinf ...map[string]string) {
	totalbalance, err := strconv.Atoi(custinf[0]["totalbalance"])
	if err != nil {
		print("parsing error")
	}
	nprtousd := float32(totalbalance) / usdrate
	fmt.Printf(" YOUR BANK AMOUNT IS %f USD $. THE CURRENT EXCHANGE RATE OF USA IS %f \n", nprtousd, usdrate)
}

func RepayLoan() string {
	return "hello:"
}
