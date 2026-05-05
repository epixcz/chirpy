package auth

import "testing"

func TestMakeRefreshToken(t *testing.T) {
	token := MakeRefreshToken()

	if len(token) != 64 {
		t.Fatalf("got token length %d, want 64", len(token))
	}
}

func TestMakeRefreshTokenReturnsUniqueTokens(t *testing.T) {
	firstToken := MakeRefreshToken()
	secondToken := MakeRefreshToken()

	if firstToken == secondToken {
		t.Fatal("MakeRefreshToken returned the same token twice")
	}
}
