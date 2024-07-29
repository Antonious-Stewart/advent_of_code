package dayone

import (
	"fmt"
	"strconv"

	aoc_utils "github.com/Antonious-Stewart/advent_of_code/utils"
)

func Puzzle() int64 {
	var result int64
	scanner := aoc_utils.ScanFile()

	for scanner.Scan() {
		txt := scanner.Text()
		var start, end int64
		for _, char := range txt {
			if num, err := strconv.ParseInt(string(char), 10, 64); err == nil {
				if start == 0 {
					start = num
				}
				end = num
			}
		}

		if code, err := strconv.ParseInt(fmt.Sprintf("%d%d", start, end), 10, 64); err != nil {
			panic(err)
		} else {
			result += code
		}
	}

	return result
}
