package main

import (
    "os"
    "fmt"
    "strings"
    "strconv"
)

type Set map[int]struct{}
type Rulebook map[int]Set

// Check if according to the rule book having key before other is a safety
// violation
func Violates(rulebook Rulebook, key int, other int) bool {
    set, p := rulebook[other]
    if !p {
        return false
    }
    _, p = set[key]
    return p
}

func IsItSafe(pages []int, rulebook Rulebook) bool {
    L := len(pages)
    for i := 0; i < L; i++ {
        for j := i+1; j < L; j++ {
            if Violates(rulebook, pages[i], pages[j]) {
                return false
            }
        }
    }
    return true
}

// Return zero if the pages are in correct order, otherwise fix the ordering
// and return the value of middle page
func TryToFix(pages []int, rulebook Rulebook) int {
    L := len(pages)
    has_been_fixed := false
    for i := 0; i < L; i++ {
        for j := i+1; j < L; j++ {
            if !Violates(rulebook, pages[i], pages[j]) {
                continue
            }
            has_been_fixed = true
            // Fix the mistake, swap current value with violated value, and
            // start this iteration again
            t := pages[i]
            pages[i] = pages[j]
            pages[j] = t
            i-- // (repeat curret iteration)
            break;
        }
    }
    if has_been_fixed {
        return pages[L / 2]
    }
    return 0
}

func main() {
    fmt.Println("hello world")

    // Store which key should be before what values
    rulebook := make(Rulebook)

    // read & parse input
    dat, err := os.ReadFile("./input.txt")
    // dat, err := os.ReadFile("./test.txt")
    if err != nil {
        panic("Failed to read the input file")
    }

    lines := strings.Split(string(dat), "\n")
    i := 0
    for ; i < len(lines) - 1; i++ {
        line := lines[i]
        if len(line) == 0 {
            // end of rules
            break
        }
        tmp := strings.Split(line, "|")
        if len(tmp) != 2 {
            panic("Failed to parse a rule")
        }
        x, err := strconv.Atoi(tmp[0])
        if err != nil {
            panic("Failed to convert string to int")
        }
        y, err := strconv.Atoi(tmp[1])
        if err != nil {
            panic("Failed to convert string to int")
        }

        set, p := rulebook[x]
        if !p {
            set = make(Set)
            rulebook[x] = set
        }
        set[y] = struct{}{} // create an empty entry
    }
    i += 1 // skip the empty line

    // Check if updates are valid
    sum := 0
    for ; i < len(lines) - 1; i++ {
        t := strings.Split(lines[i], ",")
        pages := make([]int, len(t))
        L := len(t)
        if L % 2 == 0 {
            panic("The update has even number of pages. Which one is the middle page? This must not happen!")
        }
        for j := 0; j < L; j++ {
            pages[j], err = strconv.Atoi(t[j])
            if err != nil {
                fmt.Println("-- string failed to convert to int:", t[j])
                panic("Failed to parse int")
            }
        }
        // if IsItSafe(pages, rulebook) {
        //     mid := pages[L / 2]
        //     sum += mid
        //     // fmt.Println(pages, ":", mid)
        // }
        mid := TryToFix(pages, rulebook)
        sum += mid
    }
    fmt.Println("sum:", sum)
}
