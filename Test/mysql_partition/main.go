package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Record struct {
	Id    int64 `gorm:"column:id"`
	Ctime int   `gorm:"ctime"`
}

func (Record) TableName() string {
	return "t1"
}

func main() {
	db, err := sql.Open("mysql", "root:jiaoxueming@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	// orm
	gormdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	record1 := Record{
		Ctime: 11,
	}
	record2 := Record{
		Ctime: 1011,
	}
	gormdb.Table("p0").Create(&record1)
	gormdb.Table("p1").Create(&record2)
}
