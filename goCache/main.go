package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)
func main () {

	// 初始化一个cache, 5分钟的有效期, 每10分钟清理一次
	c := cache.New(5 * time.Minute, 10 * time.Minute)

	// 使用c 设置的有效期
	c.Set("ssd", "sb", cache.DefaultExpiration)

	// 跟上面那个等效, 这个明显更好些
	c.SetDefault("ssd2", "sb2")

	// 设置永不过期 --- 此处疑问怎么限制cache的使用内存
	c.Set("baz", 42, cache.NoExpiration)

	// 读取key
	foo, found := c.Get("foo")
	if found {
		fmt.Println(foo)
	}
}
