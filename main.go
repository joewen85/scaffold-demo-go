package main

import (
	"github.com/gin-gonic/gin"
	"scaffold-demo-go/config"
)

func main() {
	r := gin.Default()
	r.Run(config.Port)
}
