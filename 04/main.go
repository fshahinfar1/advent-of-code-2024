package main

import (
	"fmt"
	"os"
	"strings"
)

type SearchCtx struct {
	grid []string
	width int
	row int
	col int
	key string
	key_len int
}

func Search(ctx SearchCtx) int {
	count := 0
	if ctx.row >= ctx.width || ctx.col >= ctx.width {
		// out of range
		return count
	}
	is_wide := false
	is_high := false

	var excerpt strings.Builder
	var rExcerpt strings.Builder

	// The line is as wide as a key
	if ctx.col <= ctx.width - ctx.key_len {
		is_wide = true
		line := ctx.grid[ctx.row]
		// check forward
		for i := 0; i < ctx.key_len; i++ {
			excerpt.WriteByte(line[ctx.col + i])
			rExcerpt.WriteByte(line[ctx.col + ctx.key_len - 1 - i])
		}
		if excerpt.String() == ctx.key {
			// fmt.Printf("H(F): r:%d c:%d\n", ctx.row, ctx.col)
			count++
		}
		if rExcerpt.String() == ctx.key {
			// fmt.Printf("H(B): r:%d c:%d\n", ctx.row, ctx.col)
			count++
		}
	}
	// The column is as wide as a a key
	excerpt.Reset()
	rExcerpt.Reset()
	if ctx.row <= ctx.width - ctx.key_len {
		is_high = true
		for i := 0; i < ctx.key_len; i++ {
			excerpt.WriteByte(ctx.grid[ctx.row + i][ctx.col])
			rExcerpt.WriteByte(ctx.grid[ctx.row + ctx.key_len - 1 - i][ctx.col])
		}
		if excerpt.String() == ctx.key {
			// fmt.Printf("V(F): r:%d c:%d\n", ctx.row, ctx.col)
			count++
		}
		if rExcerpt.String() == ctx.key {
			// fmt.Printf("V(B): r:%d c:%d\n", ctx.row, ctx.col)
			count++
		}
	}

	if !is_wide || !is_high {
		// no space for diagonal 
		return count
	}

	// Main diagonal
	excerpt.Reset()
	rExcerpt.Reset()
	for i := 0; i < ctx.key_len; i++ {
		excerpt.WriteByte(ctx.grid[ctx.row + i][ctx.col + i])
		j := ctx.key_len - 1 - i
		rExcerpt.WriteByte(ctx.grid[ctx.row + j][ctx.col + j])
	}
	if excerpt.String() == ctx.key {
		// fmt.Printf("MD(F): r:%d c:%d\n", ctx.row, ctx.col)
		count++
	}
	if rExcerpt.String() == ctx.key {
		// fmt.Printf("MD(B): r:%d c:%d\n", ctx.row, ctx.col)
		count++
	}

	// Other diagonal
	excerpt.Reset()
	rExcerpt.Reset()
	for i := 0; i < ctx.key_len; i++ {
		j := ctx.key_len - 1 - i
		excerpt.WriteByte(ctx.grid[ctx.row + i][ctx.col + j])
		rExcerpt.WriteByte(ctx.grid[ctx.row + j][ctx.col + i])
	}
	if excerpt.String() == ctx.key {
		// fmt.Printf("OD(F): r:%d c:%d\n", ctx.row, ctx.col)
		count++
	}
	if rExcerpt.String() == ctx.key {
		// fmt.Printf("OD(B): r:%d c:%d\n", ctx.row, ctx.col)
		count++
	}
	return count
}

func SearchXed(ctx SearchCtx) int {
	if ctx.key_len % 2 == 0 {
		panic("The Xed shaped is not defined for even size key")
	}
	if ctx.row >= ctx.width || ctx.col >= ctx.width {
		// out of range
		return 0
	}
	// Not enough space for diagonal key
	if (ctx.col > ctx.width - ctx.key_len) ||
		 (ctx.row > ctx.width - ctx.key_len) {
		return 0
	}

	md_match := false
	od_match := false

	var excerpt, rExcerpt strings.Builder

	// Main diagonal
	for i := 0; i < ctx.key_len; i++ {
		excerpt.WriteByte(ctx.grid[ctx.row + i][ctx.col + i])
		j := ctx.key_len - 1 - i
		rExcerpt.WriteByte(ctx.grid[ctx.row + j][ctx.col + j])
	}
	if excerpt.String() == ctx.key || rExcerpt.String() == ctx.key {
		md_match = true
	}

	// Other diagonal
	excerpt.Reset()
	rExcerpt.Reset()
	for i := 0; i < ctx.key_len; i++ {
		j := ctx.key_len - 1 - i
		excerpt.WriteByte(ctx.grid[ctx.row + i][ctx.col + j])
		rExcerpt.WriteByte(ctx.grid[ctx.row + j][ctx.col + i])
	}
	if excerpt.String() == ctx.key || rExcerpt.String() == ctx.key {
		od_match = true
	}

	if md_match && od_match {
		return 1
	}
	return 0
}

func main() {
	fmt.Println("hello world")
	dat, err := os.ReadFile("input.txt")
	// dat, err := os.ReadFile("test.txt")
	if err != nil {
		panic("Failed to read the file")
	}
	str := string(dat)
	grid := strings.Split(str, "\n")
	count_line := len(grid) - 1 // last line is empty
	if count_line < 1 {
		panic("Empty file!")
	}
	width_line := len(grid[0])
	fmt.Println("count lines:", count_line)
	fmt.Println("width of line:", width_line)
	if count_line != width_line {
		panic("Input file does not describe a grid")
	}

	// key := "XMAS"
	key := "MAS"
	key_len := len(key)
	ctx := SearchCtx {
		grid: grid,
		width: count_line,
		row: 0,
		col: 0,
		key: key,
		key_len: key_len,
	}
	total := 0
	for i := 0; i < count_line; i++ {
		ctx.row = i
		for j := 0; j < width_line; j++ {
			ctx.col = j
			// tmp := Search(ctx)
			tmp := SearchXed(ctx)
			total += tmp
		}
	}
	fmt.Println("total:", total)
}
