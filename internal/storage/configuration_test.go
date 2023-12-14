package storage

import (
	"testing"
)

func Test(t *testing.T) {
	name := "test"
	want := "test"

	if name != want {
		t.Fatalf(`Error`)
	}
}
