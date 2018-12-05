package gocleanups

import (
	"testing"
)

func TestCleanups(t *testing.T) {
	var count int
	cls := NewCleanups()
	cls.Add(func() {
		count++
	})
	cls.Run()
	cls.Export()()
	wantCount := 2
	if count != 2 {
		t.Errorf("expect count++ will be called %d times, got: %d", wantCount, count)
	}
	cls.RunAndReset()
	cls.Export()()
	if count != 3 {
		t.Errorf("expect count++ will be called %d times, got: %d", wantCount, count)
	}

}
