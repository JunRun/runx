/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-10-30 15:49
 */
package rImpl

import (
	"fmt"
	"github.com/runx/server/rIterface"
	"strconv"
)

type MsgHandler struct {
	Apis map[uint64]rIterface.IRouter
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
	}
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)

}
