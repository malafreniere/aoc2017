package first

import (
	"testing"

	"github.com/malafreniere/aoc2017/2/spreadsheet"
)

func TestRowChecksum(t *testing.T) {
	tables := []struct {
		row      string
		expected int
	}{
		{"5 1 9 5", 8},
		{"7 5 3", 4},
		{"2 4 6 8", 6},
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
		{"5 1 9 5\r\n7 5 3\r\n2 4 6 8", 18},
	}

	for ti, table := range tables {
		ss := spreadsheet.ParseString(table.rows)
		actual := Execute(ss)

		if actual != table.expected {
			t.Errorf("(%d) expected %d but got %d with rows %s", ti, table.expected, actual, table.rows)
		}
	}
}
