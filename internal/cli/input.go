package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input(message string) string {
	fmt.Print(message)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(input)
}