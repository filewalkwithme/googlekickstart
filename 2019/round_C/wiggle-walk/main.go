package main

import (
	"bufio"
	"fmt"

	//"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	nLines := strToInt(readLine())
	id := 0
	for nLines > 0 {
		id++

		f := strings.Split(readLine(), " ")
		n := int(strToInt(f[0]))
		r := int(strToInt(f[1]))
		c := int(strToInt(f[2]))
		sr := int(strToInt(f[3]))
		sc := int(strToInt(f[4]))

		directions := []rune(readLine())

		fmt.Printf("Case #%d: %s", id, solve(r, c, sr-1, sc-1, directions, n))
		nLines--
	}
}

func solve(r0, c0, sr, sc int, directions []rune, n int) string {
	mR := make([][]interval, r0)
	mC := make([][]interval, c0)

	r := sr
	c := sc

	addInterval(&mR[r], c)
	addInterval(&mC[c], r)

	for i := int(0); i < n; i++ {
		d := directions[i]

		if d == 'E' {
			for i := 0; i < len(mR[r]); i++ {
				if c >= mR[r][i].begin && c <= mR[r][i].end {
					c = mR[r][i].end + 1
					break
				}
			}
		}

		if d == 'W' {
			for i := 0; i < len(mR[r]); i++ {
				if c >= mR[r][i].begin && c <= mR[r][i].end {
					c = mR[r][i].begin - 1
					break
				}
			}
		}

		if d == 'S' {
			for i := 0; i < len(mC[c]); i++ {
				if r >= mC[c][i].begin && r <= mC[c][i].end {
					r = mC[c][i].end + 1
					break
				}
			}
		}

		if d == 'N' {
			for i := 0; i < len(mC[c]); i++ {
				if r >= mC[c][i].begin && r <= mC[c][i].end {
					r = mC[c][i].begin - 1
					break
				}
			}
		}

		addInterval(&mR[r], c)
		addInterval(&mC[c], r)
	}

	return fmt.Sprintf("%d %d\n", r+1, c+1)
}

type interval struct {
	begin int
	end   int
}

func addInterval(input *[]interval, newElem int) {
	newIntv := interval{begin: newElem, end: newElem}

	// If the arrya is empty we simply need to insert a new interval
	if len(*input) == 0 {
		*input = append(*input, newIntv)
		return
	}

	// If the array contains a single element, we have to check if we need to
	// grow an existing interval or insert a new interval
	if len(*input) == 1 {
		// Grow left
		if newElem == (*input)[0].begin-1 {
			(*input)[0].begin--
			return
		}

		// Grow right
		if newElem == (*input)[0].end+1 {
			(*input)[0].end++
			return
		}

		// Insert the new interval at left
		if newElem < (*input)[0].begin {
			new := []interval{newIntv}
			new = append(new, (*input)[0])
			*input = new
			return
		}

		// Insert the new interval at right
		if newElem > (*input)[0].end {
			*input = append(*input, newIntv)
			return
		}
	}

	// If the array contains at least 2 intervals, we need to check if we need
	// to merge existing intervals
	if len(*input) >= 2 {
		// Insert the new interval at left
		if newElem < (*input)[0].begin-1 {
			new := []interval{newIntv}
			new = append(new, (*input)...)
			*input = new
			return
		}

		// Insert the new interval at right
		if newElem > (*input)[len(*input)-1].end+1 {
			*input = append(*input, newIntv)
			return
		}

		// Check if we need to insert a new interval between two existing
		// intervals
		for i := 0; i < len((*input))-1; i++ {
			// The new interval is between two intervals that are separated by
			// a single unit. Ex:
			// [1..4] - [6..7], new element: 5
			// Merge them!
			// [1..7]
			if newElem == (*input)[i].end+1 && newElem == (*input)[i+1].begin-1 {
				(*input)[i].end = (*input)[i+1].end
				new := []interval{}
				new = append(new, (*input)[:i+1]...)
				new = append(new, (*input)[i+2:]...)
				(*input) = new
				return
			}

			// The new element will be placed between two intervals, which are
			// separate by more than 2 units. Ex:
			// [1..4] - [10..12], new element: 7
			// [1..4] [7..7] [10..12]
			if newElem > (*input)[i].end+1 &&
				newElem < (*input)[i+1].begin-1 {
				new := []interval{}
				new = append(new, (*input)[:i+1]...)
				new = append(new, newIntv)
				new = append(new, (*input)[i+1:]...)
				(*input) = new
				return
			}
		}

		for i := 0; i < len((*input)); i++ {
			// element is already contained in an existing interval
			if newElem >= (*input)[i].begin &&
				newElem <= (*input)[i].end {
				// Do nothing
				return
			}

			// Grow left
			if newElem == (*input)[i].begin-1 {
				(*input)[i].begin--
				return
			}

			// Grow right
			if newElem == (*input)[i].end+1 {
				(*input)[i].end++
				return
			}
		}
	}
}

// UTILS

var reader *bufio.Reader

func init() {
	rand.Seed(time.Now().UnixNano())
	reader = bufio.NewReader(os.Stdin)
}

func readLine() string {
	buf := []byte{}
	for {
		b, isPrefix, err := reader.ReadLine()
		if err != nil {
			return fmt.Sprintf("ERROR-readline: %v\n", err)
		}

		buf = append(buf, b...)
		if !isPrefix {
			break
		}

	}
	return string(buf)
}

func strToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
