package user

import (
	"github.com/gin-gonic/gin"
	"scaffold-demo-go/controllers/users"
)

func UserRouters(prefix *gin.RouterGroup) {
	prefix.POST("/user/login", users.Login)
	prefix.GET("/user/logout", users.Logout)
}
