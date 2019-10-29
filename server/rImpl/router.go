/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-10-29 13:30
 */
package rImpl

import "github.com/runx/server/rIterface"

type BaseRouter struct {
}

func (b *BaseRouter) PreHandle(request rIterface.IRequest) {
}
func (b *BaseRouter) Handle(request rIterface.IRequest) {

}
func (b *BaseRouter) PostHandle(request rIterface.IRequest) {

}
