package main

import (
	"database/sql"
	"fmt"

	"github.com/cznic/ql"
	_ "github.com/cznic/ql/driver"
	_ "github.com/mattn/go-sqlite3"
)

var (
	qlDB *sql.DB
	//sqlDB *sql.DB
)

func createQLTable(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Begin:", err)
	}
	_, err = tx.Exec(`
	CREATE TABLE ip2location_db5 (
		ip_from int64,
		ip_to int64,
		latitude float64,
		longitude float64);
	`)
	// CREATE TABLE ip2location_db5(ip_from INTEGER,ip_to INTEGER,country_code TEXT,country_name TEXT,region_name TEXT,city_name TEXT,latitude DOUBLE,longitude DOUBLE)
	if err != nil {
		fmt.Println("Exec:", err)
	}
	if err = tx.Commit(); err != nil {
		fmt.Println("Commit:", err)
	}
	return nil
}

func openQLDB() (*sql.DB, error) {
	ql.RegisterDriver()
	return sql.Open("ql-mem", "mem.db")
}

func openSQLite3DB() (*sql.DB, error) {
	return sql.Open("sqlite3", "geo.sqlite3")
}

func importSQLite2QL(from, to *sql.DB) {
	rows, err := from.Query("SELECT * FROM ip2location_db5")
	if err != nil {
		fmt.Println(err)
		return
	}

	var line uint32
	for rows.Next() {
		var ip_from, ip_to uint32
		var country_code string
		var country_name string
		var region_name string
		var city_name string
		var lat, lon float64

		err := rows.Scan(
			&ip_from, &ip_to,
			&country_code, &country_name, &region_name, &city_name,
			&lat, &lon)
		if err != nil {
			fmt.Println(err)
			continue
		}

		s := fmt.Sprintf(
			`INSERT INTO ip2location_db5 VALUES (%v, %v, %v, %v);`,
			ip_from, ip_to,
			lat, lon)
		if line == 0 {
			line++
			continue
		}
		tx, err := to.Begin()
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = tx.Exec(s)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if err := tx.Commit(); err != nil {
			fmt.Println(err)
			continue
		}
		line++
		fmt.Println(line)
	}
}

func main() {
	var err error
	qlDB, err = openQLDB()
	if err != nil {
		fmt.Println("openQLDB:", err)
		return
	}
	if err := createQLTable(qlDB); err != nil {
		fmt.Println("createQLTable:", err)
		return
	}
	sqlDB, err = openSQLite3DB()
	if err != nil {
		fmt.Println("openSQLite3DB:", err)
		return
	}
	importSQLite2QL(sqlDB, qlDB)

	defer func() {
		if err := qlDB.Close(); err != nil {
			return
		}

		fmt.Println("OK")
	}()

	select {}
}
