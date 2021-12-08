package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	Fname  string
	Lname  string
	City   string
	Mobile int64
}

func main() {
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
	fmt.Println(user_creds[0])
	for index, element := range user_creds {
		temp := user_creds[index]
		fmt.Println(temp.input)
		username := char_limiter(temp.creds, temp.i)
		user_informations[element.creds] = username
	}
	fmt.Println(user_informations["username"])

}

// this function scans user input and limit
func char_limiter(s string, limit int) string {
	text := ""
	fmt.Scanf("%s", &text)
	if len(text) > limit {
		fmt.Printf("YOU CANNOT ENTER CHARACTERS MORE THAN %d .. PLEASE RE ENTER-%s \n", limit, s)
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
