package questions

import (
	"log"
	"regexp"
)

func ParseDBdsn() {
	dsnPattern := regexp.MustCompile(
		`^(?:(?P<user>.*?)(?::(?P<passwd>.*))?@)?` + // [user[:password]@]
			`(?:(?P<net>[^\(]*)(?:\((?P<addr>[^\)]*)\))?)?` + // [net[(addr)]]
			`\/(?P<dbname>.*?)` + // /dbname
			`(?:\?(?P<params>[^\?]*))?$`) // [?param1=value1&paramN=valueN]
	dsn := "jack:ajksfhk:kcxjl@tcp(127.0.0.1:3307)/user?param1=v1&param2=v2&param3=v3"

	names := dsnPattern.SubexpNames()
	log.Println(names)
	match := dsnPattern.FindStringSubmatch(dsn)
	log.Println(match)
}
