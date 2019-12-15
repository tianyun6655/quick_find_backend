package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"quick_find_backend/common"
	"quick_find_backend/user/grpc"
)


// @Tags User Service
// @Summary Login
// @Title Login
// @Param	phone  query 	string	true	"phone number"
// @Param	code  query 	int	true	"phone code"
// @Failure 800 params error
// @Failure 1000 system error
// @router  /quick_find/user/login [post]
func Login(c *gin.Context) {
	var(
		loginRequest=new(RegisterRequest)
	)

	if err:=ValidateParams(c,loginRequest);err!=nil{
		c.JSON(200,GetErrorResponse(800))
		return
	}

	user, err := controller.UserService.Login(
		context.Background(),
		&quick_find_services_user.LoginRequest{
		Phone: loginRequest.Phone,
		Code:  int32(loginRequest.Code),
	})

	if err != nil {
		c.JSON(200,GetErrorResponse_Message(err.Error()))
		return
	}

	c.JSON(200,GetResponse("success",user))
}


// @Tags User Service
// @Summary Register
// @Title Register
// @Param	phone  query 	string	true "phone number"
// @Param	code  query 	int	true	 "phone code"
// @Failure 800 params error
// @Failure 1000 system error
// @router  /quick_find/user/register [post]
func Register(c *gin.Context) {
	var (
		registerRequest = new(RegisterRequest)
	)

	if err := ValidateParams(c, registerRequest); err != nil {
		c.JSON(200, GetErrorResponse(800))
		return
	}

	user, err := controller.UserService.Register(
		context.Background(),
		&quick_find_services_user.RegisterRequest{
			Phone: registerRequest.Phone,
			Code:  int32(registerRequest.Code),
		})

	if err!=nil{
		common.Logger.Error("[c.user] Register(%#v) error: ",registerRequest,err)
		c.JSON(200, GetErrorResponse_Message(err.Error()))
		return
	}

	c.JSON(200,GetResponse("login success",user))
}



func Update(c *gin.Context){
   var(
   	 updateReqeust =  new(UpdateRequest)
   )
   if err := ValidateParams(c, updateReqeust); err != nil {
		c.JSON(200, GetErrorResponse(800))
		return
   }
	_, err := controller.UserService.Update(
		context.Background(),
		&quick_find_services_user.User{
			Uid:int32(updateReqeust.Uid),
			UserName:updateReqeust.Nickname,
			Age:int32(updateReqeust.Age),
			Sex:int32(updateReqeust.Sex),
	})
	if err!=nil{
		c.JSON(200,GetErrorResponse(1000))
		return
	}
	c.JSON(200,GetResponse("success",nil))
}

