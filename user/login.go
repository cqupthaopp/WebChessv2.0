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

func UserLoginFunc(c *gin.Context) {

	LoginUser := mysqlutil.User{c.PostForm("username"), c.PostForm("password")}

	if err := mysqlutil.FindUser(LoginUser.Username); err != nil {
		c.JSON(405, utils.GetErrorRes(gin.H{
			"response": "User " + LoginUser.Username + "is not Exist",
		}))
		return
	}

	UserInfo := mysqlutil.GetUser(LoginUser.Username)

	if UserInfo.Password != LoginUser.Password {
		c.JSON(403, utils.GetErrorRes(gin.H{
			"response": "Password is not match",
		}))
	} else {

		token, token_err := utils.GetToken(LoginUser.Username)
		retoken, retoken_err := utils.GetRefreshToken(LoginUser.Username)

		if token_err != nil || retoken_err != nil {
			c.JSON(405, utils.GetErrorRes(gin.H{
				"response": "GetTokenError",
			}))
			return
		}

		c.JSON(200, utils.GetNormalRes(gin.H{
			"token":         token,
			"refresh_token": retoken,
			"username":      LoginUser.Username,
		}))

	}

}
