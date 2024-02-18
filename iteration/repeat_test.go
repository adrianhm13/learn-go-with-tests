package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	expected := "aaaaa"

	if got != expected {
		t.Errorf("Got %q expected %q", got, expected)
	}
}
