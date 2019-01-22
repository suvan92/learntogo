package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Eden")
	a := "Welcome my dear Eden"
	if s != a {
		t.Error("Expected", a, "got", s)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Eden"))
	// Output:
	// Welcome my dear Eden
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Eden")
	}
}
