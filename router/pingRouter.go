/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-10-29 14:30
 */
package router

import (
	"fmt"
	"github.com/runx/server/rImpl"
	"github.com/runx/server/rIterface"
)

type PingRouter struct {
	rImpl.BaseRouter
}

func (p *PingRouter) PreHandle(request rIterface.IRequest) {
	fmt.Println("PreHandle start")
	fmt.Printf("id:%d recv :%s\n", request.GetMessageID(), request.GetData())
	request.GetConnection().SendMsg(0, []byte("preKing"))
}

func (p *PingRouter) Handle(request rIterface.IRequest) {

}

func (p *PingRouter) PostHandle(request rIterface.IRequest) {

}

type Hell struct {
	rImpl.BaseRouter
}

func (h *Hell) PreHandle(request rIterface.IRequest) {
	fmt.Println("hell PreHandle start")
	fmt.Printf("id:%d recv :%s\n", request.GetMessageID(), request.GetData())
	request.GetConnection().SendMsg(1, []byte("ll"))
}

func (h *Hell) Handle(request rIterface.IRequest) {

}

func (h *Hell) PostHandle(request rIterface.IRequest) {

}
