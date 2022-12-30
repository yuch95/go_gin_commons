package gin_ext

import "github.com/gin-gonic/gin"

// Namespace 命名空间 路由组、蓝图
type Namespace struct {
	Name       string
	Middleware []MiddlewareHandler
	group      *gin.RouterGroup
}

// NewNamespace Namespace的构造函数 创建新的命名空间
func NewNamespace(app *gin.Engine, name string, path string, middle ...MiddlewareHandler) *Namespace {
	group := app.Group(path)
	return &Namespace{Name: name, Middleware: middle, group: group}
}

// NewNamespace 创建新的子命名空间
func (n *Namespace) NewNamespace(name string, path string, middle ...MiddlewareHandler) *Namespace {
	group := n.group.Group(path)
	newMiddle := append(n.Middleware, middle...)
	return &Namespace{Name: name, Middleware: newMiddle, group: group}
}

// createHandlerFunc 请求处理函数解析 构造gin的处理函数
func (n *Namespace) createHandlerFunc(handler RequestHandler, middle ...MiddlewareHandler) gin.HandlerFunc {
	// 经请求处理函数转换成 自定义中间件函数 并将返回结果JSON序列化
	requestTrance := func(r *Resource) {
		resp := handler(r)
		if resp != nil {
			r.JSON(200, resp)
		}
	}

	// 构建新的函数中间件处理数组 处理函数放在最后一个执行
	newMiddle := append(append(n.Middleware, middle...), requestTrance)
	// 创建gin的处理函数 并在其中创建自定义资源函数 并开始执行自定义中间件
	return func(c *gin.Context) {
		resource := &Resource{handlers: newMiddle, Context: c}
		resource.begin()
	}
}

func (n *Namespace) Get(path string, handler RequestHandler, middle ...MiddlewareHandler) {
	n.group.GET(path, n.createHandlerFunc(handler, middle...))
}

func (n *Namespace) Put(path string, handler RequestHandler, middle ...MiddlewareHandler) {
	n.group.PUT(path, n.createHandlerFunc(handler, middle...))
}

func (n *Namespace) Post(path string, handler RequestHandler, middle ...MiddlewareHandler) {
	n.group.POST(path, n.createHandlerFunc(handler, middle...))
}

func (n *Namespace) Patch(path string, handler RequestHandler, middle ...MiddlewareHandler) {
	n.group.PATCH(path, n.createHandlerFunc(handler, middle...))
}

func (n *Namespace) Delete(path string, handler RequestHandler, middle ...MiddlewareHandler) {
	n.group.DELETE(path, n.createHandlerFunc(handler, middle...))
}

func (n *Namespace) Handle(httpMethod string, path string, handler RequestHandler, middle ...MiddlewareHandler) {
	n.group.Handle(httpMethod, path, n.createHandlerFunc(handler, middle...))
}
