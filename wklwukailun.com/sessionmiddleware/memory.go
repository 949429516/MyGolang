package sessionmiddleware

import (
	"errors"
	"sync"
)

// 对象
type MemorySession struct {
	sessionId string
	// 储存k-v
	data   map[string]interface{}
	rwlock sync.RWMutex
}

// 构造函数
func NewMemorySession(id string) *MemorySession {
	return &MemorySession{
		sessionId: id,
		data:      make(map[string]interface{}, 16),
	}
}
func (m *MemorySession) Set(key string, value interface{}) (err error) {
	// 加锁
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	// 设置值
	m.data[key] = value
	return
}
func (m *MemorySession) Get(key string) (value interface{}, err error) {
	// 加锁
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	value, ok := m.data[key]
	if !ok {
		err = errors.New("key not exists in session")
		return
	}
	return
}
func (m *MemorySession) Del(key string) (err error) {
	// 加锁
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	delete(m.data, key)
	return
}
func (m *MemorySession) Save() (err error) {
	return
}
