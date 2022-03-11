package integers

import "testing"

func TestAdder(t *testing.T) {
	
	assertCorrectResult := func(t testing.TB, want, got int) {
		t.Helper()
		if got != want {
			t.Errorf("expected '%d' but got '%d'", want, got)
		}
	}

	t.Run("Expected 4", func(t *testing.T) {
		got := Add(2, 2)
		want := 4
		assertCorrectResult(t, got, want)
	})

	t.Run("Expected 6", func(t *testing.T) {
		got := Add(5, 1)
		want := 6
		assertCorrectResult(t, got, want)
	})

}