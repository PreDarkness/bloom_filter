package main

import (
	hash "bloom_filter"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	con,err := redis.Dial("tcp", ":6379")//连接redis
	print(err,"connect")
	defer con.Close()

	bloom := hash.NewBloom(con) //创建过滤器
	bloom.Add("newClient") //往过滤器写入数据
	b := bloom.Exist("aaa") //判断是否存在这个值
	fmt.Println(b)
}