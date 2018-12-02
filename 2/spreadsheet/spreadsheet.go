package spreadsheet

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Spreadsheet struct {
	Rows []Row
}

type Row []int

func ParseFile(fileName string) *Spreadsheet {
	f, e := os.Open(fileName)

	if e != nil {
		log.Fatal(e)
	}

	bytes, e := ioutil.ReadAll(f)

	if e != nil {
		log.Fatal(e)
	}

	content := fmt.Sprintf("%s", bytes)

	return ParseString(content)
}

func ParseString(value string) *Spreadsheet {
	ss := &Spreadsheet{
		Rows: make([]Row, 0, 100),
	}

	for _, r := range strings.Split(value, "\r\n") {
		cells := strings.Fields(r)
		row := make(Row, 0, 100)

		for _, c := range cells {
			i, err := strconv.Atoi(c)

			if err != nil {
				panic(fmt.Errorf("invalid number %s", c))
			}

			row = append(row, i)
		}

		ss.Rows = append(ss.Rows, row)
	}

	return ss
}
