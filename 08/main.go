package main

import (
    "os"
    "fmt"
    "strings"
)

type Pos struct {
    r int
    c int
}

func IsInBounds(r, c, rows, cols int) bool {
    if r < 0 || c < 0 || r >= rows || c >= cols {
        return false
    }
    return true
}

func MarkAntinodes(state [][]bool, H, W int, all_pos []Pos) {
    count_pos := len(all_pos)
    for i := 0; i < count_pos; i++ {
        a := all_pos[i]
        for j := i+1; j < count_pos; j++ {
            b := all_pos[j] 
            // vertical dist
            v := a.c - b.c
            h := a.r - b.r
            if v < 0 {
                v *= -1
            }
            if h < 0 {
                h *= -1
            }

            // For the second part of the qeuestion
            state[a.r][a.c] = true
            state[b.r][b.c] = true

            var n_row, n_col int
            // For second part of the question, apply the jump calculation
            // until out of the map (naive way to check different position)
            var max_v int = W / v
            var max_h int = H / h
            m := max_v
            if max_h < max_v {
                m = max_h
            }
            for k := 0; k < m; k++ {
                h2 := h * (k+1)
                v2 := v * (k+1)
                if a.r <= b.r {
                    // a is above b
                    if a.c < b.c {
                        // a is to the left
                        n_row = a.r - h2
                        n_col = a.c - v2
                        if IsInBounds(n_row, n_col, H, W) {
                            state[n_row][n_col] = true
                        }
                        n_row = b.r + h2
                        n_col = b.c + v2
                        if IsInBounds(n_row, n_col, H, W) {
                            state[n_row][n_col] = true
                        }
                    } else {
                        // a is to the right
                        n_row = a.r - h2
                        n_col = a.c + v2
                        if IsInBounds(n_row, n_col, H, W) {
                            state[n_row][n_col] = true
                        }
                        n_row = b.r + h2
                        n_col = b.c - v2
                        if IsInBounds(n_row, n_col, H, W) {
                            state[n_row][n_col] = true
                        }
                    }
                } else {
                    // a is below b
                    if a.c < b.c {
                        // a is to the left of b
                        n_row = a.r + h2
                        n_col = a.c - v2
                        if IsInBounds(n_row, n_col, H, W) {
                            state[n_row][n_col] = true
                        }
                        n_row = b.r - h2
                        n_col = b.c + v2
                        if IsInBounds(n_row, n_col, H, W) {
                            state[n_row][n_col] = true
                        }
                    } else {
                        // a is to the right of b
                        n_row = a.r + h2
                        n_col = a.c + v2
                        if IsInBounds(n_row, n_col, H, W) {
                            state[n_row][n_col] = true
                        }
                        n_row = b.r - h2
                        n_col = b.c - v2
                        if IsInBounds(n_row, n_col, H, W) {
                            state[n_row][n_col] = true
                        }
                    }
                }
            }

        }
    }
}

func main() {
    fmt.Println("hello world")
    dat, err := os.ReadFile("input.txt")
    // dat, err := os.ReadFile("test.txt")
    if err != nil {
        panic("Failed to read file")
    }
    lines := strings.Split(string(dat), "\n")
    H := len(lines) - 1
    if H == 0 {
        panic("Zero sized map (no lines)")
    }
    W := len(lines[0])
    fmt.Println("map size:", H, "x", W)

    // Create an empty grid on which we mark where anti nodes are
    anti_node_state := make([][]bool, H)
    for i := 0; i < H; i++ {
        anti_node_state[i] = make([]bool, W)
        for j := 0; j < W; j++ {
            anti_node_state[i][j] = false
        }
    }
    // Let's create a map of where each antena of same type is
    book := make(map[byte][]Pos)
    for i := 0; i < H; i++ {
        for j := 0; j < W; j++ {
            c := lines[i][j]
            if c == '.' {
                continue
            }
            book[c] = append(book[c], Pos { i, j })
        }
    }
    // Number of frequencies
    fmt.Println("count frequencies:", len(book))
    for freq, all := range book {
        fmt.Println(string(freq), all)
        MarkAntinodes(anti_node_state, H, W, all)
    }
    count := 0
    for i := 0; i < H; i++ {
        for j := 0; j < W; j++ {
            if anti_node_state[i][j] {
                count += 1
            }
        }
    }
    fmt.Println("Unique places with anitnode:", count)
}
