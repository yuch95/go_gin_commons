package libs

type StatusCode int

// Jsonify 标准Json响应
type Jsonify struct {
	Code StatusCode `json:"Code"`
	Msg  string     `json:"Msg"`
	Data any        `json:"data"`
}

// SuccessResp 构建成功响应
func (j *Jsonify) SuccessResp(data any) *Jsonify {
	j.Code = Success.Code
	j.Msg = Success.Msg
	j.Data = data
	return j
}

// ErrorResp 构建失败响应
func (j *Jsonify) ErrorResp(returnCode returnCode, msg ...string) *Jsonify {
	j.Code = returnCode.Code
	if len(msg) > 0 {
		for _, s := range msg {
			j.Msg += s
		}
	} else {
		j.Msg = returnCode.Msg
	}
	return j
}
