package components

import(
    "github.com/garyburd/redigo/redis"
    "github.com/astaxie/beego/logs"
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
        logs.Info("redis init complete")
    }
    return RedisHandleObj
}
