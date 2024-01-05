package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 8)
	expected := "aaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q, but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 8)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 3)
	fmt.Println(repeated)
	// Output: aaa
}

func TestStringsRepeat(t *testing.T) {
	repeated := StringsRepeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("Exptected %q, but got %q", expected, repeated)
	}
}

func BenchmarkStandardRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringsRepeat("a", 8)
	}
}

func TestStringsBuilderRepeat(t *testing.T) {
	repeated := StringsBuilderRepeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("Exptected %q, but got %q", expected, repeated)
	}
}

func BenchmarkStringsBuilderRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringsBuilderRepeat("a", 8)
	}
}

func TestStringsBuilderWithRunesRepeat(t *testing.T) {
	repeated := StringsBuilderWithRunesRepeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("Exptected %q, but got %q", expected, repeated)
	}
}

func BenchmarkStringsBuilderWithRunesRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringsBuilderWithRunesRepeat("a", 8)
	}
}
