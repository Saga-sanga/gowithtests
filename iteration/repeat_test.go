package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 7)
	fmt.Println(repeated)
	// Output: bbbbbbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
