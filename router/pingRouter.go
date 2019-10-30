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
	request.GetConnection().SendMsg(1, []byte("preKing"))
}

func (p *PingRouter) Handle(request rIterface.IRequest) {

}

func (p *PingRouter) PostHandle(request rIterface.IRequest) {

}
