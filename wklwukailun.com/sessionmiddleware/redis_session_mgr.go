package sessionmiddleware

import (
	"errors"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
)

// 定义对象
type RedisSessionMgr struct {
	addr       string      // redis地址
	passwd     string      // 密码
	pool       *redis.Pool // 连接池
	sessionMap map[string]Session
	rwlock     sync.RWMutex
}

func NewRedisSessionMgr() SessionMgr {
	return &RedisSessionMgr{
		sessionMap: make(map[string]Session, 32),
	}
}
func myPool(addr, passwd string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: time.Duration(time.Second * 240),
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			// 如有有密码
			if _, err := conn.Do("AUTH", passwd); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, err
		},
		// 连接测试,开发时写上线注释
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
}
func (r *RedisSessionMgr) Init(addr string, options ...string) (err error) {
	// 如有其他参数
	if len(options) > 0 {
		r.passwd = options[0]
	}
	// 创建连接池
	r.pool = myPool(addr, r.passwd)
	r.addr = addr
	return
}
func (r *RedisSessionMgr) CreateSession() (session Session, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	id := uuid.NewV4() //uuid作为sessionId
	// 转换为string类型
	sessionId := id.String()
	// 创建session
	session = NewRedisSession(sessionId, r.pool)
	r.sessionMap[sessionId] = session
	return
}
func (r *RedisSessionMgr) Get(sessionId string) (session Session, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	session, ok := r.sessionMap[sessionId]
	if !ok {
		err = errors.New("get session failed from RedisSessionMgr")
		return
	}
	return
}
