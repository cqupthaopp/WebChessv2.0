/**
 * @Author: Hao_pp
 * @Data: 2022年7月17日-16点03分
 * @Desc: 无无无
 */

package gameroom

import (
	"WebChess/utils"
	"github.com/gin-gonic/gin"
)

func CreateRoomFunc(c *gin.Context) {

	info, err := utils.GetUserFromToken(c.GetHeader("Authorization"))

	if err != nil {
		c.JSON(403, utils.GetErrorRes(gin.H{"Response": "TokenError"}))
		return
	}

	password := c.PostForm("password")

	err = CreateRoom(info+"Room", info, password)
	PlayingGame[info+"Room"].Prepare[0] = 1

	if err != nil {
		c.JSON(403, utils.GetErrorRes(gin.H{"Info": "RoomExist"}))
		return
	}

	c.JSON(200, utils.GetNormalRes(gin.H{"RoomName": info + "Room"}))
	return

}
