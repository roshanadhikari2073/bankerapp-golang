package main

import (
	"bufio"
	getLogo "cliapplications/assets"
	sqlconn "cliapplications/dataconfig"
	"fmt"
	"os"
	"time"
)

//start portion of the application...
var customerGlobalScope customerInfo
var clearT string = "clearT"
var checkCreds string = `** WARNING ** PLEASE PRESS CORRECT INPUTS `
var exitSignalText string = "PRESS ENTER TO RETURN TO THE MAIN MENU OR PRESS 9 TO EXIT"
var goodByeNote string = " -------  CLOSING THE APPLICATION $$$$$$$ THANKS FOR VISITING --------"
var timeToday string

// adding the needed functions

type customerInfo struct {
	name          string
	user_type     string
	address       string
	phone         int
	total_balance int
	total_loan    int
}

func main() {
	loginPage()
}

func init() {
	timeToday = time.Now().Format(time.RFC850)
}

func clearTheTerminal(s string) bool {
	forceClear := func() {
		fmt.Print("\033[H\033[2J")
	}
	if s == clearT {
		forceClear()
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(s)
		text, _ := reader.ReadString('\n')
		if text == "\n" || text != "\n" {
			forceClear()
		}
	}
	return false
}

// func inner() {
// 	panic("unimplemented")

// }

func welcomeloop(cont bool, status string, updateInfo bool) {
	clearTheTerminal(clearT)
	customerGlobalScope := sqlconn.Show("roshan")

	if cont {
		// add the ending parameters
		fmt.Println(status)
		fmt.Println(getLogo.BankLogo())
		spacingToTheExit("", 3)
		fmt.Println("-------   WELCOME TO THE BANKING APPLICATIONS    -------    ")
		fmt.Printf("-  %s  -    ", timeToday)
		spacingToTheExit("", 3)
		fmt.Println("HINT -> TYPE NUMBERS ASSOCIATED WITH THE MODULES MENTIONED BELOW")
		spacingToTheExit("", 2)
		fmt.Println("|-----------------------------------------------|")
		fmt.Printf(" Customer Name  |        %s          \n|", customerGlobalScope.Name)
		fmt.Println("-----------------------------------------------|")
		fmt.Printf(" Address        |        %s          \n|", customerGlobalScope.Address)
		fmt.Println("-----------------------------------------------|")
		fmt.Printf(" Phone Number   |        %d          \n|", customerGlobalScope.Phone)
		fmt.Println("-----------------------------------------------|")
		spacingToTheExit("", 2)
		fmt.Println("[ 1 ]  -> |      CHECK BALANCE              |")
		fmt.Println("")
		fmt.Println("[ 2 ]  -> |      TAKE LOAN                  |")
		fmt.Println("")
		fmt.Println("[ 3 ]  -> |      TOP UP BALANCE             |")
		fmt.Println("")
		fmt.Println("[ 4 ]  -> |      CHECK YOUR BANK STATEMENT  |")
		fmt.Println("")
		if true {
			fmt.Println("[ 5 ]  -> |      REPAY THE LOAN             |")
		}
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("******  press 9 to exit  ******")
		_, int := takeTheUserInput("int")
		switch int {
		case 1, 2, 3, 4:
			bankingModules(int, "")
		case 5:
			if true {
				bankingModules(int, "")
			} else {
				welcomeloop(true, checkCreds, false)
			}

		case 9:
			clearTheTerminal(clearT)
			spacingToTheExit(".", 4)
			println(goodByeNote)
			spacingToTheExit(".", 4)
			cont = false
		default:
			welcomeloop(true, checkCreds, false)
		}
	} else {
		// add a ending paramter
		spacingToTheExit(".", 4)
		println(goodByeNote)
		spacingToTheExit(".", 4)
	}

}

// the main login page of the application
func loginPage() {
	clearTheTerminal(clearT)
	// println("ENTER THE RIGHT CREDENTIALS TO ACCESS THE BANKING APPLICATION")
	// print("USERNAME - ")
	// username, _ := takeTheUserInput("str")
	// print("PASSWORD - ")
	// _, password := takeTheUserInput("int")
	// if username == "roshan" && password == 123 {
	if true {
		welcomeloop(true, "", true)
	} else {
		// right credentials to get the input
		println("ERROR")

	}

}

// TODO: Implement Interface here and learn more about it
func takeTheUserInput(dataType string) (string, int) {
	username, password := "", 0
	if dataType == "str" {
		fmt.Scanf("%s", &username)
		return username, password
	} else if dataType == "int" {
		fmt.Scanf("%d", &password)
		return username, password
	} else {
		panic("error while parsing the correct datatype")
	}
}

func bankingModules(head int, blockStat string) {
	clearTheTerminal(clearT)
	if blockStat == "blocked" {
		fmt.Println(checkCreds)
		spacingToTheExit("", 4)
	}
	header := [6]string{"", "MAIN BALANCE", "TAKE THE LOAN", "TOP UP BALANCE", "CHECK EXPENDITURE"}
	// create if there is loan to be paid
	if true {
		header[5] = "REPAY THE LOAN"
	}
	println(header[head])
	spacingToTheExit("", 4)
	if head == 1 {
		println(customerGlobalScope.total_balance)
		// to check the main balance of the user
	} else if head == 2 {
		customerGlobalScope.total_balance = 100
		// take the loan
	} else if head == 3 {
		// top up the balance

	} else if head == 4 {
		// check expenditure

	} else if head == 5 {
		println("5")
		//repay the loan
	}
	checkStat, Status := exitTextSignal(head)
	if Status == "" {
		welcomeloop(checkStat, Status, false)
	}
}

func exitTextSignal(currentInt int) (bool, string) {
	println(exitSignalText)
	var reader string
	fmt.Scanf("%s", &reader)
	if reader == "" {
		return true, ""
	} else if reader == "9" {
		return false, ""
	} else {
		bankingModules(currentInt, "blocked")
		return false, ""
	}
}

func spacingToTheExit(char string, totalspace int) {
	j := 0
	for {
		fmt.Println(char)
		j++
		if j >= totalspace {
			break
		}
	}
}
