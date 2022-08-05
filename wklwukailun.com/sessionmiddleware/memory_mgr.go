package sessionmiddleware

import (
	"errors"
	"sync"

	uuid "github.com/satori/go.uuid"
)

// 定义对象
type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwlock     sync.RWMutex
}

// 构造函数
func NewMemorySessionMgr() *MemorySessionMgr {
	return &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
}

func (m *MemorySessionMgr) Init(addr string, options ...string) (err error) {
	return
}

// 创建一个session
func (m *MemorySessionMgr) CreateSession() (session Session, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	id := uuid.NewV4() //uuid作为sessionId
	// 转换为string类型
	sessionId := id.String()
	// 创建session
	session = NewMemorySession(sessionId)
	m.sessionMap[sessionId] = session
	return
}
func (m *MemorySessionMgr) Get(sessionId string) (session Session, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	session, ok := m.sessionMap[sessionId]
	if !ok {
		err = errors.New("get session failed from MemorySessionMgr")
		return
	}
	return
}
