package gin_ext

import (
	"github.com/gin-gonic/gin"
	"tools.com/libs/libs/resp"
)

// MiddlewareHandler 中间件方法函数
type MiddlewareHandler func(c *Context)

// RequestHandlerFunc 请求方法函数
type RequestHandlerFunc func(c *Context) *resp.Jsonify

type RequestHandler interface {
	Handler(c *Context) *resp.Jsonify
}

// Context 基于资源的上下文
// 可以在Context中添加任意资源 用户信息 数据库对象 等
type Context struct {
	index    int8
	handlers []MiddlewareHandler
	*gin.Context
}

// begin 开始运行资源中的方法
func (r *Context) begin() {
	for r.index < int8(len(r.handlers)) {
		r.handlers[r.index](r)
		r.index++
	}
}

// Next 中间件中执行下个方法
func (r *Context) Next() {
	r.index++
	for r.index < int8(len(r.handlers)) {
		r.handlers[r.index](r)
		r.index++
	}
}
