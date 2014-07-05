package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

)


func main() {
	db, err := sql.Open("mysql", "dz:123456@tcp(192.168.11.110:3306)/xweibo11_67?charset=utf8")
	checkErr(err)
	defer db.Close()

	//查询数据
	rows, err := db.Query("SELECT * FROM xwb11_sys_config")
	checkErr(err)

	for rows.Next() {
		var key string
		var value string
		var group_id int

		err = rows.Scan(&key, &value, &group_id)
		checkErr(err)
		fmt.Println(key)
		fmt.Println(value)
		fmt.Println(group_id)

	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

