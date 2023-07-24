package routers

import (
	"github.com/gin-gonic/gin"
	"scaffold-demo-go/routers/user"
)

func RegisterRouters(r *gin.Engine) {
	prefix := r.Group("/api")
	user.UserRouters(prefix)
}
