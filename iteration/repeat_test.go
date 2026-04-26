package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 4)
	expected := "aaaa"

	if repeated != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	// ... setup ...
	for b.Loop() {
		// ... code to benchmark ...
		Repeat("a", 4)
	}
	// ... teardown ...
}
