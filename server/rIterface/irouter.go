/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-10-29 13:29
 */
package rIterface

type IRouter interface {
	//处理conn连接之前的业务方法 hook
	PreHandle(request IRequest)
	//处理主业务方法
	Handle(request IRequest)
	//处理业务之后的方法
	PostHandle(request IRequest)
}
