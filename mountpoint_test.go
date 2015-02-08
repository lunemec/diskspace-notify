package main

import (
	"syscall"
	"testing"
)

func TestPercentAvailable(t *testing.T) {
	data := new(syscall.Statfs_t)
	data.Blocks = uint64(0)
	data.Bavail = uint64(0)
	got := percentAvailable(data)
	want := uint8(0)

	if got != want {
		t.Errorf("func percentAvailable got %v, want %v", got, want)
	}

	data.Blocks = uint64(1000000)
	data.Bavail = uint64(500000)
	got = percentAvailable(data)
	want = uint8(50)

	if got != want {
		t.Errorf("func percentAvailable got %v, want %v", got, want)
	}
}
