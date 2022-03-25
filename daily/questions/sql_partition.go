package questions

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
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

func SQLPartation() {
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

type Rules struct {
	Id            int64  `gorm:"column:id"`
	Group         string `gorm:"column:group"`
	RuleName      string `gorm:"column:rule_name"`
	BusinessID    string `gorm:"business_id"`
	ServiceSource string `gorm:"column:source"`
	Rules         string `gorm:"column:rules"`
}

func SQLType() {
	// db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/rmp?charset=utf8")
	// if err != nil {
	// 	log.Fatal(err)
	// }
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
		DSN: "root:@tcp(127.0.0.1:3306)/rmp?charset=utf8",
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
	// e := gormdb.Model(&rules{Id: 5}).Update("rules", "updated data").Error
	// if e != nil {
	// 	log.Println(e)
	// }
	// // 当以不存在的主键进行更新时不会返回错误
	// e = gormdb.Model(&rules{Id: 3}).Updates(rules{Id: 3, RuleName: "new_name"}).Error
	// if e != nil {
	// 	log.Println(e)
	// }
	// 删除

	// 查询
	res := make([]*Rules, 0, 1)
	err = gormdb.Find(&res, "date(ctime) >= '1000-01-01' and date(ctime) < '1000-01-02'").Where("business_id=?", 1).Error // where不起作用
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
}

type Users struct {
	ID          int64     `xorm:"id"`
	IMSI        string    `xorm:"imsi"`
	RootK       string    `xorm:"root_k"`
	Opc         string    `xorm:"opc"`
	Mnc         string    `xorm:"mnc"`
	Mcc         int       `xorm:"mcc"`
	Apn         string    `xorm:"apn"`
	IP          string    `xorm:"ip"`
	SipUsername string    `xorm:"sip_username"`
	DNS         string    `xorm:"sip_dns"`
	Ctime       time.Time `xorm:"ctime"` // 记录创建时刻
	Utime       time.Time `xorm:"utime"` // 记录更新时刻
}

func XormGet() {
	engin, err := xorm.NewEngine("mysql", "root:@tcp(127.0.0.1:3306)/volte?charset=utf8")
	if err != nil {
		log.Fatalln(err)
	}
	engin.ShowSQL(true)

	one := Users{
		IMSI: "123456789",
		Mcc:  1,
	}
	engin.Get(&one)
	all := []Users{}
	engin.Find(&all)
	log.Println(one)
}

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
