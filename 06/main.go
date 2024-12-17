package main

import (
    "fmt"
    "os"
    "strings"
)

func NextCell(r, c, dir int) (int, int) {
        var nr, nc int
        // Find next position
        switch dir {
            case 0:
                // up
                nr = r - 1
                nc = c
            case 1:
                // right
                nr = r
                nc = c + 1
            case 2:
                // down
                nr = r + 1
                nc = c
            case 3:
                //left
                nr = r
                nc = c - 1
        }
        return nr, nc
}

func main() {
    fmt.Println("hello world")
    dat, err := os.ReadFile("input.txt")
    // dat, err := os.ReadFile("test.txt")
    if err != nil {
        panic("Failed to read file")
    }

    grid := strings.Split(string(dat), "\n")
    L := len(grid) - 1
    // Assuming each line has the same length and is the same width as the
    // height of the grid
    if L < 1 {
        fmt.Println("Map is empty")
        return 
    }

    fmt.Println("L:", L)
    // All values must be initialized to false
    var visited [131][131]bool

    // Row and column of current position of the guard
    var r, c, nr, nc int
    var dir int // 0: up, 1: right, 2: down, 3: left

    // Find the position of guard
    for i := 0; i < L; i++ {
        for j := 0; j < L; j++ {
            if grid[i][j] == '^' {
                r = i
                c = j
                break
            }
        }
    }

    // Simulate
    count := 1
    visited[r][c] = true

    for ;; {
        nr, nc = NextCell(r, c, dir)
        if nr < 0 || nr >= L || nc < 0 || nc >= L {
            // out of the map
            break
        }

        if grid[nr][nc] == '#' {
            dir++
            dir %= 4
            continue
        }

        r = nr
        c = nc
        if !visited[r][c] {
            count++
            visited[r][c] = true
        }
    }
    fmt.Println("Number of unique visited positions:", count)
}
