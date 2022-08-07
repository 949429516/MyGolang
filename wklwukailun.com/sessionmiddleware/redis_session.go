package sessionmiddleware

import (
	"encoding/json"
	"errors"
	"sync"

	"github.com/garyburd/redigo/redis"
)

type RedisSession struct {
	sessionId string
	pool      *redis.Pool
	// 设置session先放在内存map中,批量倒入redis中提升性能
	// 缺点容易丢失且不是实时
	sessionMap map[string]interface{}
	rwlock     sync.RWMutex //读写锁
	// 记录内存中map是否被操作
	flag int
}

// 用常量定义状态
const (
	// 内存没有变化
	SessionFlagNone = iota
	// 有变化
	SessionFlagModify
)

// 构造函数
func NewRedisSession(id string, pool *redis.Pool) *RedisSession {
	return &RedisSession{
		sessionId:  id,
		pool:       pool,
		sessionMap: make(map[string]interface{}, 0),
		flag:       SessionFlagNone,
	}
}
func (r *RedisSession) Get(key string) (result interface{}, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	// 先从内存中获取
	result, ok := r.sessionMap[key]
	if !ok {
		err = errors.New("key not exists")
		return
	}
	return
}

// 从redis里加载
func (r *RedisSession) loadFromRedis() (err error) {
	conn := r.pool.Get()
	reply, err := conn.Do("GET", r.sessionId)
	if err != nil {
		return
	}
	// 转换字符
	data, err := redis.String(reply, err)
	if err != nil {
		return
	}
	// 反序列化
	err = json.Unmarshal([]byte(data), &r.sessionMap)
	if err != nil {
		return
	}
	return
}

// session存储到redis
func (r *RedisSession) Set(key string, value interface{}) (err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	r.sessionMap[key] = value
	// 标记记录
	r.flag = SessionFlagModify
	return
}
func (r *RedisSession) Del(key string) (err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	// 标记记录
	r.flag = SessionFlagModify
	delete(r.sessionMap, key)
	return
}
func (r *RedisSession) Save() (err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	// 如果数据没有改变则不需要存储redis
	if r.flag != SessionFlagModify {
		return
	}
	// 内存中的redis序列化
	data, err := json.Marshal(r.sessionMap)
	if err != nil {
		return
	}
	// 获取redis连接
	conn := r.pool.Get()
	_, err = conn.Do("SET", r.sessionId, string(data))
	r.flag = SessionFlagNone
	if err != nil {
		return
	}
	return
}
