package config

import (
	"testing"
)

func TestValidateProgramFlag(t *testing.T) {
	args := [][]string{
		{"OneCommander", "Spyder"},
		{"OneCommander"},
	}
	for ix, arg := range args {
		got := ValidateProgramFlag(arg)
		if got != nil {
			t.Errorf("%d. got: %v", ix, got)
		}
	}
}
