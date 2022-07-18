package User

/**
 * @Author: Hao_pp
 * @Data: 2022年7月17日-16点03分
 * @Desc: 无无无
 */

import (
	"WebChess/utils"
	"WebChess/utils/mysql"
	"github.com/gin-gonic/gin"
)

func UserRegisterFunc(c *gin.Context) {

	user := mysqlutil.User{c.PostForm("username"), c.PostForm("password")}

	if err := mysqlutil.FindUser(user.Username); err == nil {
		c.JSON(405, utils.GetErrorRes(gin.H{
			"response": "User " + user.Username + " Exist",
		}))
		return
	}

	err := mysqlutil.AddUser(user)

	if err != nil {
		c.JSON(405, utils.GetErrorRes(gin.H{
			"response": "Register Failed",
		}))
	} else {
		token, token_err := utils.GetToken(user.Username)
		retoken, retoken_err := utils.GetRefreshToken(user.Username)

		if token_err != nil || retoken_err != nil {
			c.JSON(405, utils.GetErrorRes(gin.H{
				"response": "GetTokenError",
			}))
		}

		c.JSON(200, utils.GetNormalRes(gin.H{
			"token":         token,
			"refresh_token": retoken,
			"username":      user.Username,
		}))
	}

}
