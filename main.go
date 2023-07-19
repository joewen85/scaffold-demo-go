package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"scaffold-demo-go/config"
	"scaffold-demo-go/middlewares/jwtauth"
)

func main() {
	r := gin.Default()
	token, _ := jwtauth.GenerateJwt("joe")
	claims, err := jwtauth.ParseJwt(token)
	fmt.Println(claims)
	fmt.Println(err)
	r.Run(config.Port)
}
