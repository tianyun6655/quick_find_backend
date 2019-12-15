package http

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"quick_find_backend/common"
)

var errors map[int]string

func init() {
	errors = map[int]string{
		800: "params error",
		900: "register failed",
		1000:"update user information failed",
		1001:"验证码错误",
		1002:"请先填写企业额外信息",
	}
}

type PageRequest struct {
	Page int 	`json:"page"`
	Count int `json:"count"`
}

type PageRequest2 struct {
	Page int 	`json:"page"`
	Count int `json:"limit"`
}

type PageResponse struct {
	List  interface{} `json:"list"`
	Total int  `json:"total"`
}
type PageResponseSpecail struct {
	List  interface{} `json:"list"`
	Total int  `json:"total"`

}
func ValidateParams(c *gin.Context, param interface{}) error {
	validate := validator.New()
	err := c.BindJSON(param)
	if err != nil {
		goto SKIP
	}
	err = validate.Struct(param)
SKIP:
	if err != nil {
		common.Logger.Error("[gateway.middleware] ValidateParams(%s) error:%s",c.Request.URL,err.Error())
	}
	return err
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetErrorResponse(code int) *Response {

	return &Response{
		Code:    code,
		Message: errors[code],
		Data:    nil,
	}
}
func GetErrorResponse_Message(message string) *Response {

	return &Response{
		Code:    400,
		Message: message,
		Data:    nil,
	}
}


func GetResponse(Message string, data interface{}) *Response {
	response := &Response{
		Code:    200,
		Message: Message,
		Data:    data,
	}
	return response
}




type IdRequest struct {
	Id int  `json:"id"`
}
