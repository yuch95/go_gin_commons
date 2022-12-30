# 自用封装go工具

## gin_ext 
> 对gin框架中上下文进行再封装
### MiddlewareHandler

  >中间件函数方法类型 
### RequestHandler

  >请求函数的方法类型
### Resource

  > 继承gin.Context上下文，拥有原上下文所有方法，并重写了Next方法，使其可以调用自定义资源

```go
// Resource 基于资源的上下文
// 可以在Resource中添加任意资源 用户信息 数据库对象 等
type Resource struct {
	index    int8
	handlers []MiddlewareHandler
	*gin.Context
	*Jsonify
	UserInfo *UserInfo
	DbDao    *DbDao
	...
}
```



***Tips:该结构体在使用中按需添加自己需要的资源***

### Namespace

  > 对gin.RouterGroup进行的再封装，可以用于创建新的路由组。用于路由模块化

创建gin处理函数的逻辑：

1.  将请求函数转成中间件函数 并将请求函数的返回数据进行JSON返回
2.  合并中间件。将创建路由组的中间件和改方法的中间件进行合并，并在最后增加转换后的请求函数
3.  返回gin的请求方法。并在其中创建自己的资源对象，开始执行自定义的中间件函数

```go
// createHandlerFunc 请求处理函数解析 构造gin的处理函数
func (n *Namespace) createHandlerFunc(handler RequestHandler, middle ...MiddlewareHandler) gin.HandlerFunc {
	// 1.经请求处理函数转换成 自定义中间件函数 并将返回结果JSON序列化
	requestTrance := func(r *Resource) {
		resp := handler(r)
		if resp != nil {
			r.JSON(200, resp)
		}
	}

	// 2.构建新的函数中间件处理数组 处理函数放在最后一个执行
	newMiddle := append(append(n.Middleware, middle...), requestTrance)
	// 3.创建gin的处理函数 并在其中创建自定义资源函数 并开始执行自定义中间件
	return func(c *gin.Context) {
		resource := &Resource{handlers: newMiddle, Context: c}
		resource.begin()
	}
}
```



### Response

>   JSON响应结构规范化