package main

import "testing"

func TestDayTwo(t *testing.T) {
	t.Run("testing part one example file", func(t *testing.T) {
		got := partOne("example.txt")
		want := 2

		if got != want {
			t.Errorf("expected %d but got %d", want, got)
		}
	})

	t.Run("testing part two example file", func(t *testing.T) {
		got := partTwo("example.txt")
		want := 4

		if got != want {
			t.Errorf("expected %d but got %d", want, got)
		}
	})
}
