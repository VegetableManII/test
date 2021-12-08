package questions

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SQLUpdateGorm() {
	gormdb, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:@tcp(127.0.0.1:3306)/rmp?charset=utf8",
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// 插入
	// 在建表SQL中指定的 default 不会起作用，如果对象某些字段未赋值则表中该字段会初始化为零值
	// 插入失败（unique出现重复）会浪费一个自增id
	e := gormdb.Create(&Rules{
		BusinessID:    "3",
		ServiceSource: "tp",
		RuleName:      "haha",
		Rules:         "rules",
	}).Error
	if e != nil {
		log.Println(e)
	}
	// 更新
	// 当以主键更新时，只需要在数据对象中包含主键字段即可；以主键更新时即使主键不存在也不会发生错误
	// 如果以非主键进行更新则必须使用Where子句来指定
	r := &Rules{
		Id:            3,
		BusinessID:    "5",
		ServiceSource: "tp",
		RuleName:      "haha",
		Rules:         "updated rules",
	}
	e = gormdb.Model(r).Updates(r).Error
	if e != nil {
		log.Println(e)
	}
	// 当以不存在的主键进行更新时不会返回错误
	// e = gormdb.Model(&Rules{Id: 6}).Updates(Rules{RuleName: "name"}).Error
	// if e != nil {
	// 	log.Println(e)
	// }
	// 删除

	// 查询
	// res := make([]*Rules, 0, 1)
	// err = gormdb.Find(&res, "date(ctime) >= '1000-01-01' and date(ctime) < '1000-01-02'").Where("business_id=?", 1).Error // where不起作用
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(res)
}
