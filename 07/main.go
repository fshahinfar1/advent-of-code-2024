package main
import (
	"os"
	"fmt"
	"strings"
	"errors"
	"strconv"
)

func ParseLine(line string) (int, []int, error) {
	/* find the `:` */
	L := len(line)
	colon_index := -1
	for i := 0; i < L; i++ {
		if line[i] == ':' {
			colon_index = i
			break
		}
	}
	if colon_index == -1 {
		return 0, nil, errors.New("did not found `:`")
	}
	res, err := strconv.Atoi(line[:colon_index])
	if err != nil {
		return 0, nil, err
	}

	num_str := strings.Split(line[colon_index+2:], " ")
	count := len(num_str)
	nums := make([]int, count)
	for i := 0; i < count; i++ {
		nums[i], err = strconv.Atoi(num_str[i])
		if err != nil {
			return 0, nil, err
		}
	}

	return res, nums, nil
}

type State struct {
	offset int
	sum int
}

func ConcatInts(a, b int) int {
	count := 0
	tmp := b
	if tmp == 0 {
		count = 1
	} else {
		for ; tmp > 0; {
			tmp /= 10
			count++
		}
	}
	tmp = a
	for i := 0; i < count; i++ {
		tmp *= 10
	}
	tmp += b
	return tmp
}

func IsEq(res int, numbers []int) bool {
	L := len(numbers)
	if L == 0 {
		return false
	}
	queue := make([]State, 0)
	queue = append(queue, State {
		offset: 1,
		sum: numbers[0],
	})
	// BFS
	for ;; {
		if len(queue) == 0 {
			break;
		}
		// pop head
		s := queue[0]
		queue = queue[1:]
		if s.offset >= L {
			if s.sum == res {
				// we found a combination that can work
				return true
			}
			// there is no more number to +/*. Investigate next state
			continue
		}
		// append tail
		queue = append(queue, State {
			sum: s.sum + numbers[s.offset],
			offset: s.offset + 1,
		});
		queue = append(queue, State {
			sum: s.sum * numbers[s.offset],
			offset: s.offset + 1,
		});
		// For the second part of the question (support || operator)
		queue = append(queue, State {
			sum: ConcatInts(s.sum, numbers[s.offset]),
			offset: s.offset + 1,
		});
	}
	return false
}

func main() {
	fmt.Println("hello world")
	dat, err := os.ReadFile("input.txt")
	// dat, err := os.ReadFile("test.txt")
	if err != nil {
		panic("Failed to read input file")
	}
	lines := strings.Split(string(dat), "\n")
	count_line := len(lines) - 1
	sum := 0
	for i:= 0; i < count_line; i++ {
		res, numbers, err := ParseLine(lines[i])
		if err != nil {
			panic("Failed to parse the line")
		}
		if !IsEq(res, numbers) {
			// skip to the next line
			continue
		}
		// fmt.Println(res, "=", numbers)
		sum += res
	}
	fmt.Println("Total sum:", sum)
}
