package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	firstCounter := GetInstance()

	if firstCounter == nil {
		t.Error("pointer to singleton should not be nil")
	}

	secondCounter := GetInstance()

	if firstCounter != secondCounter {
		t.Error("pointer to singleton should return the same pointer")
	}
}

func TestAddOne(t *testing.T) {
	firstCounter := GetInstance()
	currentCount := firstCounter.AddOne()

	if currentCount != 1 {
		t.Error("AddOne should be called once")
	}

	secondCounter := GetInstance()
	currentCount = secondCounter.AddOne()

	if currentCount != 2 {
		t.Error("AddOne should act on the same singleton instance")
	}
}
