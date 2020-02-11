package iteration

import (
	"fmt"
	"testing"
)

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", b.N)
	}
}
func ExampleRepeat() {
	fmt.Println(Repeat("a", 5))
	// Output: aaaaa
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}
