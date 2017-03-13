package main

import (
	"testing"
	"time"
)

func TestPow(t *testing.T) {
	time.Sleep(time.Duration(15) * time.Second)
	if answer := pow(2, 4); answer != 16 {
		t.Errorf("Incorrect exponentation! Expected %v, got %v", 16, answer)
	}
	if answer := pow(2, 3); answer != 8 {
		t.Errorf("Incorrect exponentation! Expected %v, got %v", 8, answer)
	}
	if answer := pow(3, 3); answer != 27 {
		t.Errorf("Incorrect exponentation! Expected %v, got %v", 27, answer)
	}
}

func TestSquare(t *testing.T) {
	if answer := square(5); answer != 25 {
		t.Errorf("Incorrect square! Expected %v, got %v", 25, answer)
	}
}
