// cui: http request/response tui
// Copyright 2022 Mario Finelli
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"testing"
)

func TestGetStrSliceIndex(t *testing.T) {
	search := []string{"one", "two", "three"}
	tests := []struct {
		input  string
		output int
	}{
		{"one", 0},
		{"three", 2},
		{"four", -1},
	}

	for _, test := range tests {
		if res := getStrSliceIndex(&search, test.input); res != test.output {
			t.Errorf("expected %v got %v", test.output, res)
		}
	}
}
