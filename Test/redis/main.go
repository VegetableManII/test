package main

import (
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

const script = `
local res = redis.call('get',KEYS[1])
redis.call('expire',KEYS[1],ARGV[1])
return res`

var lua *redis.Script

func init() {
	lua = redis.NewScript(1, script)
}

func main() {
	//lua = redis.NewScript(1, script)

	lua1 := redis.NewScript(1, script)
	c, e := redis.Dial("tcp", "127.0.0.1:6379",
		redis.DialClientName("test"),
		redis.DialConnectTimeout(1*time.Second),
		redis.DialReadTimeout(1*time.Second),
		redis.DialWriteTimeout(1*time.Second),
	)
	if e != nil {
		log.Fatalln(e)
	}
	lua1.Load(c)
	args := redis.Args{}.Add(interface{}("haha"), interface{}(10))
	reply, err := redis.String(lua.Do(c, args...))
	if err != nil {
		log.Println("err", err)
	}
	log.Println("replay", reply)
}
