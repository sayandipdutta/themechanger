package main

import (
	"reflect"
	"testing"
)

var _ = func() bool {
	testing.Init()
	return true
}()

func TestValidateFlags(t *testing.T) {
	tests := []struct {
		theme string
		prog  string
	}{
		{"light", "OneCommander Spyder"},
		{"dark", "OneCommander"},
		{"light", "Spyder"},
		{"dark", "all"},
	}
	for _, tt := range tests {
		if got := validateFlags(tt.theme, tt.prog); !reflect.DeepEqual(got, nil) {
			t.Errorf("validateFlags() = %v, want %v", got, nil)
		}
	}
	tests = []struct {
		theme string
		prog  string
	}{
		{"lights", "OneCommander"},
		{"dark", "pyder"},
		{"light", "all Spyder"},
		{"dark", "hello"},
		{"light", "Sp der"},
	}
	for _, tt := range tests {
		if got := validateFlags(tt.theme, tt.prog); reflect.DeepEqual(got, nil) {
			t.Errorf("validateFlags() = %v, expected error", nil)
		}
	}
}
