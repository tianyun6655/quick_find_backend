package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"quick_find_backend/topic/grpc"
)

// @Tags Topic Service
// @Summary publish
// @Title publish
// @Param	content  query 	string	true	"content of the topic"
// @Param	title  query 	string	true	"title of the topic"
// @Param	uid  query 	int	true	"want to add"
// @Failure 800 params error
// @Failure 1000 system error
// @router  /quick_find/topic/publish [post]
func Publish(c *gin.Context){
	var(
       publishRequest =new(PublishRequest)
	)

	if err:=ValidateParams(c,publishRequest);err!=nil{
		c.JSON(200, GetErrorResponse(800))
		return
	}
	publish:=&quick_find_services_topic.Publish{
		Title:publishRequest.Title,
		Content:publishRequest.Content,
		Uid:publishRequest.Uid,
	}
	_,err:=controller.TopicService.Push(context.TODO(),publish)
	if err!=nil{
		c.JSON(200, GetErrorResponse(1000))
		return
	}
	c.JSON(200,GetResponse("success",nil))
	return
}


// @Tags Topic Service
// @Summary publish
// @Title publish
// @Param	page  query 	int	true	"page"
// @Param	count  query 	int	true	"count"
// @Param	uid  query 	int	true	"want to add"
// @Failure 800 params error
// @Failure 1000 system error
// @router  /quick_find/topic/list [post]
func List(c *gin.Context){
	 var(
	 	pageRequest =new(PageRequest)
	 )
	 if err:=ValidateParams(c,pageRequest);err!=nil{
		 c.JSON(200, GetErrorResponse(800))
		 return
	 }
	 list:=&quick_find_services_topic.ListRequest{
	 	Page:int32(pageRequest.Page),
	 	Count:int32(pageRequest.Count),
	 }
	 response,err:=controller.TopicService.List(context.TODO(),list)
	 if err!=nil{
		 c.JSON(200, GetErrorResponse(1000))
		 return
	 }
	c.JSON(200,GetResponse("success",response))
	return
}