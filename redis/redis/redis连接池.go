package redis

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

/**
redis连接池
创建一个redis连接池放入多个redisConn 当需要用redisConn时就从连接池里取出 不使用时就放入redis连接池
这样可以节省临时获取redis链接的时间 从而提高效率
注意：当pool被Close()后是不能取出conn的
*/

var pool *redis.Pool //全局变量

func Case_5() {

	//创建redis连接池并放入redis链接
	pool = &redis.Pool{
		MaxIdle:     5,                                                                         //最大空闲链接数
		MaxActive:   0,                                                                         //和redis数据库的最大连接数 0表示没有限制
		IdleTimeout: 100,                                                                       //链接的最大空闲时间 超过100s链接未使用就放入连接池中
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", "localhost:6379") }, //创建redis连接池
	}

	//从pool中取出一个conn
	conn := pool.Get()

	//函数退出时关闭conn链接
	defer func() {
		connCloseErr := conn.Close()
		if connCloseErr != nil {
			fmt.Printf("connCloseErr:%v\n", connCloseErr)
			return
		}
		fmt.Println("connClose suc...")
	}()

	//使用从pool中取出的conn 连接redis操作数据
	//写入
	writeRedis, writeRedisErr := conn.Do("set", "name", "golang")
	if writeRedisErr != nil {
		fmt.Printf("writeRedisErr:%v\n", writeRedisErr)
		return
	}
	fmt.Printf("writeRedis suc:%v\n", writeRedis)
	//读取
	readRedis, readRedisErr := redis.String(conn.Do("get", "name"))
	if readRedisErr != nil {
		fmt.Printf("readRedisErr:%v\n", readRedisErr)
		return
	}
	fmt.Printf("readRedis suc:%v\n", readRedis)
}
