package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	var res *sql.Rows
	if res, err = db.Query("select type.a from type"); err != nil {
		log.Println(err)
	}
	// type.a smallint 256
	var a byte
	var b int
	for res.Next() {
		err = res.Scan(&a)
		log.Println(err)
		log.Println(a)
		res.Scan(&b)
		log.Println(b)
	}
	// orm
	gormdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	gormdb.Table("type").Select("type.a").Find(&a)
	log.Println(a)
	gormdb.Table("type").Select("type.a").Find(&b)
	log.Println(b)
}
