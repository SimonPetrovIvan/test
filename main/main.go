package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type User struct {
	name                 string
	age                  uint16
	money                int16
	avgGrades, happiness float64
}

var database *sql.DB

func (u User) getAllinfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money egual: %d", u.name, u.age, u.money)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	bob := User{name: "Simon", age: 25, money: -1000, avgGrades: 5.5, happiness: 10.0}

	fmt.Fprintf(w, bob.getAllinfo())
}
func contacts(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	rows, err := database.Query("select * from actor where actor_id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	var (
		actorId             int64
		firstName, lastName string
		lastUpdate          string
	)
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&actorId, &firstName, &lastName, &lastUpdate); err != nil {
			log.Fatal(err)
		}
	}

	w.Write([]byte(fmt.Sprintf("actor id: %d, name: %s", actorId, firstName)))
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contact/", contacts)
	http.ListenAndServe("localhost:8080", nil)

}

func main() {
	db, err := sql.Open("mysql", "Simon:Aza20141202@tcp(localhost:3306)/sakila")

	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	database = db
	handleRequest()

}
