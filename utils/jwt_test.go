package utils

import "testing"

func TestGetToken(t *testing.T) {

	get, _ := GetToken("Hao_pp")

	if s, _ := PraseToken(get); s.Username != "Hao_pp" {
		t.Errorf("expected Hao_pp,got %s", s.Username)
	}

}

func TestGetRefreshToken(t *testing.T) {

	get, _ := GetRefreshToken("Hao_pp")

	if s, _ := PraseToken(get); s.Username != "Hao_pp" {
		t.Errorf("expected Hao_pp,got %s", s.Username)
	}

}
