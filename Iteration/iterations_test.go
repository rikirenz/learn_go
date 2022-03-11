package iteration

import "testing"

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		Repeat("a", 5)
	}
}

func TestRepeat(t *testing.T) {

	assertCorrectResult := func(t testing.TB, want, got string) {
		t.Helper()
		if got != want {
			t.Errorf("expected %q but got %q", want, got)
		}
	}


	t.Run("5 times a", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertCorrectResult(t, got, want)
	})

	t.Run("0 times empty string", func(t *testing.T) {
		got := Repeat("", 5)
		want := ""
		assertCorrectResult(t, got, want)
	})

}