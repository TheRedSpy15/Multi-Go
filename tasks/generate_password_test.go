package tasks_test

/*
   Copyright 2020 TheRedSpy15

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

import (
	"bytes"
	"fmt"
	"testing"

	tks "github.com/TheRedSpy15/Multi-Go/tasks"
)

func TestGeneratePassword(t *testing.T) {

	var buff bytes.Buffer

	tks.GeneratePassword("12", &buff)
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
	tks.GeneratePassword(brokenInput, &buff)
	got = buff.String()
	expected := fmt.Sprintf("Cannot use \"%s\" as length for password\n", brokenInput)

	if got != expected {
		t.Errorf("for non integer input: expected %s, got %s", expected, got)
	}

}
