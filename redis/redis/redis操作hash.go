package redis

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

/**
通过go向redis写入和读取数据 操作hash
*/

func Case_3() {

	//连接到redis
	connRedis, connRedisErr := redis.Dial("tcp", "127.0.0.1:6379")
	if connRedisErr != nil {
		fmt.Printf("connRedisErr=%v\n", connRedisErr)
		return
	}

	//函数执行完毕 关闭connRedis
	defer func() {
		connRedisCloseErr := connRedis.Close()
		if connRedisCloseErr != nil {
			fmt.Printf("connRedisCloseErr=%v\n", connRedisCloseErr)
			return
		}
		fmt.Printf("connRedisClose suc")
	}()

	//向redis写入数据 hash[[key-value][key-value]]
	//写入user[name-go语言]
	writeRedis01, writeRedisErr01 := connRedis.Do("hset", "user", "name", "go语言")
	if writeRedisErr01 != nil {
		fmt.Printf("writeRedisErr01=%v\n", writeRedisErr01)
		return
	}
	fmt.Printf("writeRedis01 suc：%v\n", writeRedis01)
	//写入user[age-18]
	writeRedis02, writeRedisErr02 := connRedis.Do("hset", "user", "age", 18)
	if writeRedisErr02 != nil {
		fmt.Printf("writeRedisErr02=%v\n", writeRedisErr02)
		return
	}
	fmt.Printf("writeRedis02 suc：%v\n", writeRedis02)

	//从redis读取数据 string[key-value]
	//connRedis.Do()返回的readRedis01是interface{}数据类型 要转为string
	//通过redis.String()方法将readRedis01转为string类型 (如果是多个字符串就用redis.Strings())
	//读取user中key为name的value
	readRedis01, readRedisErr01 := redis.String(connRedis.Do("hget", "user", "name"))
	if readRedisErr01 != nil {
		fmt.Printf("readRedisErr01=%v\n", readRedisErr01)
		return
	}
	//connRedis.Do()返回的readRedis02是interface{}数据类型 要转为int
	//通过redis.Int()方法将readRedis02转为int类型 (如果是多个字符串就用redis.Ints())
	//读取user中key为age的value
	readRedis02, readRedisErr02 := redis.Int(connRedis.Do("hget", "user", "age"))
	if readRedisErr02 != nil {
		fmt.Printf("readRedisErr02=%v\n", readRedisErr02)
		return
	}

	//打印读取结果
	fmt.Printf("readRedis01:%v\nreadRedis02:%v\n", readRedis01, readRedis02)
}

//批量操作set
func Case_4() {

	//连接到redis
	connRedis, connRedisErr := redis.Dial("tcp", "127.0.0.1:6379")
	if connRedisErr != nil {
		fmt.Printf("connRedisErr=%v\n", connRedisErr)
		return
	}

	//函数执行完毕 关闭connRedis
	defer func() {
		connRedisCloseErr := connRedis.Close()
		if connRedisCloseErr != nil {
			fmt.Printf("connRedisCloseErr=%v\n", connRedisCloseErr)
			return
		}
		fmt.Printf("connRedisClose suc\n")
	}()

	//向redis写入数据 hash[[key-value][key-value]]
	//写入user[[name-go语言][age-18]]
	writeRedis, writeRedisErr := connRedis.Do("hmset", "user", "name", "golang", "age", 22)
	if writeRedisErr != nil {
		fmt.Printf("writeRedisErr=%v\n", writeRedisErr)
		return
	}
	fmt.Printf("writeRedis suc：%v\n", writeRedis)

	//从redis读取数据 string[key-value]
	//connRedis.Do()返回的readRedis01是interface{}数据类型 要转为string
	//通过redis.String()方法将readRedis01转为string类型 (如果是多个字符串就用redis.Strings())
	//一次性读取user中key为name的value 和key为age的value
	readRedis, readRedisErr := redis.Strings(connRedis.Do("hmget", "user", "name", "age"))
	if readRedisErr != nil {
		fmt.Printf("readRedisErr=%v\n", readRedisErr)
		return
	}

	//打印读取结果
	//批量操作时返回的readRedis是切片类型 遍历readRedis切片
	for i, v := range readRedis {
		fmt.Printf("readRedis[%v]：%v\n", i, v)
	}

}
