package boltdb

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("./dbFile/123.db", 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte("bucket0"))
		if bkt == nil {
			//没有bucket
			bkt, err = tx.CreateBucket([]byte("bucket0"))
			if err != nil {
				log.Panic(err)
				return err
			}
		}
		bkt.Put([]byte("test_1"), []byte("hello world"))
		bkt.Put([]byte("test_2"), []byte("hello world"))
		return nil
	})
	db.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte("bucket0"))
		if bkt == nil {
			log.Panic("非法:Bucket内容为空")
		}
		v1 := bkt.Get([]byte("test_1"))
		v2 := bkt.Get([]byte("test_2"))

		fmt.Printf("取得数据库中的数据 test_1==%s\n", v1)
		fmt.Printf("取得数据库中的数据 test_2==%s", v2)

		return nil
	})

}
