package random

import (
	"testing"
)

func TestRandomString(t *testing.T) {
	size := 10
	testLen := len(RandString(size))

	// Not millisecond accurate above, so...
	if testLen != size {
		t.Fatalf("expected length `%d`, got `%d`", size, testLen)
	}
}
