package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"quick_find_backend/im/grpc"
)


// @Tags IM Service
// @Summary send
// @Title send
// @Param	from  query 	int	true	"from id"
// @Param	to  query 	int	true	"to id"
// @Param	message  query 	string	true	"message"
// @Failure 800 params error
// @Failure 1000 system error
// @router  /quick_find/im/send [post]
func SendMessage(c *gin.Context) {
	var (
		sendRequest = new(SendMessageRequest)
	)
	if err := ValidateParams(c, sendRequest); err != nil {
		c.JSON(200, GetErrorResponse(800))
	}
	_, err := controller.IMServices.SendMessage(context.TODO(),
		&quick_find_services_im.SendMessageRequest{
		From:    sendRequest.From,
		To:      sendRequest.To,
		Message: sendRequest.Message,
	})
	if err != nil {
		c.JSON(200, GetErrorResponse_Message(err.Error()))
		return
	}
	c.JSON(200, GetErrorResponse(800))
	return

}

// @Tags IM Service
// @Summary translate
// @Title translate
// @Param	text  query 	string	true	"text"
// @Failure 800 params error
// @Failure 1000 system error
// @router  /quick_find/im/translate [post]
func Translate(c *gin.Context){
	 var(
	 	translate=new(TranslateRequest)
	 )
	 if err:=ValidateParams(c,translate);err!=nil{
		 c.JSON(200, GetErrorResponse(800))
	 }
	result, err :=controller.IMServices.Translate(context.TODO(),&quick_find_services_im.TranslateRequest{
		Text:translate.Text,
	})
	if err != nil {
		c.JSON(200, GetErrorResponse_Message(err.Error()))
		return
	}
	 	c.JSON(200, GetResponse("success",result))
	return
}
