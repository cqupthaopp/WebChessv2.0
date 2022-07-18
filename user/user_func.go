package User

/**
 * @Author: Hao_pp
 * @Data: 2022年7月17日-16点03分
 * @Desc: 无无无
 */

import "github.com/gin-gonic/gin"

type User struct {
	Username string `gorm:"primary_key"`
	Password string
}

func LoadUserFuncs(r *gin.Engine) {

	r.POST("/user", UserRegisterFunc)
	r.GET("/user", UserLoginFunc)

}
