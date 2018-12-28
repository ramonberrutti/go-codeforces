package main

import "testing"

func TestDecode(t *testing.T) {
	tables := []struct {
		In  []uint
		Out uint
	}{
		{[]uint{1, 3, 3, 7}, 2},
		{[]uint{1, 10000}, 0},
		{[]uint{1, 10000, 3, 4, 8}, 7},
		{[]uint{1, 10000, 2}, 1},
		{[]uint{1, 2, 3, 4, 5, 6, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}, 20},
	}

	for _, table := range tables {
		out := instability(table.In)
		if out != table.Out {
			t.Errorf("Decode of (%v) was incorrect, got: %d, want: %d", table.In, out, table.Out)
		}
	}
}
