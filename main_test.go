package gocleanups

import (
	"testing"
)

func TestCleanups(t *testing.T) {
	var count int
	cls := NewCleanups(func() {
		count++
	}, func() {
		count++
	}, func() {
		count++
	})
	cls.Add(func() {
		count++
	}, func() {
		count++
	})
	cls.Run()
	wantCount := cls.Len()
	if count != wantCount {
		t.Errorf("expect count++ will be called %d times, got: %d", wantCount, count)
	}
	cls.RunAndReset()
	cls.Export()()
	wantCount = wantCount * 2
	wantLen := 0
	if count != wantCount {
		t.Errorf("expect count++ will be called %d times, got: %d", wantCount, count)
	}
	if cls.Len() != 0 {
		t.Errorf("expect the length of Cleanups is %d, got: %d", wantLen, cls.Len())
	}

}
