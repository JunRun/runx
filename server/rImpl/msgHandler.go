/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-10-30 15:49
 */
package rImpl

import (
	"fmt"
	"runx/server/rIterface"
	"runx/util"
	"strconv"
)

type MsgHandler struct {
	Apis map[uint64]rIterface.IRouter

	WorkPoolSize uint64

	TaskQueue []chan rIterface.IRequest
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis:         make(map[uint64]rIterface.IRouter),
		WorkPoolSize: uint64(util.Config.GetInt("server.work-pool-size")),
		TaskQueue:    make([]chan rIterface.IRequest, util.Config.GetInt("server.work-pool-size")),
	}
}
func (m *MsgHandler) AddRouter(Id uint64, router rIterface.IRouter) {
	if _, ok := m.Apis[Id]; ok {
		panic("Id already exists,Id:" + strconv.Itoa(int(Id)))
	} else {
		m.Apis[Id] = router
		fmt.Println("id:", strconv.Itoa(int(Id)), "router add success")

	}

}
func (m *MsgHandler) DoMsgHandler(request rIterface.IRequest) {
	handler, ok := m.Apis[request.GetMessageID()]
	if !ok {
		fmt.Println("router has err,Id:", strconv.Itoa(int(request.GetMessageID())))
		return
	}
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)

}

func (m *MsgHandler) StartWorkPool() {
	for i := 0; i < int(m.WorkPoolSize); i++ {
		m.TaskQueue[i] = make(chan rIterface.IRequest, util.Config.GetInt("TaskQueueLen"))
		go m.StartTaskQueue(i, m.TaskQueue[i])
	}
}

func (m *MsgHandler) StartTaskQueue(workPoolId int, taskQueue chan rIterface.IRequest) {
	fmt.Println("start workPool ID:", workPoolId)
	for {
		select {
		case request := <-taskQueue:
			m.DoMsgHandler(request)
		}
	}
}

func (m *MsgHandler) SendMsgToTask(request rIterface.IRequest) {
	//将接受到的请求发送到连接池通道 ，将来做负载均衡 现在做平均分配
	workId := request.GetConnection().GetConnID() % m.WorkPoolSize
	m.TaskQueue[workId] <- request
}
