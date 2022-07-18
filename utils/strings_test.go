package utils

import "testing"

func TestMapToJson(t *testing.T) {

	get := MapToJson(map[string]interface{}{
		"test":    "Hao_pp",
		"testMsg": "TestMsg",
	})

	t.Skipf(get)

}