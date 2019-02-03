package util

import (
	"bufio"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func ReadLine() string {
	scanner.Scan()
	return scanner.Text()
}
