package main

import "testing"

func TestDecode(t *testing.T) {
	tables := []struct {
		In  string
		Out string
	}{
		{"baabbb", "bab"},
		{"ooopppssss", "oops"},
		{"z", "z"},
		{"zww", "zw"},
		{"cooooonnnnttttteeeeeeeeeeeeessssssssttttttttttttttttttt", "coonteestt"},
		{"coodddeeeecccccoooooo", "codeco"},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaa"},
		{"abbcccddddeeeeeffffffggggggghhhhhhhh", "abcdefgh"},
	}

	for _, table := range tables {
		out := decode(table.In)
		if out != table.Out {
			t.Errorf("Decode of (%s) was incorrect, got: %s, want: %s", table.In, out, table.Out)
		}
	}
}
