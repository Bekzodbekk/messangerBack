package http

import (
	"api-gateway/internal/http/handlers"
	"api-gateway/internal/service"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewGin(cli service.ServiceRepositoryClient) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://messangervue.netlify.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"content-type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	r.Use(RedirectHTTPSMiddleware())

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

func RedirectHTTPSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("X-Forwarded-Proto") != "https" {
			c.Redirect(http.StatusMovedPermanently, "https://"+c.Request.Host+c.Request.RequestURI)
			return
		}
		c.Next()
	}
}
