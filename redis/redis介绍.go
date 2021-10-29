package main

/**
redis环境配置：
使用第三方开源的redis库：github.com/garyburd/redigo/redis
1.安装配置git (git官网下载安装)
2.在GOPATH环境变量路径下运行命令提示符执行安装命令：go get github.com/garyburd/redigo/redis
3.执行后如果src文件夹下面新增一个名为 github.com 文件夹，就表示redis环境配置成功。
 */

/**
redis介绍：
1.redis支持的五大数据：String(字符串)、Hash(哈希表)、List(列表)、Set(集合)、zset(sorted set:有序集合)
2.redis有16[0-15号]个默认数据库，初始默认0号。

redis基本使用：
1.查看当前所有的key[keys *]
2.查看当前key-val的数量[dbsize]
3.清空当前数据库key-val[flushdb]  清空所有数据库key-val[flushall]
4.切换数据库[select i]

redis使用细节：
string：
1.string是redis最基本得类型，一个key对应一个value。
2.string类型是二进制安全的，除普通的字符串外，也可以存放图片等数据。
3.redis中value最大是512M
4.set [set key01 golang]  (添加一个key为key01 value为golang的键值对) (有就修改 没有就添加)
5.get [get key01]  (查看key为key01的键值对的value)
6.setex [setex key 时间 value] (添加一个定时的string键值对，时间一到自动删除。时间单位：秒。)
7.mset [mset key01 value01 key02 value02]  (添加多个string键值对)
8.mget [mget key01 key02]  (查看多个string键值对)

hash
redis hash是一个键值对集合，是一个string类型的field和value的映射表。
1.hset [hset user01 name golang] (给user01存储一个key-valuse key是name value是golang) (有就修改 没有就添加)
2.hget [hget user01 name] (查看user01的key为name的value)
3.hgetall [hgetall user01] (取出user01的所有键值对)
4.hdel [hdel user01 name](删除user01的key为name的键值对)
5.hmset [hmset user02 name "hello go" age 20] (字段中间有空格的要用引号括起来) (一次性给user02添加 name-"hello go" age-20 多个键值对)
6.hmget [hmset user02 name age] (查看user02的 key为name的value 和 key为age的value)
7.hlen [hlen user02] (查看user02有多少个键值对)
8.hexists [hexists user02 name] (查看user02是否有key为name的键值对)

List(列表)
list是简单的string列表，按照插入顺序排序，你可以添加一个元素到头部或尾部。(可以从左或从右插入添加)
list本质是个链表，list元素是有序的，元素的value可以重复。
注意：如果list的元素全部移除，那么list也就消失了。
1.lpush [lpush city aa bb cc] (给名为city的list列表(没有就创建)，从右向左放入元素aa bb cc) (lpush表示从右向左放入元素 此时city list列表中元素分布位置为cc bb aa)
2.rpush [rpush city dd ee] (给名为city的list列表(没有就创建)，从左向右放入元素aa bb cc) (rpush表示从右向左放入元素 此时city list列表中元素分布位置为cc bb aa dd ee)
3.lrange [lrange city 0 2] (将名为city的list列表下标为0到下标为2之间的元素取出) (也可以用负数作为下标：-1为倒数第一个元素 -2为倒数第二个元素 以此类推)
4.lpop [lpop city] (将名为city的list列表最左边的一个元素取出并移除) (最左边的cc被取出并移除 此时city list列表中元素分布位置为bb aa dd ee)
5.rpop [rpop city] (将名为city的list列表最右边的一个元素取出并移除) (最右边的ee被取出并移除 此时city list列表中元素分布位置为bb aa dd)
6.lindex [lindex city 3](查看名为city的list列表从左到右开始 下标为3的元素)
7.del [del city] (删除名为city的list列表) (city以及city所有的元素将被删除)
8.llen [llen city](查看名为city的list列表的长度，即元素个数)(如果city不存在 就返回0)

Set(集合)
redis的Set是string类型的无序集合，底层是hashTable数据结构
Set也是存放很多字符串元素，字符串元素是无需的，且不能重复。
1.sadd [sadd set01 aa bb cc] (给名为set01的Set集合(没有就创建) 添加aa bb cc三个元素)
2.smembers [smembers set01] (查看名为set01的Set集合的所有元素)
3.sismember [sismember set01 bb] (查看名为set01的Set集合是否有bb这个元素)
4.srem [srem set01 cc] (删除set01中cc元素)
*/