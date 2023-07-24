package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scaffold-demo-go/config"
	"scaffold-demo-go/utils/jwtauth"
	"scaffold-demo-go/utils/logs"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	userinfo := UserInfo{}
	returndata := config.ReturnData{}
	if err := c.ShouldBindJSON(&userinfo); err != nil {
		returndata.Status = 200500
		returndata.Message = err.Error()
		returndata.Data = map[string]interface{}{}
		c.JSON(http.StatusOK, returndata)
		return
	}
	//TODO 修改为读取数据库判断
	if userinfo.Username == config.Username && userinfo.Password == config.Password {
		token, err := jwtauth.GenerateJwt(userinfo.Username)
		if err != nil {
			returndata.Status = 200410
			returndata.Message = "生成token失败"
			returndata.Data = map[string]interface{}{}
			c.JSON(http.StatusOK, returndata)
			return
		}
		returndata.Status = 200200
		returndata.Message = "登陆成功"
		returndata.Data = map[string]interface{}{"token": token}
		c.JSON(http.StatusOK, returndata)
	} else {
		returndata.Status = 200410
		returndata.Message = "账号或密码不对"
		returndata.Data = map[string]interface{}{}
		c.JSON(http.StatusOK, returndata)
		return
	}

	logs.Deubg(map[string]interface{}{"username": userinfo.Username, "password": userinfo.Password}, "login ok")

	//returndata.Status = 200000
	//returndata.Message = ""
	//returndata.Data = map[string]interface{}{"username": userinfo.Username, "password": userinfo.Password}
	//c.JSON(http.StatusOK, returndata)
}

func Logout(c *gin.Context) {
	returnData := config.ReturnData{}
	returnData.Status = 200000
	returnData.Message = ""
	returnData.Data = make(map[string]interface{})
	c.JSON(200, returnData)
}
