package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)
import "io/ioutil"

type Actor struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Character string `json:"character"`
	Seasons int `json:"seasons"`
	Alive bool `json:"alive"`
}

type JsonStructure struct {
	Show string `json:"show"`
	Director string `json:"director"`
	Actors []Actor `json:"actors"`

}

type User struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}


func main() {
	data, err := readjson()
	if err != nil {
		log.Fatal(err)
	}

	db,err := dbconnection()
	if err !=nil {
		log.Fatal(err)
	}
	inserttodb(db, data)
	readfromdb(db)
}
	func dbconnection() (connection *sql.DB, err error) {
		dburi:= "root:daddy@tcp(localhost:3306)/assignment"

		connection, _ = sql.Open("mysql", dburi)
		err = connection.Ping()
		if err !=nil {
			fmt.Println(err)
			return
		}
		return
	}

	func readjson () (data JsonStructure, err error){
		file, err := ioutil.ReadFile("./users.json")
		if err != nil {
			return
		}

		if err := json.Unmarshal(file, &data);  err!=nil {
			return
		}
		return
	}

	func inserttodb (dbconnection *sql.DB, data JsonStructure) {
		for _,actor := range data.Actors{
			if actor.Alive{
				fmt.Println(actor.FirstName, actor.LastName)
				query := "insert into user (first_name, last_name) values (?,?)"
				_, err:= dbconnection.Exec(query, actor.FirstName, actor.LastName)
				if err !=nil{
					fmt.Println("unable to insert because", err)
				}
			}
		}
	}


	func readfromdb (db *sql.DB) {
		// reading from the database
		var users []User
		query:= "select distinct (first_name), last_name from user order by first_name asc"
		rows, err := db.Query(query)
		if err != nil{
			log.Fatal(err)
		}

		for rows.Next(){
			var user User
			if err := rows.Scan(&user.FirstName, &user.LastName); err != nil{
				fmt.Println("Unable to read data", err)
				continue
			}
			users= append(users, user)
		}
		fmt.Println(users)
	}