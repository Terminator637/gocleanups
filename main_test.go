package gocleanups

import (
	"testing"
)

func TestCleanups(t *testing.T) {
	cls := NewCleanups()
	cls.Add(func() {
		t.Log("Add")
	})
	cls.Run()
	cls.Export()()
}
