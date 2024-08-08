package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b")
	fmt.Println(repeated)
	// Output: bbbbb
}
