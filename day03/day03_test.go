package main

import "testing"

func TestPartOne(t *testing.T) {
	t.Run("testing example one", func(t *testing.T) {
		got := partOne("example.txt")
		want := 161

		if got != want {
			t.Errorf("expected %d but got %d", want, got)
		}
	})

	t.Run("testing example one", func(t *testing.T) {
		got := partTwo("example2.txt")
		want := 48

		if got != want {
			t.Errorf("expected %d but got %d", want, got)
		}
	})
}
