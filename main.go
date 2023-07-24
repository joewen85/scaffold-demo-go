package main

import (
	"github.com/gin-gonic/gin"
	"scaffold-demo-go/config"
	"scaffold-demo-go/middlewares/jwtvalid"
	"scaffold-demo-go/routers"
)

func main() {
	r := gin.Default()
	r.Use(jwtvalid.JwtValid)
	routers.RegisterRouters(r)
	r.Run(config.Port)
}
