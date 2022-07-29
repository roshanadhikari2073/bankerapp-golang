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
	flag, salt := sqlconn.VerifyTheCredentials(un, "pass")
	if flag && len(salt) != 0 {
		t := strconv.Itoa(pass)
		err = e.Compare(salt, t)
		if err == nil {
			return true
		}
	}
	return false
}

// this function takes the username and verifies if the user exists
func VerifyTheUserName(un string) bool {
	flag, _ := sqlconn.VerifyTheCredentials(un, "user")
	return flag
}

func CreateNewAccount() string {
	var status string
	user_creds := []struct {
		input string
		creds string
		i     int
	}{
		{"ENTER YOUR USERNAME :", "username", 10},
		{"ENTER YOUR 4 DIGIT PIN :", "password", 4},
		{"ENTER TOTAL AMOUNT YOU WANT TO DEPOSIT :", "total_balance", 6},
		{"ENTER YOUR ADDRESS :", "address", 10},
		{"ENTER YOUR PHONE :", "phone", 9},
	}
	user_informations := make(map[string]string)
	for index, element := range user_creds {
		temp := user_creds[index]
		fmt.Println(temp.input)
		username := char_limiter(temp.creds, temp.i)
		user_informations[element.creds] = username
	}
	if len(user_informations) > 0 {
		flag, _ := sqlconn.VerifyTheCredentials(user_informations["username"], "user")
		if flag {
			return "user_exists"
		} else {
			status = sqlconn.CreateBankAccount(user_informations)
			return status
		}
	}
	return status
}

// this function scans user input and limit
func char_limiter(s string, limit int) string {
	text := ""
	// create a hash value of the password
	if s == "password" {
		fmt.Println("\033[8m") // Hide input
		fmt.Scanf("%s", &text)
		fmt.Println("\033[28m")
		if len(text) < 4 {
			fmt.Print("PLEASE ENTER ATLEAST FOUR DIGIT PIN:")
			char_limiter(s, limit)
		}
		hashed := Hash{}
		hashedval, _ := hashed.Generate(text)
		return hashedval
	} else {
		fmt.Scanf("%s", &text)
	}

	if len(text) > limit || len(text) == 0 {
		fmt.Printf("PLEASE ENTER CHARACTERS OF SIZE %d .. PLEASE RE ENTER-%s \n", limit, s)
		char_limiter(s, limit)
	}
	if limit == 4 || limit == 6 || limit == 9 {
		if i, err := strconv.Atoi(text); err != nil {
			if i == 0 {
				fmt.Printf("YOU CANNOT ENTER STRING PLEASE RE ENTER-%s \n", s)
				char_limiter(s, limit)
			}
		}
	}
	return text
}
