/**
 * @Author: Hao_pp
 * @Data: 2022年7月17日-16点03分
 * @Desc: 无无无
 */

package gameroom

import (
	"WebChess/chess"
	"WebChess/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
	"unsafe"
)

type ChessBoard struct {
	ChessBoard [10][9]int
	time       time.Time
	op         string
}

var upGrader = websocket.Upgrader{

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var lock = sync.Mutex{}
var BoardCh map[string]chan ChessBoard
var room = sync.Map{}
var PlayingGame map[string]*chess.ChessGameRoom

func UpDateChessBoard(roomName string, board [10][9]int, op string) {

	BoardCh[roomName] <- ChessBoard{
		ChessBoard: board,
		time:       time.Now(),
		op:         op,
	} //

}

func BroadCast(roomName string) {

	for {

		var err error

		select {

		case Board := <-BoardCh[roomName]:

			if Board.time.Unix()-time.Now().Unix() >= 50000 {
				continue
			}

			m := append(S2B(utils.MapToJson(gin.H{
				"ChessBoard": Board.ChessBoard,
				"Operator":   Board.op,
			})))

			room.Range(func(key, value any) bool {
				conn := value.(*websocket.Conn)
				err = conn.WriteMessage(websocket.TextMessage, m)
				if err != nil {
					log.Println("conn.WriteMessage err: ", err)
				}
				return true
			})
		}
	}

}

func S2B(str string) (bytes []byte) {
	x := *(*[2]uintptr)(unsafe.Pointer(&str))
	bytes = *(*[]byte)(unsafe.Pointer(&[3]uintptr{x[0], x[1], x[1]}))
	return
}

func CreateRoom(roomName string, p1 string, password string) error {

	if len(BoardCh) == 0 {
		BoardCh = make(map[string]chan ChessBoard)
	}

	if len(PlayingGame) == 0 {
		PlayingGame = make(map[string]*chess.ChessGameRoom)
	}

	if _, ok := BoardCh[roomName]; ok {
		return errors.New("Room exist")
	}

	BoardCh[roomName] = make(chan ChessBoard, 500) //make chan
	PlayingGame[roomName] = chess.NewChessBoard(p1, password)

	return nil
}
