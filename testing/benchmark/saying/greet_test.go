package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Tanmay Kakodia")
	expected := "Welcome my dear Tanmay Kakodia"
	if s != expected {
		t.Error("Expected", expected, "but Got", s)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Yashika Kakodia"))
	// Output:
	// Welcome my dear Yashika Kakodia
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Rekha Yadav")
	}
}

func BenchmarkGreetT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GreetT("Rekha Yadav")
	}
}
