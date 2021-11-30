package main

import (
	"bufio"
	getLogo "cliapplications/assets"
	sqlconn "cliapplications/dataconfig"
	"cliapplications/src"
	"fmt"
	"os"
	"strings"
	"time"
)

// This leads to the main login page
func main() {
	loginPage()
}

// the main page
func welcomeloop(cont bool, status string, updateTheTable bool, params ...map[string]string) {
	clearTheTerminal(src.CLEARTERMINAL)
	var customerGlobalScope = make(map[string]string)
	if updateTheTable {
		customerGlobalScope = sqlconn.Show("roshan")
	} else {
		customerGlobalScope = params[0]
	}
	if cont {
		// add the ending parameters
		fmt.Println(status)
		fmt.Println(getLogo.BankLogo())
		spacingToTheExit("", 3)
		fmt.Printf("-------   WELCOME TO THE BANKING APPLICATIONS Mr. %s    -------    ", strings.ToUpper(customerGlobalScope["fullname"]))
		fmt.Printf("-  %s  -    ", time.Now().Format(time.RFC850))
		spacingToTheExit("", 3)
		fmt.Println("HINT -> TYPE NUMBERS ASSOCIATED WITH THE MODULES MENTIONED BELOW")
		spacingToTheExit("", 2)
		fmt.Println("|-----------------------------------------------|")
		fmt.Printf(" Customer Name  |        %s          \n|", customerGlobalScope["fullname"])
		fmt.Println("-----------------------------------------------|")
		fmt.Printf(" Address        |        %s          \n|", customerGlobalScope["address"])
		fmt.Println("-----------------------------------------------|")
		fmt.Printf(" Phone Number   |        %s          \n|", customerGlobalScope["phone"])
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
			bankingModules(int, "", customerGlobalScope)
		case 5:
			if true {
				bankingModules(int, "")
			} else {
				welcomeloop(true, src.CHECKCREDS, false)
			}

		case 9:
			clearTheTerminal(src.CLEARTERMINAL)
			spacingToTheExit(".", 4)
			println(src.GOODBYENOTE)
			spacingToTheExit(".", 4)
			cont = false
		default:
			welcomeloop(true, src.CHECKCREDS, false)
		}
	} else {
		// add a ending paramter
		spacingToTheExit(".", 4)
		println(src.GOODBYENOTE)
		spacingToTheExit(".", 4)
	}

}

// the main login page of the application
func loginPage() {
	clearTheTerminal(src.CLEARTERMINAL)
	println("ENTER THE RIGHT CREDENTIALS TO ACCESS THE BANKING APPLICATION")
	spacingToTheExit("", 2)
	print("USERNAME - ")
	username, _ := takeTheUserInput("str")
	print("PASSWORD - ")
	_, password := takeTheUserInput("int")
	successFlag := src.TakeTheUserCreds(username, password)
	if successFlag {
		// after the login gets successful
		welcomeloop(true, "", true)
	} else {

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

// modules for the bank application
func bankingModules(head int, blockStat string, custinf ...map[string]string) {
	updateTheTable := false
	if head == 2 {
		updateTheTable = true
	}
	clearTheTerminal(src.CLEARTERMINAL)
	if blockStat == "blocked" {
		fmt.Println(src.CHECKCREDS)
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
		// to check the main balance of the user
		src.CheckBalance(custinf[0])

	} else if head == 2 {
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
		welcomeloop(checkStat, Status, updateTheTable, custinf...)
	}
}

// this function gives the exiting text
func exitTextSignal(currentInt int) (bool, string) {
	spacingToTheExit("", 4)
	println(src.EXITAPP)
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

//this function gives spacing
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

// This function clears the terminal and prints the designated text
func clearTheTerminal(s string) bool {
	forceClear := func() {
		fmt.Print("\033[H\033[2J")
	}
	if s == src.CLEARTERMINAL {
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
