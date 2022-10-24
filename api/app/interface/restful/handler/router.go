package handler

import (
	"log"
	"server/app/interface/restful/handler/endpoints"
	"server/app/interface/restful/handler/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	log.Println("Server Handler Started Successfully")
}

func Routing() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://ducthang.dev", "http://localhost:5001", "http://128.199.202.198", "https://128.199.202.198", "https://peekaboo.ducthang.dev"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Origin", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// r.Use(middleware.RateLimit(2, 5))
	r.Use(middleware.Tracer())

	r.POST("/api/login", endpoints.Login)
	r.POST("/api/register", endpoints.CreateUser)
	r.POST("/api/logout", middleware.ValidateToken(), endpoints.Logout)

	v1 := r.Group("/api/v1")
	v1.Use(middleware.ValidateToken())

	v1.GET("/validateRole", middleware.ValidateToken(), endpoints.ValidateRole)
	v1.GET("/ID", endpoints.GetID)
	endpoints.UserHandler(v1.Group("/user"))
	endpoints.TestHandler(v1.Group("/test"))
	endpoints.ClassHandler(v1.Group("/class"))
	endpoints.AdminHandler(v1.Group("/admin"))

	return r
}
