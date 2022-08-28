package registry

import (
	"context"
	"fmt"
	"sync"
)

// 声明管理者结构体
type PluginMgr struct {
	// map维护所有的插件
	plugins map[string]Registry
	lock    sync.RWMutex
}

var (
	pluginMgr = &PluginMgr{
		plugins: make(map[string]Registry),
	}
)

// 插件注册
func RegisterPlugin(registry Registry) (err error) {
	return pluginMgr.registerPlugin(registry)
}

// 注册插件
func (p *PluginMgr) registerPlugin(plugin Registry) (err error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	// 先去判断是否有
	_, ok := p.plugins[plugin.Name()]
	if ok {
		err = fmt.Errorf("registry plugin exist")
		return
	}
	// 加入map中
	p.plugins[plugin.Name()] = plugin
	return
}

// 进行初始化注册中心
func InitRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	return pluginMgr.initRegistry(ctx, name, opts...)
}

func (p *PluginMgr) initRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	// 先去判断服务列表服务是否存在，若不存在说明没有注册不需要初始化，则报错
	plugin, ok := p.plugins[name]
	if !ok {
		err = fmt.Errorf("plugin %s not exist", name)
		return
	}
	// 存在则返回值赋值
	registry = plugin
	// 插件初始化
	err = plugin.Init(ctx, opts...)
	return
}
