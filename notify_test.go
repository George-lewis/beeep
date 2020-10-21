package beeep

import (
	"testing"
)

func TestNotify(t *testing.T) {
	SetAppID("App Name")
	err := Notify("Notify title", "Message body", "assets/information.png")
	if err != nil {
		t.Error(err)
	}
}
