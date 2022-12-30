package gin_ext

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func validateErr(errs validator.ValidationErrors, c *gin.Context) {
	errMsgMap := make(map[string]string)
	for _, e := range errs {
		name := e.StructField()
		msg := e.Error()
		errMsgMap[name] = msg
	}
	c.JSON(200, &Jsonify{Code: ParamsError.code, Msg: ParamsError.msg, Data: errMsgMap})
}

// Recover 全局异常捕获中间件
func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case validator.ValidationErrors:
				validateErr(err.(validator.ValidationErrors), c)
			default:
				c.JSON(200, &Jsonify{Code: RequestError.code, Msg: "系统开小差!"})
			}
		}
	}()
	c.Next()
}
