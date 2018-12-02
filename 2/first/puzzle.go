package first

import "github.com/malafreniere/aoc2017/2/spreadsheet"

func Execute(ss *spreadsheet.Spreadsheet) int {
	checksum := 0
	for _, r := range ss.Rows {
		checksum += RowChecksum(r)
	}

	return checksum
}

func RowChecksum(r spreadsheet.Row) int {
	largest := r[0]
	smallest := r[0]

	for _, c := range r {
		if c < smallest {
			smallest = c
		}

		if c > largest {
			largest = c
		}
	}

	return largest - smallest
}
