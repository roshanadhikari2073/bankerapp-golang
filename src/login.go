// author: roshanadhikai
// this page immplements the main start page of the application that initiaties the login, signup and account termination

package src

import (
	sqlconn "cliapplications/dataconfig"
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

// func CreateBankAccount() {

// }

// func Login() {

// }

// func Terminate() {

// }
