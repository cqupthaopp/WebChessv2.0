/**
 * @Author: Hao_pp
 * @Data: 2022年7月17日-16点03分
 * @Desc: 无无无
 */

package chess

import (
	"time"
)

type ChessGameRoom struct {
	ChessBoard [10][9]int

	PlayerName [2]string
	OutOfTime  int64

	MovingMem    time.Time //上一次移动的时间
	MovingPlayer int
	GameProcess  bool

	password string
	Prepare  [2]int
}

func (r ChessGameRoom) StartGame() {

	r.MovingMem = time.Now()
	r.MovingPlayer = 0
	r.GameProcess = false

}

func NewChessBoard(player1 string, password string) *ChessGameRoom {

	var room ChessGameRoom

	room.password = password

	room.PlayerName[0] = player1
	room.Prepare[0] = 0
	room.Prepare[0] = 0

	{
		for v := 1; v <= 5; v++ {
			room.ChessBoard[0][4+v-1] = v
			room.ChessBoard[0][4-v+1] = v
			room.ChessBoard[9][4+v-1] = v + 10
			room.ChessBoard[9][4-v+1] = v + 10
		}
		for k := 0; k < 9; k += 2 {
			room.ChessBoard[3][k] = BlockFlag["兵"]
			room.ChessBoard[6][k] = BlockFlag["兵"] + 10
		}
		room.ChessBoard[2][1] = BlockFlag["炮"]
		room.ChessBoard[2][7] = BlockFlag["炮"]
		room.ChessBoard[7][1] = BlockFlag["炮"] + 10
		room.ChessBoard[7][7] = BlockFlag["炮"] + 10
	} //initBoard

	room.OutOfTime = 50

	return &room

}

func (c ChessGameRoom) GameOver() {

	c.GameProcess = true

}
