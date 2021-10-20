package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type A struct {
	Field_a int       `json:"field_a"`
	Field_b string    `json:"field_b"`
	Field_c time.Time `json:"field_c"`
	Field_d int       `json:"field_d"`
	Field_e string    `json:"field_e"`
	Field_f string    `json:"field_f"`
	Field_g int       `json:"field_g"`
	Field_h string    `json:"field_h"`
	Field_i int       `json:"field_i"`
	Field_j int       `json:"field_j"`
	Field_k string    `json:"field_k"`
	Field_l string    `json:"field_l"`
}
type Rule struct {
	Field     string `json:"field"`
	Type      string `json:"type"`
	Operation string `json:"operation"`
	Members   string `json:"members"`
}

type Object struct {
	ObjA  *A
	Rules []*Rule
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	atest := A{
		Field_c: time.Date(1999, 1, 1, 0, 0, 0, 0, time.Local),
	}
	Obj := new(Object)
	Obj.ObjA = &atest
	res, e := db.Query("SELECT * FROM t2")
	if e != nil {
		log.Println(e)
	}
	var col []byte
	res.Next()
	res.Scan(&col)

	Obj.Rules = make([]*Rule, 0, 1)
	e = json.Unmarshal(col, &Obj.Rules)
	if e != nil {
		log.Println(e)
	}

	for i, v := range Obj.Rules {
		log.Println("rule ", i+1, v)
	}

	log.Println(paraseRules(Obj.ObjA, Obj.Rules, 3))

}

var aStuFieldMap map[string]interface{}

// 支持解析三种数据类型  整数、字符串、时间（必须符合格式1999-08-01 01:01:01）
// 整数支持 < > <= >= !(不等于) =(等于) 六种运算
// 字符串支持 !(不等于) =(等于) 两种运算
// 时间支持 < > <= >= =(等于) 六种运算
func paraseRules(obj *A, rules []*Rule, fieldNum int) bool {
	aStuFieldMap = getFieldMap(obj, fieldNum)
	for _, v := range rules {
		switch v.Type {
		case "int":
			left := aStuFieldMap[v.Field].(int)
			rightNumbers := strings.Split(v.Members, ",")
			switch v.Operation {
			case "<":
				for _, v := range rightNumbers {
					right, _ := strconv.Atoi(v)
					if left < right || left == 0 {
						continue
					} else {
						return false
					}
				}
			case "<=":
				for _, v := range rightNumbers {
					right, _ := strconv.Atoi(v)
					if left <= right || left == 0 {
						continue
					} else {
						return false
					}
				}
			case ">":
				for _, v := range rightNumbers {
					right, _ := strconv.Atoi(v)
					if left > right || left == 0 {
						continue
					} else {
						return false
					}
				}
			case ">=":
				for _, v := range rightNumbers {
					right, _ := strconv.Atoi(v)
					if left >= right || left == 0 {
						continue
					} else {
						return false
					}
				}
			case "!":
				if !isStringsContainsInteger(rightNumbers, left) {
					continue
				} else {
					return false
				}
			case "=":
				if isStringsContainsInteger(rightNumbers, left) {
					continue
				} else {
					return false
				}
			default:
				return false
			}
		case "string":
			left := aStuFieldMap[v.Field].(string)
			rightVars := v.Members
			switch v.Operation {
			case "!":
				if !strings.Contains(rightVars, left) {
					continue
				} else {
					return false
				}
			case "=":
				if strings.Contains(rightVars, left) {
					continue
				} else {
					return false
				}
			default:
				return false

			}
		case "time":
			left := aStuFieldMap[v.Field].(time.Time)
			rightTimes := strings.Split(v.Members, ",")
			switch v.Operation {
			case "<":
				for _, v := range rightTimes {
					right, e := time.ParseInLocation("2006-01-02 15:04:05", v, time.Local)
					if e != nil {
						log.Println(e)
					}
					if left.Before(right) {
						continue
					} else {
						return false
					}
				}
			case "<=":
				for _, v := range rightTimes {
					right, e := time.ParseInLocation("2006-01-02 15:04:05", v, time.Local)
					if e != nil {
						log.Println(e)
					}
					if left.Before(right) || left.Equal(right) {
						continue
					} else {
						return false
					}
				}
			case ">":
				for _, v := range rightTimes {
					right, e := time.ParseInLocation("2006-01-02 15:04:05", v, time.Local)
					if e != nil {
						log.Println(e)
					}
					if left.After(right) {
						continue
					} else {
						return false
					}
				}
			case ">=":
				for _, v := range rightTimes {
					right, e := time.ParseInLocation("2006-01-02 15:04:05", v, time.Local)
					if e != nil {
						log.Println(e)
					}
					if left.After(right) || left.Equal(right) {
						continue
					} else {
						return false
					}
				}
			}
		}
	}
	return true
}

func getFieldMap(obj interface{}, fieldnum int) map[string]interface{} {
	tType := reflect.TypeOf(obj).Elem()
	tVal := reflect.ValueOf(obj).Elem()
	fields := make(map[string]interface{}, fieldnum)
	for i := 0; i < tType.NumField(); i++ {
		tag := tType.Field(i).Tag.Get("json")
		field := tVal.Field(i).Interface()
		fields[tag] = field
	}
	return fields
}

func isStringsContainsInteger(numbers []string, target int) bool {
	p1, p2 := 0, len(numbers)-1
	for p1 < p2 {
		num1, e := strconv.Atoi(numbers[p1])
		if e != nil {
			return false
		}
		if target == num1 {
			return true
		}
		p1--
		num2, e := strconv.Atoi(numbers[p1])
		if e != nil {
			return false
		}
		if target == num2 {
			return true
		}
		p2++
	}
	return false
}
