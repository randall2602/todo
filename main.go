// Sample datastore-quickstart fetches an entity from Google Cloud Datastore.
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type database struct {
	conn *sql.DB
	user string
	pass string
	host string
	port string
	db   string
}

func (db *database) Init(user, pass, host, port, use string) error {
	var err error
	driver := "mysql"
	protocol := "tcp"
	connString := user + ":" + pass + "@" + protocol + "(" + host + ":" + port + ")/" + use
	db.conn, err = sql.Open(driver, connString)
	return err
}

func main() {
	// set up database
	var db database
	err := db.Init("app", "NWlmeA27jk", "127.0.0.1", "3306", "todo")
	if err != nil {
		log.Fatal(err)
	}
	// run a query
	var (
		name        string
		description string
	)
	rows, err := db.conn.Query("select name, description from tasks_summary where completed=?", false)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name, &description)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("list: ", name, "\ttask: ", description)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer db.conn.Close()
}
