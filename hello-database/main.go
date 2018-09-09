package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Id         int64
	Name       sql.NullString
	Price      int
	Active     bool
	ModifiedAt time.Time
}

func main() {

	printSection("SQL Drivers")
	fmt.Println(sql.Drivers())
	fmt.Println()

	driverName := "sqlite3"
	var dbSourceName string
	if len(os.Args) > 1 {
		dbSourceName = os.Args[1]
	} else {
		dbSourceName = ":memory:"
	}

	var err error
	var db *sql.DB

	db, err = sql.Open(driverName, dbSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	printSection("Ping")
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	printSection("Create tables")
	_, err = db.Exec(`
		CREATE TABLE products (
        	id INTEGER NOT NULL PRIMARY KEY,
        	name TEXT,
        	price INTEGER NOT NULL DEFAULT 0,
        	active INT(1) NOT NULL DEFAULT 0,
        	modified_at DATETIME
		)`)
	if err != nil {
		log.Fatal(err)
	}

	printSection("Insert rows")
	products := []Product{
		Product{1, nullableString("Product1"), 1230, true, time.Now()},
		Product{2, nullableString("Product2"), 2340, true, time.Now()},
		Product{3, nullableString("Product3"), 3450, false, time.Now()},
		Product{4, nullString(), 4560, true, time.Now()},
	}
	var tx *sql.Tx
	tx, err = db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	for _, x := range products {
		r, err := tx.Exec("INSERT INTO products VALUES ($1, $2, $3, $4, $5)",
			x.Id, x.Name, x.Price, x.Active, x.ModifiedAt)
		if err != nil {
			log.Fatal(err)
		}
		n, err := r.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("rows affefected: ", n)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	printSection("Select rows")
	selectedRows, err := db.Query("SELECT * FROM products WHERE active = $1", true)
	if err != nil {
		log.Fatal(err)
	}
	defer selectedRows.Close()

	selectedColumnTypes, err := selectedRows.ColumnTypes()
	if err != nil {
		log.Fatal(err)
	}
	for _, x := range selectedColumnTypes {
		fmt.Println(describeColumnType(x))
	}
	fmt.Println()

	for selectedRows.Next() {
		row := Product{}
		if err := selectedRows.Scan(&row.Id, &row.Name, &row.Price,
			&row.Active, &row.ModifiedAt); err != nil {
			log.Fatal(err)
		}
		fmt.Println(row)
	}

	printSection("Select rows using prepared statements")
	var findByIdStmt *sql.Stmt
	findByIdStmt, err = db.Prepare("SELECT * FROM products WHERE id = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer findByIdStmt.Close()
	for _, x := range products {
		rows, err := findByIdStmt.Query(x.Id)
		if err != nil {
			log.Fatal(err)
		}
		if rows.Next() {
			row := Product{}
			if err := rows.Scan(&row.Id, &row.Name, &row.Price,
				&row.Active, &row.ModifiedAt); err != nil {
				log.Fatal(err)
			}
			fmt.Println(row)
		}
		rows.Close()
	}
}

func nullableString(v string) sql.NullString {
	return sql.NullString{String: v, Valid: true}
}

func nullString() sql.NullString {
	return sql.NullString{}
}

func describeColumnType(ct *sql.ColumnType) string {
	var s string
	if presision, scale, ok := ct.DecimalSize(); ok {
		s = fmt.Sprintf("%d,%d", presision, scale)
	} else if length, ok := ct.Length(); ok {
		s = fmt.Sprintf("%d", length)
	} else {
		s = ""
	}
	return fmt.Sprintf("%s %s %s", ct.Name(), ct.DatabaseTypeName(), s)
}

func printSection(s string) {
	fmt.Println()
	fmt.Println()
	fmt.Println(strings.Title(s))
	fmt.Println(strings.Repeat("=", len(s)))
	fmt.Println()
}
