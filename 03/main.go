package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

const RED = 31
const GREEN = 32
const colorNone = "\033[0m"

func PrintMulStr(mem string, from int, bytes int, color byte) {
    // fmt.Println("--", from, bytes)
    fmt.Printf("\033[0;%dm%s%s", color, mem[from:from+bytes+1], colorNone)
}

func IsDigit(x byte) bool {
    if x < '0' || x > '9' {
        return false
    }
    return true
}

func ParseSecondHalf(mem string, off int, X, Y, move *int) bool {
    L := len(mem)
    second_off := off + 4 
    var tmp strings.Builder
    state := 0 // 0: looking for X, 1: looking for comma, 2: looking for end-par 
    for i := second_off; i < L; i++ {
        if i >= L  {
            return false
        }
        c := mem[i]
        is_digit   := IsDigit(c)
        is_comma   := c == ','
        is_end_par := c == ')'
        switch (state) {
        case 0:
            if !is_digit {
                return false
            }
            state = 1
            tmp.WriteByte(c)
        case 1:
            if is_digit {
                tmp.WriteByte(c)
            } else if is_comma {
                if tmp.Len() > 3 {
                    return false // the number is more than 3 digits
                }
                tmpint, err := strconv.Atoi(tmp.String())
                if err != nil {
                    return false
                }
                *X = tmpint
                tmp.Reset()
                state = 2
            } else {
                return false
            }
        case 2:
            if is_digit {
                tmp.WriteByte(c)
            } else if is_end_par {
                if tmp.Len() > 3 {
                    return false // the number is more than 3 digits
                }
                tmpint, err := strconv.Atoi(tmp.String())
                if err != nil {
                    return false
                }
                *Y = tmpint
                *move = i - off
                return true
            } else {
                return false
            }
        }
    }
    return false
}

func main() {
    fmt.Println("hello world")
    dat, err := os.ReadFile("./input.txt")
    if err != nil {
        fmt.Println("Failed to read the file")
        panic(err)
    }

    do_mul := true
    sum := 0
    mem := string(dat)
    L := len(mem)
    // Parsing "mul(X,Y)"
    for i := 0; i < L; i++ {
        fmt.Printf("%c", mem[i])
        if L - i < 5 {
            break
        }

        // Check for do()
        if (mem[i] == 'd' &&
            mem[i+1] == 'o' &&
            mem[i+2] == '(' &&
            mem[i+3] == ')') {
                do_mul = true
                PrintMulStr(mem, i, 3, RED)
                i += 3 // skip
                continue
        }

        if (i + 6 < L &&
            mem[i] == 'd' &&
            mem[i+1] == 'o' &&
            mem[i+2] == 'n' &&
            mem[i+3] == '\'' &&
            mem[i+4] == 't' &&
            mem[i+5] == '(' &&
            mem[i+6] == ')') {
                do_mul = false
                PrintMulStr(mem, i, 6, RED)
                i += 6
                continue
        }

        if !do_mul {
            continue
        }

        if (mem[i] == 'm' &&
            mem[i+1] == 'u' &&
            mem[i+2] == 'l' &&
            mem[i+3] == '(') {
                var X, Y, move int
                if (ParseSecondHalf(mem, i, &X, &Y, &move)) {
                    // successfully parsed the second half
                    sum += X * Y
                    PrintMulStr(mem, i, move, GREEN)
                    i += move
                }
        }
    }
    fmt.Println("\nresult:", sum)
}
