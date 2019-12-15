package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"quick_find_backend/friend/grpc"
)



// @Tags Friend Service
// @Summary Add
// @Title Add
// @Param	uid  query 	string	true	"the user id"
// @Param	friend_id  query 	int	true	"want to add"
// @Failure 800 params error
// @Failure 1000 system error
// @router  /quick_find/friend/add [post]
func Add(c *gin.Context){
    var(
    	addFriendRequest=new(AddFriendRequest)
	)

    if err:=ValidateParams(c,addFriendRequest);err!=nil{
    	c.JSON(200,GetErrorResponse(800))
		return
	}
    _,err:=controller.FriendService.AddFriend(
    	context.TODO(), &quick_find_services_friend.AddFriendRequest{
    	Uid:addFriendRequest.Uid,
    	FriendId:addFriendRequest.FriendId,
	})

    if err!=nil{
		c.JSON(200,GetErrorResponse(1000))
		return
	}

	c.JSON(200,GetResponse("success",nil))
}

// @Tags Friend Service
// @Summary List
// @Title List
// @Param	uid  query 	string	true	"the user id"
// @Failure 800 params error
// @Failure 1000 system error
// @router  /quick_find/friend/list [post]
func FriendList(c *gin.Context){
	var(
		listFriendRequest=new(FriendListRequest)
	)
	if err:=ValidateParams(c,listFriendRequest);err!=nil{
		c.JSON(200,GetErrorResponse(800))
		return
	}
	list,err:=controller.FriendService.List(
		context.TODO(), &quick_find_services_friend.ListRequest{
			Uid:listFriendRequest.Uid,
		})
	if err!=nil{
		c.JSON(200,GetErrorResponse(1000))
		return
	}
	c.JSON(200,GetResponse("success",list))

}