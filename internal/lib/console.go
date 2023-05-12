package lib

import (
	"bufio"
	"os"
)

func WaitForConsoleInput(success string) bool {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "y" {
			return true
		}
		return false
	}
	return false
}
