package magic

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

// RedisClient Redis缓存结构体
type RedisClient struct {
	pool              *redis.Pool
	defaultExpiration time.Duration
}

// NewRedisCache 返回cache 对象, 在多个工具之间建立一个 中间初始化的时候使用
func NewRedisCache() RedisClient {
	redisConfig := Config.Redis
	pool := &redis.Pool{
		MaxActive:   100,                             // 最大连接数，即最多的tcp连接数，一般建议往大的配置，但不要超过操作系统文件句柄个数（centos下可以ulimit -n查看）
		MaxIdle:     100,                             // 最大空闲连接数，即会有这么多个连接提前等待着，但过了超时时间也会关闭。
		IdleTimeout: time.Duration(60) * time.Second, // 空闲连接超时时间|查询:CONFIG GET timeout|设置CONFIG SET timeout 65
		Wait:        true,                            // 当超过最大连接数 是报错还是等待， true 等待 false 报错
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", redisConfig.Addr, redis.DialDatabase(redisConfig.DB), redis.DialPassword(redisConfig.Password))
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			return conn, nil
		},
	}
	return RedisClient{pool: pool, defaultExpiration: 86400 * time.Second}
}

// Option 默认传参
type Option struct {
}

// TODO:用字符串包替换

// Serialization 序列化
func Serialization(value interface{}) ([]byte, error) {
	if bytes, ok := value.([]byte); ok {
		return bytes, nil
	}

	switch v := reflect.ValueOf(value); v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []byte(strconv.FormatInt(v.Int(), 10)), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return []byte(strconv.FormatUint(v.Uint(), 10)), nil
	case reflect.Map:
	}
	k, err := json.Marshal(value)
	return k, err
}

// Deserialization 反序列化
func Deserialization(byt []byte, ptr interface{}) (err error) {
	if bytes, ok := ptr.(*[]byte); ok {
		*bytes = byt
		return
	}
	if v := reflect.ValueOf(ptr); v.Kind() == reflect.Ptr {
		switch p := v.Elem(); p.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			var i int64
			i, err = strconv.ParseInt(string(byt), 10, 64)
			if err != nil {
				fmt.Printf("Deserialization: failed to parse int '%s': %s", string(byt), err)
			} else {
				p.SetInt(i)
			}
			return

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			var i uint64
			i, err = strconv.ParseUint(string(byt), 10, 64)
			if err != nil {
				fmt.Printf("Deserialization: failed to parse uint '%s': %s", string(byt), err)
			} else {
				p.SetUint(i)
			}
			return
		}
	}
	err = json.Unmarshal(byt, &ptr)
	return
}

// TODO:用字符串包替换

// 设置字符串key,valule
func (c RedisClient) Set(key string, val string, expiration int64) (err error) {
	// 获取链接
	conn := c.pool.Get()
	defer conn.Close()

	// 设置kv和过期时间
	if expiration == -1 {
		_, err = conn.Do("SET", key, val)
	} else {
		_, err = conn.Do("SETEX", key, expiration, val)
	}
	return err
}

// SetNx 设置唯一键
func SetNx() error {
	return nil
}

// 获取 字符串类型的值
func (c RedisClient) Get(name string) (val string, err error) {
	conn := c.pool.Get()
	defer conn.Close()
	val, err = redis.String(conn.Do("Get", name))
	return val, err
}

// 判断所在的 key 是否存在
func (c RedisClient) Exist(name string) (bool, error) {
	conn := c.pool.Get()
	defer conn.Close()
	v, err := redis.Bool(conn.Do("EXISTS", name))
	return v, err
}

// 自增
func (c RedisClient) Incr(name string) (int, error) {
	conn := c.pool.Get()
	defer conn.Close()
	v, err := redis.Int(conn.Do("INCR", name))
	return v, err
}

// 设置过期时间 （单位 秒）
func (c RedisClient) Expire(name string, newSecondsLifeTime int64) error {
	// 设置key 的过期时间
	conn := c.pool.Get()
	defer conn.Close()
	_, err := conn.Do("EXPIRE", name, newSecondsLifeTime)
	return err
}

// 删除指定的键
func (c RedisClient) Del(keys ...interface{}) (bool, error) {
	conn := c.pool.Get()
	defer conn.Close()
	v, err := redis.Bool(conn.Do("DEL", keys...))
	return v, err
}

// 查看指定的长度
func (c RedisClient) StrLen(name string) (int, error) {
	conn := c.pool.Get()
	defer conn.Close()
	v, err := redis.Int(conn.Do("STRLEN", name))
	return v, err
}

// 删除指定的 hash 键
func (c RedisClient) Hdel(name, key string) (bool, error) {
	conn := c.pool.Get()
	defer conn.Close()
	var err error
	v, err := redis.Bool(conn.Do("HDEL", name, key))
	return v, err
}

// 查看hash 中指定是否存在
func (c RedisClient) HExists(name, field string) (bool, error) {
	conn := c.pool.Get()
	defer conn.Close()
	var err error
	v, err := redis.Bool(conn.Do("HEXISTS", name, field))
	return v, err
}

// 获取hash 的键的个数
func (c RedisClient) HLen(name string) (int, error) {
	conn := c.pool.Get()
	defer conn.Close()
	v, err := redis.Int(conn.Do("HLEN", name))
	return v, err
}

// 传入的 字段列表获得对应的值
func (c RedisClient) HMget(name string, fields ...string) ([]interface{}, error) {
	conn := c.pool.Get()
	defer conn.Close()
	args := []interface{}{name}
	for _, field := range fields {
		args = append(args, field)
	}
	value, err := redis.Values(conn.Do("HMGET", args...))

	return value, err
}

// 设置单个值, value 还可以是一个 map slice 等
func (c RedisClient) HSet(name string, key string, value interface{}) (err error) {
	conn := c.pool.Get()
	defer conn.Close()
	v, _ := Serialization(value)
	_, err = conn.Do("HSET", name, key, v)
	return
}

// 设置多个值 , obj 可以是指针 slice map struct
func (c RedisClient) HMSet(name string, obj interface{}) (err error) {
	conn := c.pool.Get()
	defer conn.Close()
	_, err = conn.Do("HSET", redis.Args{}.Add(name).AddFlat(&obj)...)
	return
}

// 获取单个hash 中的值
func (c RedisClient) HGet(name, field string, v interface{}) (err error) {
	conn := c.pool.Get()
	defer conn.Close()
	temp, _ := redis.Bytes(conn.Do("Get", name, field))
	err = Deserialization(temp, &v) // 反序列化
	return
}

// 获取 set 集合中所有的元素, 想要什么类型的自己指定
func (c RedisClient) Smembers(name string, v interface{}) (err error) {
	conn := c.pool.Get()
	defer conn.Close()
	temp, _ := redis.Bytes(conn.Do("smembers", name))
	err = Deserialization(temp, &v)
	return err
}

// 获取集合中元素的个数
func (c RedisClient) ScardInt64s(name string) (int64, error) {
	conn := c.pool.Get()
	defer conn.Close()
	v, err := redis.Int64(conn.Do("SCARD", name))
	return v, err
}

// LIST左边插入数据
func (c RedisClient) Lpush(name string, val string) (int64, error) {
	conn := c.pool.Get()
	defer conn.Close()
	v, err := redis.Int64(conn.Do("LPUSH", name, val))
	return v, err
}

// LIST右边弹出数据
func (c RedisClient) Rpop(key string) (string, error) {
	conn := c.pool.Get()
	defer conn.Close()
	v, err := redis.String(conn.Do("RPOP", key))
	return v, err
}
