package redis

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

/**
通过go向redis写入和读取数据 操作string
*/
func Case_1() {

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

	//向redis写入数据 string[key-value]
	writeRedis, writeRedisErr := connRedis.Do("set", "name", "go语言")
	if writeRedisErr != nil {
		fmt.Printf("writeRedisErr=%v\n", writeRedisErr)
		return
	}
	fmt.Printf("writeRedis suc：%v\n", writeRedis)

	//从redis读取数据 string[key-value]
	//connRedis.Do()返回的readRedis是interface{}数据类型 要转为string
	//通过redis.String()方法将readRedis转为string类型 (如果是多个字符串就用redis.Strings())
	readRedis, readRedisErr := redis.String(connRedis.Do("get", "name"))
	if readRedisErr != nil {
		fmt.Printf("readRedisErr=%v\n", readRedisErr)
		return
	}

	//打印读取结果
	fmt.Printf("readRedis：%v\n", readRedis)
}

//批量操作string
func Case_2() {

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

	//向redis写入数据 string[key-value][key-value]
	writeRedis, writeRedisErr := connRedis.Do("mset", "name", "go语言", "address", "us")
	if writeRedisErr != nil {
		fmt.Printf("writeRedisErr=%v\n", writeRedisErr)
		return
	}
	fmt.Printf("writeRedis suc：%v\n", writeRedis)

	//从redis读取数据 string[key-value][key-value]
	//connRedis.Do()返回的readRedis是interface{}数据类型 要转为string
	//通过redis.String()方法将readRedis转为string类型 (如果是多个字符串就用redis.Strings())
	readRedis, readRedisErr := redis.Strings(connRedis.Do("mget", "name", "address"))
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
