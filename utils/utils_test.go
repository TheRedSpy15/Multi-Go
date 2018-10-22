package utils

/*
   Copyright 2018 TheRedSpy15

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
	"errors"
	"testing"
)

func TestBytesToGigabytes(t *testing.T) {
	got := BytesToGigabytes(4034846720)
	want := 4.03

	if got != want {
		t.Errorf("got '%f' want '%f'", got, want)
	}
}

func TestRandomString(t *testing.T) {
	got := len(RandomString(5))
	want := 5

	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

func TestCheckEmptyTargetShouldPanic(t *testing.T) {
	assertPanic(t, func() {
		CheckTarget("")
	})
}

func TestCheckErrorShouldPanic(t *testing.T) {
	assertPanic(t, func() {
		CheckErr(errors.New("an unknown error"))
	})
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	f()
}
