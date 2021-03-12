package messaging

import (
	"fmt"
	"regexp"
)

func removePassword(connectionString string) string {
	re := regexp.MustCompile(`:(\w+)@`)
	return fmt.Sprintf(re.ReplaceAllStringFunc(connectionString, func(a string) string { return ":*******@" }))
}
