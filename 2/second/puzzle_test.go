package second

import (
	"testing"

	"github.com/malafreniere/aoc2017/2/spreadsheet"
)

func TestRowChecksum(t *testing.T) {
	tables := []struct {
		row      string
		expected int
	}{
		{"5 9 2 8", 4},
		{"9 4 7 3", 3},
		{"3 8 6 5", 2},
	}

	for ti, table := range tables {
		ss := spreadsheet.ParseString(table.row)
		firstRow := ss.Rows[0]
		actual := RowChecksum(firstRow)

		if actual != table.expected {
			t.Errorf("(%d) expected %d but got %d with row %s", ti, table.expected, actual, table.row)
		}
	}
}

func TestExecute(t *testing.T) {
	tables := []struct {
		rows     string
		expected int
	}{
		{"5 9 2 8\r\n9 4 7 3\r\n3 8 6 5", 9},
	}

	for ti, table := range tables {
		ss := spreadsheet.ParseString(table.rows)
		actual := Execute(ss)

		if actual != table.expected {
			t.Errorf("(%d) expected %d but got %d with rows %s", ti, table.expected, actual, table.rows)
		}
	}
}
