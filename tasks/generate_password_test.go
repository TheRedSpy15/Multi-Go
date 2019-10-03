package tasks

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGeneratePassword(t *testing.T) {

	var buff bytes.Buffer

	GeneratePassword("12", &buff)
	got := buff.String()

	// consider preceeding "password:" and trailing "\n"
	if len(got) != 22 {
		t.Errorf("Expected length of 22, got %d", len(got))
	}
	if got[len(got)-1:] != "\n" {
		t.Errorf("Expected trailing newline, got %s", got[len(got)-1:])
	}
	buff.Reset()

	brokenInput := "this can't be parsed as integer"
	GeneratePassword(brokenInput, &buff)
	got = buff.String()
	expected := fmt.Sprintf("Cannot use \"%s\" as length for password\n", brokenInput)

	if got != expected {
		t.Errorf("for non integer input: expected %s, got %s", expected, got)
	}

}
