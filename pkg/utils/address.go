package utils

import "strings"

func NormalizeAddress(address string) string {
	return strings.ToLower(address)
}
