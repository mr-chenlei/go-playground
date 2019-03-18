package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/boltdb/bolt"
	sqlite3 "github.com/mattn/go-sqlite3"
)

const (
	dbFileName = "geo.sqlite3"
)

var (
	sqlDB  *sql.DB
	boltDB *bolt.DB
)

// CreateDBInst 获取数据库实例
// 建立内存数据库，然后复制文件数据库中的内容到内存数据库中，以后通过内存数据库访问数据
// 参数：
//      无
// 返回值：
//      error: 出错信息
func createSQLDBInst() error {
	if _, err := os.Stat(dbFileName); err != nil {
		return err
	}
	var err error
	sqlite3conn := []*sqlite3.SQLiteConn{}
	sql.Register("sqlite3_with_hook",
		&sqlite3.SQLiteDriver{
			ConnectHook: func(conn *sqlite3.SQLiteConn) error {
				sqlite3conn = append(sqlite3conn, conn)
				return nil
			},
		})
	fileDB, err := sql.Open("sqlite3_with_hook", dbFileName)
	if err != nil {
		return err
	}
	defer fileDB.Close()
	fileDB.Ping()
	sqlDB, err = sql.Open("sqlite3_with_hook", "file::memory:?mode=memory&cache=shared")
	if err != nil {
		return err
	}
	sqlDB.Ping()
	bk, err := sqlite3conn[1].Backup("main", sqlite3conn[0], "main")
	if err != nil {
		return err
	}
	_, err = bk.Step(-1)
	if err != nil {
		return err
	}
	err = bk.Finish()
	if err != nil {
		return err
	}
	return nil
}

func createBoltDBInst() error {
	var err error
	boltDB, err = bolt.Open("ip-geo.db", 0600, nil) //, &bolt.Options{ReadOnly: true})
	if err != nil {
		return err
	}
	return nil
}

func ipInt2String(ip uint32) string {
	return net.IP{
		byte((ip & 0xFF000000) >> 24),
		byte((ip & 0x00FF0000) >> 16),
		byte((ip & 0x0000FF00) >> 8),
		byte(ip & 0x000000FF)}.String()
}

func writeBoltDB(i, j uint32, geo string) {
	// Start a writable transaction.
	tx, err := boltDB.Begin(true)
	if err != nil {
		fmt.Println("Begin:", err)
		return
	}
	// Use the transaction...
	b, err := tx.CreateBucketIfNotExists([]byte("ip2location_ipv4_db5"))
	if err != nil {
		fmt.Println(err)
		return
	}
	b.FillPercent = 0.9
	for ; i < j; i++ {
		ip := ipInt2String(i)
		result := b.Get([]byte(ip))
		if result != nil {
			continue
		}
		if err := b.Put([]byte(ip), []byte(geo)); err != nil {
			fmt.Println("Put:", err)
		}
	}
	if err := tx.Commit(); err != nil {
		fmt.Println("Commit:", err)
		return
	}
}

func main() {
	if err := createSQLDBInst(); err != nil {
		fmt.Println(err)
		return
	}
	if err := createBoltDBInst(); err != nil {
		fmt.Println(err)
		return
	}

	rows, err := sqlDB.Query("SELECT * FROM ip2location_db5")
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
			fmt.Println("Scan:", err)
			continue
		}

		geo := strconv.FormatFloat(lat, 'f', 3, 32) + ":" + strconv.FormatFloat(lon, 'f', 3, 32)
		writeBoltDB(ip_from, ip_to, geo)
		fmt.Printf("line: %d, %d -> %d, %s\n", line, ip_from, ip_to, geo)
		line++
	}
}
