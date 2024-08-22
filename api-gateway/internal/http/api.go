package http

import (
	"api-gateway/internal/http/handlers"
	"api-gateway/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "api-gateway/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewGin(cli service.ServiceRepositoryClient) *gin.Engine {
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowBrowserExtensions = true
	corsConfig.AllowMethods = []string{"*"}
	r.Use(cors.New(corsConfig))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	hnd := handlers.NewHandlerSt(cli)
	r.POST("/auth/register", hnd.Register)
	r.POST("/auth/verify", hnd.Verify)
	r.POST("/auth/login", hnd.SignIn)

	r.PUT("/user/:id", hnd.UpdateUser)
	r.DELETE("/user/:id", hnd.DeleteUser)
	r.GET("/user", hnd.GetUsers)
	r.GET("/user/:id", hnd.GetUserById)
	r.GET("/user/filter", hnd.GetUserByFilter)
	r.GET("/directs", hnd.GetAllDirects)

	r.POST("/message", hnd.CreateMessage)
	r.PUT("/message/:id")
	r.DELETE("/message")
	r.GET("/message", hnd.GetMessagesByTo)

	r.GET("/ws", hnd.WsHandler)
	return r
}
