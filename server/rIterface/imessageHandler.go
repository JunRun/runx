/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-10-30 15:45
 */
package rIterface

type IMsgHandler interface {
	AddRouter(Id uint64, router IRouter)
	DoMsgHandler(request IRequest)
}
