package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Name provided", func(t *testing.T) {
		got := Hello("Aron")
		want := "Hello, Aron"
		assertCorrectMessage(t, got, want)
	})

	t.Run("No name provided", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

}
