package helper

import (
	"fmt"
	"strings"
)

func ReplacePlaceholder(q string, start int) string {
	ph := "?"
	count := strings.Count(q, ph)
	for i := start; i <= count; i++ {
		substr := fmt.Sprint("$", i)
		q = strings.Replace(q, ph, substr, 1)
	}
	return q
}
