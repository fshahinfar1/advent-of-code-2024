package main

import (
    "fmt"
    "os"
    "strings"
)

type PosState struct {
    v bool
    dir[4] bool
}

type State struct {
    grid [][]byte
    L int
    r int
    c int
    dir int
    visited [131][131]PosState
}

func CloneState(s *State) *State {
    var clone State
    clone.grid = s.grid
    clone.L = s.L
    clone.r = s.r
    clone.c = s.c
    clone.dir = s.dir
    for i := 0; i < 131; i++ {
        for j := 0 ; j < 131; j++ {
            clone.visited[i][j].v = s.visited[i][j].v
            for k := 0; k < 4; k++ {
                clone.visited[i][j].dir[k] = s.visited[i][j].dir[k]
            }
        }
    }
    return &clone
}

func VisitState(s *State) {
    r := s.r
    c := s.c
    dir := s.dir
    s.visited[r][c].v = true
    s.visited[r][c].dir[dir] = true
}

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

// @returns number of unique position visited by the guard and number of loops
// found
func Simulate(state *State, mayAddBlock bool) (int, int) {
    var nr, nc int
    count_upos := 0
    count_loop := 0
    for ;; {
        nr, nc = NextCell(state.r, state.c, state.dir)
        if nr < 0 || nr >= state.L || nc < 0 || nc >= state.L {
            // out of the map
            // fmt.Println("no", nr, nc, state.L)
            break
        }

        if state.grid[nr][nc] == '#' {
            state.dir++
            state.dir %= 4
            // In this position we have a new direction
            VisitState(state)
            continue
        }

        if !state.visited[nr][nc].v {
            // fmt.Println(nr, nc)
            count_upos++
            // What if we had put an obstacle here?
            if mayAddBlock {
                tmp_cell_val := state.grid[nr][nc]
                state.grid[nr][nc] = '#'
                clone_state := CloneState(state)
                _, x := Simulate(clone_state, false)
                count_loop += x
                state.grid[nr][nc] = tmp_cell_val
                // fmt.Println("tried but there is no loop!")
            }
        } else {
            // Do we have a loop?
            if state.visited[nr][nc].dir[state.dir] {
                // fmt.Println("here", nr, nc, state.dir)
                count_loop = 1
                break
            }
        }
        state.r = nr
        state.c = nc
        VisitState(state)
    }
    return count_upos, count_loop
}

func main() {
    fmt.Println("hello world")
    dat, err := os.ReadFile("input.txt")
    // dat, err := os.ReadFile("test.txt")
    if err != nil {
        panic("Failed to read file")
    }

    tmp_grid := strings.Split(string(dat), "\n")
    L := len(tmp_grid) - 1
    grid := make([][]byte, L)
    for i := 0; i < L; i++ {
        grid[i] = make([]byte, L)
        for j := 0; j < L; j++ {
            grid[i][j] = tmp_grid[i][j]
        }
    }
    // Assuming each line has the same length and is the same width as the
    // height of the grid
    if L < 1 {
        fmt.Println("Map is empty")
        return 
    }

    fmt.Println("L:", L)

    var state State
    // Find the position of guard
    for i := 0; i < L; i++ {
        for j := 0; j < L; j++ {
            if grid[i][j] == '^' {
                state.r = i
                state.c = j
                break
            }
        }
    }
    state.dir = 0
    state.grid = grid
    state.L = L

    count := 1
    VisitState(&state)

    // For each unique position that the guard visits, either they go there, or
    // we have put an obstacle there. Simulate it!
    // We have a loop when the guard returns to a visited cell with the same direction
    clone_state := CloneState(&state)
    x, y := Simulate(clone_state, true)
    count += x

    fmt.Println("Number of unique visited positions:", count)
    fmt.Println("Number of loops found:", y)
}
