package stdlib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//ReadAllInts 读取所有数值
// 默认每行为一个整数值
func ReadAllInts() ([]int, error) {
	var results []int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if i, err := strconv.Atoi(strings.TrimSpace(scanner.Text())); err == nil {
			results = append(results, i)
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		return results, err
	}
	return results, nil
}
