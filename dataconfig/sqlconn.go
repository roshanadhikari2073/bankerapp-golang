package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type customerInfo struct {
	name          string
	user_type     string
	address       string
	phone         int
	total_balance int
	total_loan    int
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

func Show(username string) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Employee WHERE username=?", username)
	if err != nil {
		panic(err.Error())
	}
	emp := customerInfo{}
	for selDB.Next() {
		var id, total_balance, total_loan, phone int
		var name, city, user_type, address string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.name = name
		emp.user_type = user_type
		emp.address = address
		emp.phone = phone
		emp.total_balance = total_balance
		emp.total_loan = total_loan
	}
	defer db.Close()
}

// func Insert(w http.ResponseWriter, r *http.Request) {
// 	db := dbConn()
// 	if r.Method == "POST" {
// 		name := r.FormValue("name")
// 		city := r.FormValue("city")
// 		insForm, err := db.Prepare("INSERT INTO Employee(name, city) VALUES(?,?)")
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		insForm.Exec(name, city)
// 		log.Println("INSERT: Name: " + name + " | City: " + city)
// 	}
// 	defer db.Close()
// }

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

func main() {
	log.Println("establishing MySQL database connection")
}
