package gameroom

import "github.com/gin-gonic/gin"

/**
 * @Author: Hao_pp
 * @Data: 2022年7月17日-16点03分
 * @Desc: 无无无
 */

func LoadRoomFuncs(r *gin.Engine) {

	r.POST("/chess/create", CreateRoomFunc)

	r.GET("/chess/join/:name", JoinRoomFunc)

	r.POST("/chess/moving", MovingFunc)

	r.POST("/chess/prepare", PrepareFunc)

}
