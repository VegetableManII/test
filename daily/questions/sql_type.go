package questions

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type rules struct {
	Id            int64  `gorm:"column:id"`
	Appid         int    `gorm:"column:app_id"`
	RuleName      string `gorm:"column:rule_name"`
	ServiceSource int    `gorm:"column:srv_source"`
	Rules         string `gorm:"column:rules"`
}

func SQLType() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	// var res *sql.Rows
	// if res, err = db.Query("select type.a from type"); err != nil {
	// 	log.Println(err)
	// }
	// // type.a smallint 256
	// var a byte
	// var b int
	// for res.Next() {
	// 	err = res.Scan(&a)
	// 	log.Println(err)
	// 	log.Println(a)
	// 	res.Scan(&b)
	// 	log.Println(b)
	// }
	// orm
	gormdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// 插入
	// 在建表SQL中指定的 default 不会起作用，如果对象某些字段未赋值则表中该字段会初始化为零值
	// 插入失败（unique出现重复）会浪费一个自增id
	// e := gormdb.Create(&rules{
	// 	Appid:    3,
	// 	RuleName: "haha",
	// 	Rules:    "rules",
	// }).Error
	// if e != nil {
	// 	log.Println(e)
	// }
	// 更新
	// 当以主键更新时，只需要在数据对象中包含主键字段即可；以主键更新时即使主键不存在也不会发生错误
	// 如果以非主键进行更新则必须使用Where子句来指定
	e := gormdb.Model(&rules{Id: 5}).Update("rules", "updated data").Error
	if e != nil {
		log.Println(e)
	}
	// 当以不存在的主键进行更新时不会返回错误
	e = gormdb.Model(&rules{Id: 3}).Updates(rules{Id: 3, RuleName: "new_name"}).Error
	if e != nil {
		log.Println(e)
	}
	// 删除

	// 查询

}
