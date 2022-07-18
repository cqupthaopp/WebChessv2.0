package gameroom

import (
	"github.com/go-playground/assert/v2"
	"github.com/goccy/go-json"
	_ "github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestJoinRoomFunc(t *testing.T) {

	req := httptest.NewRequest(
		"POST",
		"/chess/moving",
		"Hao_pp",
	)

	w := httptest.NewRecorder()

	r := SetupRouter()

	assert.Equal(t, 200, w.Code)

	var resp map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &resp)

	assert.Nil(t, err)
	assert.Equal(t, "Hao_pp", resp["msg"])

}
