package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 1.实现3节点选举
// 2.改造为分布式选举代码，加入RPC调用

// 定义三节点常量
const raftCount = 3

// 声明leader对象
type Leader struct {
	Term     int // 任期
	LeaderId int // leader编号
}

// 声明raft
type Raft struct {
	mu              sync.Mutex // 锁
	me              int        // 节点编号
	currentTerm     int        // 当前任期
	votedFor        int        // 为哪个节点投票
	state           int        // 3个状态 0follower 1candidate 2leader
	lastMessageTime int64      // 发送最后一条数据的时间
	currentLeader   int        // 设置当前节点的领导
	message         chan bool  // 节点间发送信息的通道
	electCh         chan bool  // 选举通道
	heartBeat       chan bool  // 心跳信号通道
	heartbeatRe     chan bool  // 返回心跳信号通道
	timeout         int        // 超时时间
}

// 0还没有商人,-1没有编号
var leader = Leader{0, -1}

func Make(me int) *Raft {
	rand.Seed(time.Now().UnixNano()) // 设置随机种子
	return &Raft{
		me:            me,
		currentTerm:   0,
		votedFor:      -1, //-1谁都不投票,次为节点刚创建
		state:         0,  //0 follower
		currentLeader: -1,
		message:       make(chan bool),
		electCh:       make(chan bool),
		heartBeat:     make(chan bool),
		heartbeatRe:   make(chan bool),
		timeout:       0,
	}
}

// 设置任期
func (rf *Raft) setTerm(term int) {
	rf.currentTerm = term
}

// 随机值min到max的随机值
func randRange(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

// 获取当前时间
func millisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// 选主
func (rf *Raft) election() {
	// 设置标记，判断是否选出了leader
	var result bool
	for {
		// 设置超时,150~300ms
		timeout := randRange(150, 300)
		rf.lastMessageTime = millisecond()
		select {
		// 延迟等待ms
		case <-time.After(time.Duration(timeout) * time.Millisecond):
			fmt.Println("当前节点状态为:", rf.state)
		}
		result = false
		for !result {
			// 选主leader
		}
	}
}

// 选主逻辑
func (rf *Raft) election_one_round(leader *Leader) bool {
	var timeout int64 // 定义超时
	timeout = 100
	var vote int             // 投票数量
	var triggerHearbeat bool // 心跳,是否开始心跳信号的产生
	last := millisecond()    //时间
	success := false         // 用于返回值
	// 当前节点变成candidate
	rf.mu.Lock()
	//修改状态
	rf.becomeCandidate()
	rf.mu.Unlock()
	fmt.Println("start electing leader")
	for {
		// 遍历所有节点拉选票
		for i := 0; i < raftCount; i++ {
			if i != rf.me {
				// 拉选票
				go func() {
					if leader.LeaderId < 0 {
						// 设置投票
						rf.electCh <- true
					}
				}()
			}
		}
		// 设置投票数量
		vote = 1
		// 遍历节点增加选票
		for i := 0; i < raftCount; i++ {
			// 计算投票的数量
			select {
			case ok := <-rf.electCh:
				if ok {
					// 投票数量+1
					vote++
					// 若选票过半则成功
					success = vote > raftCount/2
					if success && !triggerHearbeat {
						// 触发心跳检测
						triggerHearbeat = true
						// 变换成主节点leader
						rf.mu.Lock()
						rf.becomeLeader()
						rf.mu.Unlock()
						// 心跳监听,由leader向其他节点发送心跳信号
						rf.heartBeat <- true
						fmt.Println(rf.me, "号节点成为了leader")
						fmt.Println("leader开发发送心跳信号")
					}
				}
			}
		}

	}
}

// 修改状态candidate
func (rf *Raft) becomeCandidate() {
	rf.state = 1
	rf.setTerm(rf.currentTerm + 1)
	rf.votedFor = rf.me
	rf.currentLeader = -1
}

// 修改为leader
func (rf *Raft) becomeLeader() {
	rf.state = 2
	rf.currentLeader = rf.me
}
func main() {
	// 过程：有3个节点，最初都是follower
	// 若有candidate状态，进行投票拉票
	// 产生leader

	// 创建3个节点
	for i := 0; i < raftCount; i++ {
		// 创建3个raft节点
	}
	for {
	}
}
