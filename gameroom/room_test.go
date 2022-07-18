package gameroom

import "testing"

func TestInMap(t *testing.T) {

	get := InMap(123, 456)

	if get == true {
		t.Errorf("expected False,got True")
	}

}

func TestJudgeMovingVaild(t *testing.T) {

	get := JudgeMovingVaild(2, 3, 4, 5, nil)

	if get == true {
		t.Errorf("expected False,got True")
	}

}

func TestJudgeNum(t *testing.T) {

	get := JudgeNum("123", "456", "test", "123")

	if get == true {
		t.Errorf("expected False,got True")
	}

}

func TestInArea(t *testing.T) {
	get := InArea(5, 7, 0)

	if get == true {
		t.Errorf("expected False,got True")
	}
}
