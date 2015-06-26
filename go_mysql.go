package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	//insert
	stmt, err := db.Prepare("insert tb_tag set name=?,count=?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("tetst", 2)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	fmt.Println(id)

	//delete
	stmt, err = db.Prepare("delete from tb_tag where id=?")
	if err != nil {
		log.Fatal(err)
	}
	res, err = stmt.Exec(1)
	if err != nil {
		log.Fatal(err)
	}
	getA, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(getA)
	db.Close()

}
