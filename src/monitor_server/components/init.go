package components

import(
    "fmt"
    "github.com/garyburd/redigo/redis"
)

var RedisHandleObj *RedisHandle

// redis类
type RedisHandle struct {
    RedisPool   *redis.Pool
}

func NewRedisHandle() *RedisHandle {
    // 初始化redis
    RedisHandleObj := &RedisHandle{
        RedisPool : NewRedisPool(),
    }
    if RedisHandleObj.RedisPool != nil {
		fmt.Println("redis init complete")
    }
    return RedisHandleObj
}
