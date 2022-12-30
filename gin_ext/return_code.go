package gin_ext

// returnCode 响应状态码
type returnCode struct {
	code StatusCode
	msg  string
}

var (
	Success      = returnCode{code: 10000, msg: "请求成功!"}
	ParamsError  = returnCode{code: 30001, msg: "参数错误!"}
	RequestError = returnCode{code: 30004, msg: "请求错误!"}
)
