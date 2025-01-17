package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/rpc"
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

// 0还没有任期,-1没有leader编号
var leader = Leader{0, -1}

func Make(me int) *Raft {
	rand.Seed(time.Now().UnixNano()) // 设置随机种子
	rf := &Raft{
		me:            me,
		currentTerm:   0,
		votedFor:      -1, //-1谁都不投票,此为节点刚创建
		state:         0,  //0 follower
		currentLeader: -1,
		message:       make(chan bool),
		electCh:       make(chan bool),
		heartBeat:     make(chan bool),
		heartbeatRe:   make(chan bool),
		timeout:       0,
	}
	// 选举的协程
	go rf.election()
	// 心跳检测的协程
	go rf.sendLeaderHeartBeat()
	return rf
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
			result = rf.election_one_round(&leader)
		}
	}
}

// 选主逻辑
func (rf *Raft) election_one_round(leader *Leader) bool {
	var vote int             // 投票数量
	var triggerHearbeat bool // 心跳,是否开始心跳信号的产生
	timeout := int64(100)    // 定义超时
	last := millisecond()    //时间
	success := false         // 用于返回值
	// 加锁修改状态
	rf.mu.Lock()
	// 当前节点变成candidate
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
		// 校验：若不超时，且票数大于一半则选举成功，break
		if timeout+last < millisecond() || (vote > raftCount/2 || rf.currentLeader > -1) {
			break
		} else {
			// 等待操作
			select {
			case <-time.After(time.Duration(10) * time.Millisecond):
			}
		}
	}
	return success
}

// 修改状态为candidate
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

// 1.leader节点发送心跳信号,2.数据传输、同步3.检查从节点是否正常
func (rf *Raft) sendLeaderHeartBeat() {
	// 死循环
	for {
		select {
		case <-rf.heartBeat:
			rf.sendAppendEntriesImpl()
		}
	}
}

// 用于返回给leader的确认信号
func (rf *Raft) sendAppendEntriesImpl() {
	// 如果是主节点不执行
	var success_count int
	if rf.currentLeader == rf.me {
		// 此时自己就是leader
		success_count = 0 // 记录确认信号的节点个数
		// 设置确认信号
		for i := 0; i < raftCount; i++ {
			if i != rf.me {
				go func() {
					//rf.heartbeatRe <- true
					// 相当于客户端
					client, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
					if err != nil {
						log.Fatal(err)
					}
					// 接收服务器返回的信息
					var ok = false // 服务端返回信息的变量
					err = client.Call("Raft.Commnication", Param{"hello"}, &ok)
					if err != nil {
						log.Fatal(err)
					}
					if ok {
						rf.heartbeatRe <- true
					}
				}()
			}
		}
	}
	// 计算返回确认信号个数
	for i := 0; i < raftCount; i++ {
		select {
		case ok := <-rf.heartbeatRe:
			if ok {
				success_count++
				if success_count > raftCount/2 {
					fmt.Println("投票选举成功,心跳信号OK")
					log.Fatal("程序结束")
				}
			}
		}
	}
}
func main() {
	// 过程：有3个节点，最初都是follower
	// 若有candidate状态，进行投票拉票
	// 产生leader

	// 创建3个节点
	for i := 0; i < raftCount; i++ {
		// 创建3个raft节点
		Make(i)
	}
	// 加入服务端监听
	rpc.Register(new(Raft))
	rpc.HandleHTTP()
	// 监听服务
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
	}
}

// 首字母大写,RPC规范
// 分布式通信
type Param struct {
	Msg string
}

// 通信方法
func (r *Raft) Commnication(p Param, a *bool) error {
	fmt.Println(p.Msg)
	*a = true
	return nil
}
