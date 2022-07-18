/**
 * @Author: Hao_pp
 * @Data: 2022年7月17日-16点03分
 * @Desc: 无无无
 */

package main

import (
	"WebChess/User"
	"WebChess/chess"
	"WebChess/gameroom"
	"github.com/gin-gonic/gin"
)

func main() {

	{
		chess.InitBlock()
	} //init

	r := gin.Default()

	{
		User.LoadUserFuncs(r)
		gameroom.LoadRoomFuncs(r)
	} //LoadFunc

	r.Run(":8080")

}
