// author - roshan adhikari
// desc - this serves as db connection and mysql is used as the driver for the db, some parts are commented and needs refinement

package sqlconn

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type customerInfo struct {
	Name          string
	User_type     string
	Address       string
	Phone         int
	Total_balance int
	Total_loan    string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Password#123"
	dbName := "shardINDEX"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err)
	} else if err = db.Ping(); err != nil {
		panic(err)
	}
	return
}
func UserInfo(username string) map[string]string {
	db := dbConn()
	selDB, err := db.Query("SELECT id, total_balance, total_loan, phone, username, user_type, address FROM user WHERE username=?", username)
	if err != nil {
		panic(err.Error())
	}
	emp := customerInfo{}
	for selDB.Next() {
		var total_loan, user_type sql.NullString
		var id, total_balance, phone int
		var name, address string
		err = selDB.Scan(&id, &total_balance, &total_loan, &phone, &name, &user_type, &address)
		if err != nil {
			panic(err.Error())
		}
		emp.Name = name
		emp.User_type = user_type.String
		emp.Address = address
		emp.Phone = phone
		emp.Total_balance = total_balance
		emp.Total_loan = total_loan.String
	}
	defer db.Close()
	return map[string]string{

		"fullname":     emp.Name,
		"usertype":     emp.User_type,
		"address":      emp.Address,
		"phone":        strconv.Itoa(emp.Phone),
		"totalbalance": strconv.Itoa(emp.Total_balance),
		"totalloan":    emp.Total_loan,
	}
}

func CreateBankAccount(s map[string]string) string {
	db := dbConn()
	query := "INSERT INTO user(username, password, total_balance, address, phone) VALUES(?,?,?,?,?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		fmt.Printf("Error %s when preparing SQL statement", err)
		return "error"
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, s["username"], s["password"], s["total_balance"], s["address"], s["phone"])
	if err != nil {
		return "error"
	}
	_, err = res.RowsAffected()
	if err != nil {
		return "error"
	}
	return "success"
}

func VerifyTheCredentials(username string, typeOfCreds string) (bool, string) {
	var salt string
	db := dbConn()
	if typeOfCreds == "pass" {
		err := db.QueryRow("SELECT password FROM user WHERE username=?", username).Scan(&salt)
		if err != nil {
			println("ERROR !")
		}
	} else if typeOfCreds == "user" {
		err := db.QueryRow("SELECT username FROM user WHERE username=?", username).Scan(&salt)
		if err != nil {
			println("ERROR !")
		}
	}
	db.Close()
	if len(salt) != 0 {
		return true, salt
	}
	return false, salt
}

// TO DO - make some changes to the bank modules

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE Employee SET name=?, city=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, city, id)
		log.Println("UPDATE: Name: " + name + " | City: " + city)
	}
	defer db.Close()
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
}

// func main() {
// 	log.Println("establishing MySQL database connection")
// }

// // CREATE TABLE IF NOT EXISTS user (
//     id INT AUTO_INCREMENT PRIMARY KEY,
//     username VARCHAR(255) NOT NULL,
// 	password VARCHAR(255) NOT NULL,
// 	total_balance BIGINT,
// 	total_loan BIGINT,
// 	phone BIGINT,
// 	user_type VARCHAR(255),
// 	address VARCHAR(255),
//     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// )  ENGINE=INNODB;

// selDB, err := db.Query("SELECT id, password, total_balance, total_loan, phone, username, user_type, address FROM user WHERE username=?", username)
