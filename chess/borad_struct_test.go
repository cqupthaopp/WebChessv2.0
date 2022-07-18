package chess

import "testing"

func TestNewChessBoard(t *testing.T) {
	get := NewChessBoard("123", "456")

	if get.ChessBoard[0][0] != 5 {
		t.Errorf("expected %d,got %d", 5, get.ChessBoard[0][0])
	}

}
