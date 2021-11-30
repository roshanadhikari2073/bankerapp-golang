package src

import "fmt"

// sqlconn "cliapplications/dataconfig"

func CheckBalance(custinf ...map[string]string) {
	fmt.Printf(" TOTAL BALANCE %s \n", custinf[0]["totalbalance"])
}
func TakeLoan() string {
	return "hello:"
}
func Topup() string {
	return "hello:"
}
func RepayLoan() string {
	return "hello:"
}
