/**
 * @Author: Hao_pp
 * @Data: 2022年7月17日-16点03分
 * @Desc: 无无无
 */

package gameroom

import (
	"WebChess/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func PrepareFunc(c *gin.Context) {

	roomName := c.PostForm("roomname")

	p, err := utils.GetUserFromToken(c.GetHeader("Authorization"))

	t := c.PostForm("PrepareType")

	if err != nil {
		c.JSON(403, utils.GetErrorRes(gin.H{
			"error": "Token error",
		}))
		return
	}

	if t != "1" && t != "2" {
		c.JSON(403, utils.GetErrorRes(gin.H{
			"error": "PrepareType Error",
		}))
		return
	}

	if PlayingGame[roomName].PlayerName[0] == p {
		PlayingGame[roomName].Prepare[0], _ = strconv.Atoi(t)
	}

	if PlayingGame[roomName].PlayerName[1] == p {
		PlayingGame[roomName].Prepare[1], _ = strconv.Atoi(t)
	}

	c.JSON(200, utils.GetNormalRes(gin.H{
		"Response": "OK!",
	}))

	if PlayingGame[roomName].Prepare[1] == 2 && PlayingGame[roomName].Prepare[0] == 2 {
		PlayingGame[roomName].StartGame()
	}

	return

}
