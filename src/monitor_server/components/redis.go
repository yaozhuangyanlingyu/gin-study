package components

import(
    "fmt"
    "monitor_server/tools"
    "github.com/garyburd/redigo/redis"
)

/**
 * 初始化redis
 */
func NewRedisPool() (*redis.Pool) {
    host := tools.IniConf.String(fmt.Sprintf("Redis%s::Host", "Dev"))
    port := tools.IniConf.String(fmt.Sprintf("Redis%s::Port", "Dev"))
    database := tools.IniConf.String(fmt.Sprintf("Redis%s::Database", "Dev"))
	address := fmt.Sprintf("%s:%s", host, port)
	return &redis.Pool{
        MaxIdle:     5,
        Dial: func() (c redis.Conn, err error) {
            c, err = redis.Dial("tcp", address)
            if err != nil {
                return nil, err
            }
            if _, err := c.Do("SELECT", database); err != nil {
                c.Close()
                return nil, err
            }
            return c, err
        },
	}
}

/*
 * 设置key
 */
func (this *RedisHandle) Set(key string, value string, expire int) error {
    cache := this.RedisPool.Get()
    defer cache.Close()
    _, err := cache.Do("SET", key, value, "EX", expire)
    return err
}

/**
 * 获取key
 */
func (this *RedisHandle) GetString(key string) string {
    cache := this.RedisPool.Get()
    defer cache.Close()
	v, err := redis.String(cache.Do("Get", key))
	if err != nil {
		return ""
	}
	return v
}


