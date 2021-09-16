package main

import (
	"errors"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

const script = `
local res = redis.call('get',KEYS[1])
redis.call('expire',KEYS[1],ARGV[1])
return res`

var luas map[string]*redis.Script

type RedisConnSt struct {
	conn   redis.Conn
	prefix string
}

func init() {
	luas = make(map[string]*redis.Script)
	scriptStr := map[string]string{
		"s1": `
		if ( redis.call('ttl',KEYS[1]) <= -1 )  -- [ 如果模板不存在或未设置ttl ]
		then
			redis.call('incr',KEYS[1])
			redis.call('expire',KEYS[1],ARGV[1])
		else
			redis.call('incr',KEYS[1])
		end
		return
		`,
		"s2": `
			local res = redis.call('get',KEYS[1])
			redis.call('expire',KEYS[1],ARGV[1])
			return res
		`,
	}
	for k, v := range scriptStr {
		luas[k] = redis.NewScript(-1, v)
	}
}

func main() {
	c, e := redis.Dial("tcp", "127.0.0.1:6379",
		redis.DialClientName("test"),
		redis.DialConnectTimeout(1*time.Second),
		redis.DialReadTimeout(1*time.Second),
		redis.DialWriteTimeout(1*time.Second),
	)
	if e != nil {
		log.Fatalln(e)
	}
	rc := RedisConnSt{
		conn:   c,
		prefix: "jojo",
	}
	value, err := rc.EvalScript("s1", 1, "dio", 600)
	if err == redis.ErrNil {
		log.Println(err)
	}
	log.Println(value)
}

// ExecScript 参数，numkey为脚本中使用到的KEY的数量，argvs为脚本中的所有的KEY和ARGV
// 要求所有的KEY排在所有的ARGV前面，并且KEY与ARGV对应，所有的KEY都为字符串类型
// eg:
//   ExecScript("script1",numkeys,key1,key2 [key...],arg [arg...])
//   ExecScript("script2",numkeys,key1,key2 [key...],arg [arg...])
func (rc *RedisConnSt) EvalScript(script string, numkey int, argvs ...interface{}) (interface{}, error) {
	if _, ok := luas[script]; !ok {
		return "", nil
	}
	lua := luas[script]
	lua.Load(rc.conn)
	for i := 0; i < numkey; i++ {
		if s, ok := argvs[i].(string); !ok {
			return "", errors.New("KEYS type assert is not string")
		} else {
			argvs[i] = rc.prefix + s
		}
	}
	rargs := []interface{}{}
	rargs = append(rargs, numkey)
	rargs = append(rargs, argvs...)
	value, err := redis.String(lua.Do(rc.conn, rargs...))
	if err != nil && err != redis.ErrNil {
		return "", err
	}
	return value, nil
}
