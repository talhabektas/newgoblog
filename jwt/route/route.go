package route

import (
	"awesomeProject2/jwt/controller"
	"awesomeProject2/jwt/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	// r.POST("/logout", controller.LogOut)
	// r.POST("/token", controller.Tokens)
	private := r.Group("/private")
	private.Use(middleware.Authenticate)
	private.GET("/token", controller.Tokens)
}
