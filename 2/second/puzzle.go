package second

import (
	"fmt"
	"math"

	"github.com/malafreniere/aoc2017/2/spreadsheet"
)

func Execute(ss *spreadsheet.Spreadsheet) int {
	checksum := 0
	for _, r := range ss.Rows {
		checksum += RowChecksum(r)
	}

	return checksum
}

func RowChecksum(r spreadsheet.Row) int {
	for i, c := range r {
		for _, c2 := range r[i+1:] {
			f := float64(c) / float64(c2)
			f2 := float64(c2) / float64(c)

			if math.Floor(f) == math.Ceil(f) {
				return int(f)
			}

			if math.Floor(f2) == math.Ceil(f2) {
				return int(f2)
			}
		}

	}

	panic(fmt.Errorf("found none !"))
}
