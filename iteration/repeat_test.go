package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	// ... setup ...
	for b.Loop() {
		// ... code to benchmark ...
		Repeat("a")
	}
	// ... teardown ...
}