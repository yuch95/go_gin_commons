package gin_ext

type StatusCode int

// Jsonify 标准Json响应
type Jsonify struct {
	Code StatusCode `json:"code"`
	Msg  string     `json:"msg"`
	Data any        `json:"data"`
}

// SuccessResp 构建成功响应
func SuccessResp(data any) *Jsonify {
	code := Success.code
	msg := Success.msg
	return &Jsonify{Code: code, Msg: msg, Data: data}
}

// ErrorResp 构建失败响应
func ErrorResp(returnCode returnCode, msg ...string) *Jsonify {
	code := returnCode.code
	var newMsg string
	if len(msg) > 0 {
		for _, s := range msg {
			newMsg += s
		}
	} else {
		newMsg = returnCode.msg
	}
	return &Jsonify{Code: code, Msg: newMsg}
}
