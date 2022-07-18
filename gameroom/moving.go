/**
 * @Author: Hao_pp
 * @Data: 2022年7月17日-16点03分
 * @Desc: 无无无
 */

package gameroom

import (
	"WebChess/chess"
	"WebChess/utils"
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
	"time"
)

func MovingFunc(c *gin.Context) {

	roomName := c.PostForm("roomname")

	p, err := utils.GetUserFromToken(c.GetHeader("Authorization"))

	if err != nil {
		c.JSON(403, utils.GetErrorRes(gin.H{"Response": "TokenError"}))
		return
	}

	if time.Now().Unix()-PlayingGame[roomName].MovingMem.Unix() >= PlayingGame[roomName].OutOfTime {
		PlayingGame[roomName].MovingPlayer ^= 1 //超市了，换人
	}

	if PlayingGame[roomName].GameProcess == true {
		c.JSON(403, utils.GetErrorRes(gin.H{
			"Response": "游戏已经结束",
		}))
		return
	}

	if PlayingGame[roomName].PlayerName[PlayingGame[roomName].MovingPlayer] != p {
		c.JSON(403, utils.GetErrorRes(gin.H{
			"Response": "现在没到你移动棋子",
		}))
		return
	}

	x1 := c.PostForm("x")
	y1 := c.PostForm("y")

	x2 := c.PostForm("to_x")
	y2 := c.PostForm("to_y")

	if !JudgeNum(x1, y1, x2, y2) {
		c.JSON(403, utils.GetErrorRes(gin.H{
			"Response": "参数错误",
		}))
		return
	}

	x, _ := strconv.Atoi(x1)
	y, _ := strconv.Atoi(y1)

	to_x, _ := strconv.Atoi(x2)
	to_y, _ := strconv.Atoi(y2)

	board := &PlayingGame[roomName].ChessBoard

	if board[x][y]/10 != PlayingGame[roomName].MovingPlayer {
		c.JSON(403, utils.GetErrorRes(gin.H{
			"Response": "你操作的不是你的棋子",
		}))
		return
	}

	if JudgeMovingVaild(x, y, to_x, to_y, board) {

		if board[to_x][to_y]%10 == 1 {
			PlayingGame[roomName].GameOver()
			c.JSON(200, utils.GetNormalRes(gin.H{
				"winner": PlayingGame[roomName].PlayerName[PlayingGame[roomName].MovingPlayer],
			}))

			return
		}

		board[to_x][to_y] = board[x][y]
		board[x][y] = 0
	} else {
		c.JSON(403, utils.GetErrorRes(gin.H{
			"Response": "操作不合法",
		}))
		return
	}

	UpDateChessBoard(roomName, *board, string("{ "+x1+","+y1+"} -> "+"{ "+x2+","+y2+"} "))
	c.JSON(200, utils.GetNormalRes(gin.H{}))
	PlayingGame[roomName].MovingPlayer ^= 1

}

func InMap(x int, y int) bool {

	return x >= 0 && x < 10 && y >= 0 && y < 9

}

func JudgeMovingVaild(x int, y int, x2 int, y2 int, i *[10][9]int) bool {

	if !InMap(x, y) || !InMap(x2, y2) {
		return false
	}

	if i[x][y] == 0 {
		return false
	}

	t := chess.BlockNumFlag[i[x][y]%10]

	{
		if t == "兵" {
			return math.Abs(float64(x-x2))+math.Abs(float64(y-y2)) <= 1
		}
		if t == "车" {
			return (x == x2 || y == y2) && GetCount(x, y, x2, y2, i) == 0
		}

		if t == "炮" {
			cnt := GetCount(x, y, x2, y2, i)
			return (x == x2 || y == y2) && ((cnt == 1 && i[x2][y2] != 0) || (cnt == 0 && i[x2][y2] == 0))
		}

		if t == "帅" {
			return math.Abs(float64(x-x2))+math.Abs(float64(y-y2)) <= 1 && InArea(x2, y2, i[x][y]/10)
		}

		if t == "马" {
			if math.Abs(float64(x2-x)) == 1 && math.Abs(float64(y2-y)) == 2 && i[x2][y+1] == 0 {
				return true
			}
			if math.Abs(float64(x2-x)) == 2 && math.Abs(float64(y2-y)) == 1 && i[x+1][y] == 0 {
				return true
			}
			return false
		}

		if t == "象" {
			return math.Abs(float64(x2-x)) == 3 && math.Abs(float64(y2-y)) == 3
		}

		if t == "士" {
			return math.Abs(float64(x2-x)) == 1 && math.Abs(float64(y2-y)) == 1 && InArea(x2, y2, i[x][y]/10)
		}

	}

	return true
}

func InArea(x2 int, y2 int, s int) bool {

	if s == 1 {
		return x2 >= 7 && x2 <= 9 && y2 >= 4 && y2 <= 6
	}

	if s == 0 {
		return x2 >= 0 && x2 <= 2 && y2 >= 4 && y2 <= 6
	}

	return false

}

func GetCount(x int, y int, x2 int, y2 int, i *[10][9]int) int {

	cnt := 0

	for k := math.Min(float64(x), float64(x2)) + 1; k < math.Max(float64(x), float64(x2)); k++ {
		for p := math.Min(float64(y), float64(y2)) + 1; p < math.Max(float64(y), float64(y2)); p++ {
			if i[int(k)][int(p)] != 0 {
				cnt++
			}
		}
	}

	return cnt

}

func JudgeNum(x1 string, y1 string, x2 string, y2 string) bool {
	return true
} //判定数字
