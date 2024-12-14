package main

import (
	"fmt"
    "os"
    "strings"
    "strconv"
    "slices"
)

func canDamp(vals []int, at int) bool {
    l := len(vals)
    if l < 3 {
        // not much to do, just get rid of one and we are left with just one or
        // less levels
        return true
    }
    if at >= l || at < 1 {
        panic("unexpected index!")
    }

    // Check if removing current value solves the issue
    vals_with_out_cur := slices.Concat(nil, vals[:at], vals[at+1:])
    safe_with_prev := IsSafe(vals_with_out_cur, false)
    if safe_with_prev {
        return true
    }

    // Check if removing the previous value solves the issue
    vals_with_out_prev := slices.Concat(nil, vals[:at-1], vals[at:])
    safe_with_cur := IsSafe(vals_with_out_prev, false)
    if safe_with_cur {
        return true
    }

    // check if removing two before can solve the issue
    if at - 2 >= 0 {
        vlas_with_out_2prev := slices.Concat(nil, vals[:at-2], vals[at-1:])
        if IsSafe(vlas_with_out_2prev, false) {
            return true
        }
    }
    return false
}

func IsSafe(vals []int, with_damp bool) bool {
    if len(vals) < 2 {
        return true
    }
    decreasing := false
    if (vals[1] < vals[0]) {
        decreasing = true
    }

    prev := vals[0]
    for i := 1; i < len(vals); i++ {
        if decreasing {
            if prev <= vals[i] {
                if with_damp {
                    return canDamp(vals, i)
                } else {
                    return false
                }
            }
        } else {
            if prev >= vals[i] {
                if with_damp {
                    return canDamp(vals, i)
                } else {
                    return false
                }
            }
        }
        diff := prev - vals[i]
        if diff < 0 {
            diff *= -1
        }
        if diff > 3 || diff < 1 {
            if with_damp {
                return canDamp(vals, i)
            } else {
                return false
            }
        }
        prev = vals[i]
    }
    return true
}

func main() {
	fmt.Println("hello world")
    dat, err := os.ReadFile("./input.txt")
    // dat, err := os.ReadFile("./test2.txt")
    if err != nil {
        panic(err)
    }
    lines := strings.Split(string(dat), "\n")
    count_safe := 0
    for i := 0; i < len(lines) - 1; i++ {
        line := lines[i]
        vals_str := strings.Split(line, " ")
        vals := make([]int, len(vals_str))
        for j := 0; j < len(vals_str); j++ {
            vals[j], err = strconv.Atoi(vals_str[j])
            if err != nil {
                panic(err)
            }
            fmt.Printf("%d ", vals[j])
        }
        fmt.Printf(":")
        if IsSafe(vals, true) {
            count_safe++
            fmt.Println("safe")
        } else {
            fmt.Println("not safe")
        }
    }
    fmt.Println("number of safe reports:", count_safe)
}
