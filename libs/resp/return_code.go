package resp

// returnCode 响应状态码
type returnCode struct {
	Code StatusCode
	Msg  string
}

var (
	Success      = returnCode{Code: 10000, Msg: "请求成功!"}
	ParamsError  = returnCode{Code: 30001, Msg: "参数错误!"}
	RequestError = returnCode{Code: 30004, Msg: "请求错误!"}
)
