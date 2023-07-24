package jwtvalid

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scaffold-demo-go/config"
	"scaffold-demo-go/utils/jwtauth"
	"scaffold-demo-go/utils/logs"
)

func JwtValid(c *gin.Context) {
	returnData := config.ReturnData{}
	requestURL := c.FullPath()
	logs.Deubg(map[string]interface{}{"请求地址": requestURL}, "")
	if requestURL == "/api/user/logout" || requestURL == "/api/user/login" {
		c.Next()
		return
	}
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		returnData.Status = 200401
		returnData.Message = "无登陆"
		returnData.Data = map[string]interface{}{}
		c.JSON(http.StatusOK, returnData)
		c.Abort()
		return
	}
	claims, err := jwtauth.ParseJwt(token)
	if err != nil {
		returnData.Status = 200401
		returnData.Message = "token不合法"
		returnData.Data = map[string]interface{}{}
		c.JSON(http.StatusOK, returnData)
		c.Abort()
		return
	}
	c.Set("claims", claims)
	c.Next()
}
