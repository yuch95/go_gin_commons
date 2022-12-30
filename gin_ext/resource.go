package gin_ext

import "github.com/gin-gonic/gin"

// MiddlewareHandler 中间件方法函数
type MiddlewareHandler func(c *Resource)

// RequestHandlerFunc 请求方法函数
type RequestHandlerFunc func(c *Resource) *Jsonify

type RequestHandler interface {
	Handler(c *Resource) *Jsonify
}

// Resource 基于资源的上下文
// 可以在Resource中添加任意资源 用户信息 数据库对象 等
type Resource struct {
	index    int8
	handlers []MiddlewareHandler
	*gin.Context
	*Jsonify
}

// begin 开始运行资源中的方法
func (r *Resource) begin() {
	for r.index < int8(len(r.handlers)) {
		r.handlers[r.index](r)
		r.index++
	}
}

// Next 中间件中执行下个方法
func (r *Resource) Next() {
	r.index++
	for r.index < int8(len(r.handlers)) {
		r.handlers[r.index](r)
		r.index++
	}
}
