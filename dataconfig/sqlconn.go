package sqlconn

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type customerInfo struct {
	Name          string
	User_type     string
	Address       string
	Phone         int
	Total_balance int
	Total_loan    int
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "1234"
	dbName := "shardINDEX"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Show(username string) map[string]string {
	db := dbConn()
	selDB, err := db.Query("SELECT id, total_balance, total_loan, phone, username, user_type, address FROM user WHERE username=?", username)
	if err != nil {
		panic(err.Error())
	}
	emp := customerInfo{}
	for selDB.Next() {
		var id, total_balance, total_loan, phone int
		var name, user_type, address string
		err = selDB.Scan(&id, &total_balance, &total_loan, &phone, &name, &user_type, &address)
		if err != nil {
			panic(err.Error())
		}
		emp.Name = name
		emp.User_type = user_type
		emp.Address = address
		emp.Phone = phone
		emp.Total_balance = total_balance
		emp.Total_loan = total_loan
	}
	defer db.Close()
	return map[string]string{

		"fullname":     emp.Name,
		"usertype":     emp.User_type,
		"address":      emp.Address,
		"phone":        strconv.Itoa(emp.Phone),
		"totalbalance": strconv.Itoa(emp.Total_balance),
		"totalloan":    strconv.Itoa(emp.Total_loan),
	}
}

func CreateBankAccount() {
	db := dbConn()
	insForm, err := db.Prepare("INSERT INTO user(username, password, total_balance, address, phone) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec("name", "pass")
	defer db.Close()
}

func VerifyTheCredentials(username string) (bool, string) {
	var salt string
	db := dbConn()
	err := db.QueryRow("SELECT password FROM user WHERE username=?", username).Scan(&salt)
	if err != nil {
		panic(err.Error())
	}
	db.Close()

	if len(salt) != 0 {
		return true, salt
	}

	return false, salt

}

// func Update(w http.ResponseWriter, r *http.Request) {
// 	db := dbConn()
// 	if r.Method == "POST" {
// 		name := r.FormValue("name")
// 		city := r.FormValue("city")
// 		id := r.FormValue("uid")
// 		insForm, err := db.Prepare("UPDATE Employee SET name=?, city=? WHERE id=?")
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		insForm.Exec(name, city, id)
// 		log.Println("UPDATE: Name: " + name + " | City: " + city)
// 	}
// 	defer db.Close()
// }

// func Delete(w http.ResponseWriter, r *http.Request) {
// 	db := dbConn()
// 	emp := r.URL.Query().Get("id")
// 	delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	delForm.Exec(emp)
// 	log.Println("DELETE")
// 	defer db.Close()
// }

// func main() {
// 	log.Println("establishing MySQL database connection")
// }
