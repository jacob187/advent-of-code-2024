package main

import "testing"

func TestDayFour(t *testing.T) {
	t.Run("testing part one", func(t *testing.T) {
		got := partOne("example.txt")
		want := 18

		if got != want {
			t.Errorf("expected %d but got %d", want, got)
		}
	})

	t.Run("testing part two", func(t *testing.T) {
		got := partTwo("example2.txt")
		want := 9

		if got != want {
			t.Errorf("expected %d but got %d", want, got)
		}
	})
}
