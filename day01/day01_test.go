package main

import "testing"

func TestDayOne(t *testing.T) {
	t.Run("testing part one", func(t *testing.T) {
		got := partOne("example.txt")
		want := 11

		if got != want {
			t.Errorf("expected %d but got %d", want, got)
		}
	})

	t.Run("testing part two", func(t *testing.T) {
		got := partTwo("example.txt")
		want := 31

		if got != want {
			t.Errorf("expected %d but got %d", want, got)
		}
	})
}
