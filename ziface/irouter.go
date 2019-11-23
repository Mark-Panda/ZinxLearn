package ziface

/**
	路由的抽象接口
	路由里的数据都是IRequest请求
 */

type IRouter interface {
	//在处理conn业务之前的钩子方法HooK
	PreHandle(request IRequest)
	//在处理conn业务的主方法
	Handle(request IRequest)
	//在处理conn业务之后的方法
	PostHandle(request IRequest)
}