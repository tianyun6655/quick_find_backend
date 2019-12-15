package http

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "quick_find_backend/docs"
	"quick_find_backend/friend/grpc"
	"quick_find_backend/im/grpc"
	"quick_find_backend/topic/grpc"
	"quick_find_backend/user/grpc"
)

var controller *Controller

type Controller struct {
	UserService   quick_find_services_user.UserServicesService
	TopicService  quick_find_services_topic.TopicServicesService
	FriendService quick_find_services_friend.FriendServicesService
	IMServices    quick_find_services_im.IMServicesService
}

//init all the services - Arvin
func Init(service micro.Service, port string) {
	controller = &Controller{
		UserService: quick_find_services_user.NewUserServicesService(
			"quick.find.services.user",
			service.Client()),
		TopicService: quick_find_services_topic.NewTopicServicesService(
			"quick.find.services.topic",
			service.Client()),
		FriendService: quick_find_services_friend.NewFriendServicesService(
			"quick.find.services.friend",
			service.Client()),
		IMServices: quick_find_services_im.NewIMServicesService(
			"quick.find.services.im",
			service.Client()),
	}
	Start(port)
}

func Start(port string) {
	//init the router
	router := gin.Default()

	quickFind := router.Group("/quick_find")
     quickFind.Use(GlobalAuthMiddleware())
	//user api
	user := quickFind.Group("/user")
	{
		//login api
		user.Any("/login", Login)
		user.Any("/register", Register)
		user.Any("/update", Update)

	}
	friend := quickFind.Group("/friend")
	{
		friend.Any("/add",Add)
		friend.Any("/list",FriendList)
	}
	topic := quickFind.Group("/topic")
	{
		topic.Any("/publish", Publish)
		topic.Any("/list", List)
	}

	im := quickFind.Group("/im")
	{
		im.Any("/send", SendMessage)
		im.Any("/translate", Translate)

	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(port)
}
