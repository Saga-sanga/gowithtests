package integers

import (
	"fmt"
	"testing"
)

func TestSubtract(t *testing.T) {
	actual := Subtract(4, 1)
	expected := 3

	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func ExampleSubtract() {
	res := Subtract(5, 2)
	fmt.Println(res)
	// Output: 3
}
