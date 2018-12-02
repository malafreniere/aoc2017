package first

import (
	"testing"

	"github.com/malafreniere/aoc2017/1/captcha"
)

func TestExecute(t *testing.T) {
	tables := []struct {
		inputs   string
		expected int
	}{
		{"1122", 3},
		{"1111", 4},
		{"1234", 0},
		{"91212129", 9},
	}

	for ti, table := range tables {
		inputs := captcha.ParseString(table.inputs)
		actual := Execute(inputs)

		if actual != table.expected {
			t.Errorf("(%d) got %d, but expected %d, for entry %s", ti, actual, table.expected, table.inputs)
		}
	}
}
