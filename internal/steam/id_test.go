package steam_test

import (
	"airforce/internal/steam"
	"testing"
)

func TestValidSteamID(t *testing.T) {
	if !steam.IsSteamID("STEAM_0:0:453245764") {
		t.Error("panic")
	}
}

func TestInvalidSteamID(t *testing.T) {
	if steam.IsSteamID("blah blah blah") {
		t.Error("panic")
	}
}

func TestValidSteamID64(t *testing.T) {
	if !steam.IsSteamID64("76561198866757256") {
		t.Error("panic")
	}
}

func TestInvalidSteamID64(t *testing.T) {
	if steam.IsSteamID64("not a steam id :(") {
		t.Error("panic")
	}
}

func TestInvalidNumberSteamID64(t *testing.T) {
	if steam.IsSteamID64("12345678901234567") {
		t.Error("panic")
	}
}
