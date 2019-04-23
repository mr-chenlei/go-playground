package main

import (
	"database/sql"
	"fmt"

	"github.com/cznic/ql"
	_ "github.com/cznic/ql/driver"
)

func main() {
	ql.RegisterDriver()
	db, err := sql.Open("ql", "ip-geo-.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	s := fmt.Sprintf(`SELECT * FROM ip2location_db5 WHERE ip_from <= %v and ip_to >=%v`, 17092096, 17092096)
	rows, err := db.Query(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	var line uint32
	for rows.Next() {
		var ip_from, ip_to uint32
		var lat, lon float64

		err := rows.Scan(&ip_from, &ip_to, &lat, &lon)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(line, ip_from, ip_to, lat, lon)
		line++
	}
}
