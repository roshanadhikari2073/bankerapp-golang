package src

import (
	"fmt"
	"strconv"

	"github.com/asvvvad/exchange"
)

//TODO implement a logic or a API to get the realtime data of the USD NPR rate, as for now use it as it is

func CheckBalance(custinf ...map[string]string) {
	fmt.Printf(" TOTAL BALANCE %s \n", custinf[0]["totalbalance"])
}

// TODO implement logic to give loan and take interest in hourly rate
func TakeLoan() string {
	return "hello:"
}
func Topup() string {
	return "hello:"
}

// TODO create a nice visual UI using mathematical geometry
func PrintBankStatement(custinf ...map[string]string) {
	totalbalance, err := strconv.Atoi(custinf[0]["totalbalance"])
	if err != nil {
		print("parsing error")
	}
	ex := exchange.New("NPR")
	xe := exchange.New("USD")

	// convert 10 USD to EUR
	nprtousd, _ := ex.ConvertTo("USD", totalbalance)
	usdrate, _ := xe.ConvertTo("NPR", 1)
	fmt.Printf(" YOUR BANK AMOUNT IS %f USD $. THE CURRENT EXCHANGE RATE OF USA IS %f \n", nprtousd, usdrate)
}

// update the db as the user pays back
func RepayLoan() string {
	return "hello:"
}
