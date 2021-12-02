// author: roshanadhikai
// this page immplements the main start page of the application that initiaties the login, signup and account termination

package src

import (
	sqlconn "cliapplications/dataconfig"
	"fmt"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

//Hash implements root.Hash
type Hash struct{}

// Adding information to create account

var userinp string = "ENTER YOUR USERNAME :"
var passinp string = "ENTER YOUR 4 DIGIT PIN :"
var totalbalinp string = "ENTER TOTAL AMOUNT YOU WANT TO DEPOSIT :"
var addressinp string = "ENTER YOUR ADDRESS :"
var phoneinp string = "ENTER YOUR PHONE :"

// setting up the hashing and the user login mechanism

func Check(pass string) {
	hashed := Hash{}
	hashedval, _ := hashed.Generate(pass)
	check := hashed.Compare(hashedval, pass)
	print(check)
}

//Generate a salted hash for the input string
func (c *Hash) Generate(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

//Compare string to generated hash
func (c *Hash) Compare(hash string, s string) error {
	incoming := []byte(s)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}

// this function takes the credentials and compares within db and pass
func TakeTheUserCreds(un string, pass int) bool {
	var err error
	e := Hash{}
	flag, salt := sqlconn.VerifyTheCredentials(un)
	if flag && len(salt) != 0 {
		t := strconv.Itoa(pass)
		err = e.Compare(salt, t)
		if err == nil {
			return true
		}
	}
	return false
}

func CreateNewAccount() string {
	user_informations := make(map[string]string)
	// inserting values in idiotmatic way needs revision
	fmt.Println(userinp)
	username := char_limiter(userinp, 10)
	user_informations["username"] = username

	fmt.Println(passinp)
	password := char_limiter(passinp, 4)
	user_informations["password"] = password

	fmt.Println(totalbalinp)
	total_balance := char_limiter(totalbalinp, 6)
	user_informations["total_balance"] = total_balance

	fmt.Println(addressinp)
	address := char_limiter(addressinp, 10)
	user_informations["address"] = address

	fmt.Println(phoneinp)
	phone := char_limiter(phoneinp, 9)
	user_informations["phone"] = phone

	println(len(user_informations))
	return "success"
}

func char_limiter(s string, limit int) string {
	text := ""
	fmt.Scanf("%s", &text)
	if len(text) > limit {
		fmt.Printf("You cannot enter character more than %d .. Please Re-Enter %s", limit, s)
		char_limiter(s, limit)
	}
	return text

}
