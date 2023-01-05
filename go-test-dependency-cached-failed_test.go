package gotestdependencycachedfailed

import (
	"testing"
)

func TestA(t *testing.T) {
	// set num value
	num[0] = 15

	// reset
	defer func() {
		num[0] = 0
		// clean cache, try to remove comment and run again
		// aCached = []int{}
	}()

	// get result
	result := A()
	if num[0] != result[0] {
		t.Fail()
	}
}

func TestB(t *testing.T) {
	// set num value
	num[0] = 16

	// reset
	defer func() {
		num[0] = 0
	}()

	// get result
	result := B()
	if num[0] != result[0][0] {
		t.Fail()
	}
}
