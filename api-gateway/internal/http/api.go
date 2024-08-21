package http

import (
	"api-gateway/internal/http/handlers"
	"api-gateway/internal/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewGin(cli service.ServiceRepositoryClient) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                                     // Ruxsat berilgan domennar
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                   // Ruxsat berilgan metodlar
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"}, // Ruxsat berilgan headerlar
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour, // CORS qoidalarini cache qilish vaqti
	}))

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
