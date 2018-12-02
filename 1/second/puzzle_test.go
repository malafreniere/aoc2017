package second

import (
	"testing"

	"github.com/malafreniere/aoe2017/1/captcha"
)

func TestExecute(t *testing.T) {
	tables := []struct {
		inputs   string
		expected int
	}{
		{"1212", 6},
		{"1221", 0},
		{"123425", 4},
		{"123123", 12},
		{"12131415", 4},
	}

	for ti, table := range tables {
		inputs := captcha.ParseString(table.inputs)
		actual := Execute(inputs)

		if actual != table.expected {
			t.Errorf("(%d) got %d, but expected %d, for entry %s", ti, actual, table.expected, table.inputs)
		}
	}
}
