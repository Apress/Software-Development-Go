package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"math/rand"
	"strconv"
)

type Record struct {
	Id   string
	Name string
}

var (
	dbHandle *sql.DB
	curNames = [5]string{"USD", "CAD", "AUD", "EUR", "GBP"}
	dbname   = "local.db"
)

func main() {
	log.Println("Initialize database - ", dbname)
	dbHandle = InitDB(dbname)

	log.Println("Creating table in -  ", dbname)
	InitTable(dbHandle)

	log.Println("Inserting data into - ", dbname)
	records := []Record{}

	for i := 0; i < 25; i++ {
		r := (rand.Intn(len(curNames)-0) + 0)
		d := strconv.Itoa(i)
		rec := Record{Id: d, Name: curNames[r]}
		records = append(records, rec)
	}

	InsertData(dbHandle, records)

	fmt.Println("Reading Table: ")
	n := ReadData(dbHandle)

	log.Println("Total rows read - ", n)
}

// InitDB initialize database
func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	return db
}

// InitTable initialize table
func InitTable(db *sql.DB) {
	q := `
	CREATE TABLE IF NOT EXISTS currencies(
		Id TEXT NOT NULL PRIMARY KEY,
		Name TEXT,
		InsertedDatetime DATETIME
	);`

	_, err := db.Exec(q)
	if err != nil {
		log.Fatal(err)
	}
}

// InsertData inserting data into table
func InsertData(db *sql.DB, records []Record) {
	q := `
	INSERT OR REPLACE INTO currencies(
		Id,
		Name,
		InsertedDatetime
	) values(?, ?,  CURRENT_TIMESTAMP)`

	stmt, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, r := range records {
		_, err := stmt.Exec(r.Id, r.Name)
		if err != nil {
			log.Fatal(err)
		}
	}
}

//ReadData reading data from table
func ReadData(db *sql.DB) []Record {
	q := `
	SELECT Id, Name  FROM currencies
	ORDER BY datetime(InsertedDatetime) DESC`

	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var records []Record
	for rows.Next() {
		item := Record{}
		err := rows.Scan(&item.Id, &item.Name)
		if err != nil {
			panic(err)
		}
		records = append(records, item)
	}
	return records
}
