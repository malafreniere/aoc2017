package captcha

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func ParseFile(fileName string) []int {
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

func ParseString(value string) []int {
	input := make([]int, 0, len(value))

	for _, c := range value {
		val, err := strconv.Atoi(string(c))

		if err != nil {
			panic(fmt.Errorf("invalid number %s", c))
		}

		input = append(input, val)
	}

	return input
}
