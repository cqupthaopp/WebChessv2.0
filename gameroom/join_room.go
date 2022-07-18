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

func JoinRoomFunc(c *gin.Context) {

	roomName := c.Param("name")

	if _, ok := BoardCh[roomName]; !ok {
		c.JSON(401, gin.H{
			"status": 10003,
			"info":   "No Room",
		})
		return
	}

	info, err := utils.GetUserFromToken(c.GetHeader("Authorization"))

	if err != nil {
		c.JSON(403, utils.GetErrorRes(gin.H{"Response": "TokenError"}))
		return
	}

	PlayingGame[roomName].Prepare[1] = 1
	PlayingGame[roomName].PlayerName[1] = info
	conn, _ := upGrader.Upgrade(c.Writer, c.Request, nil)
	room.Store(info, conn)
	BroadCast(roomName)
	defer conn.Close()

	c.JSON(200, gin.H{})

	return

}
