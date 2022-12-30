package gin_ext

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"tools.com/libs/libs"
)

func validateErr(errs validator.ValidationErrors, c *gin.Context) {
	errMsgMap := make(map[string]string)
	for _, e := range errs {
		name := e.StructField()
		msg := e.Error()
		errMsgMap[name] = msg
	}
	c.JSON(200, &libs.Jsonify{Code: libs.ParamsError.Code, Msg: libs.ParamsError.Msg, Data: errMsgMap})
}

// Recover 全局异常捕获中间件
func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case validator.ValidationErrors:
				validateErr(err.(validator.ValidationErrors), c)
			default:
				c.JSON(200, &libs.Jsonify{Code: libs.RequestError.Code, Msg: "系统开小差!"})
			}
		}
	}()
	c.Next()
}
