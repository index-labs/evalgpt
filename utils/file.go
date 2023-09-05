package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFileLines(filename string, n int) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		err = fmt.Errorf("open file failed: %v", err)
		return nil, err
	}
	defer f.Close()

	lines := make([]string, 0, n)
	scanner := bufio.NewScanner(f)
	for i := 0; i < n; i++ {
		if scanner.Scan() {
			lines = append(lines, strings.TrimSpace(scanner.Text()))
		} else {
			break
		}
	}
	return lines, nil
}
