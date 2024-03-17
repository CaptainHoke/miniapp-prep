package main

import "testing"

func TestEchoArgsNoError(t *testing.T) {
	err := echo([]string{"bin-name", "1", "2", "Ohayo"})
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestEchoNoArgsError(t *testing.T) {
	err := echo([]string{"bin-name"})
	if err == nil {
		t.Error("Expected error, but found nil")
	}
}

func TestFailsOnPurpose(t *testing.T) {
	t.Error("Makefile test")
}
