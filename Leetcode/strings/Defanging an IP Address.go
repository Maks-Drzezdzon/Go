package strings

import (
	"strings"
)

// DefangIPaddr replaces . with [.]
func DefangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}
